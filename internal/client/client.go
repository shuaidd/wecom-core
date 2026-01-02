package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/shuaidd/wecom-core/internal/auth"
	"github.com/shuaidd/wecom-core/internal/errors"
	"github.com/shuaidd/wecom-core/internal/retry"
	"github.com/shuaidd/wecom-core/pkg/logger"
)

// contextKey 用于在 context 中存储值的类型
type contextKey string

const (
	// traceIDKey TraceId 的 context key
	traceIDKey contextKey = "trace_id"
)

// WithTraceID 将 TraceId 添加到 context
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

// getTraceID 从 context 中获取 TraceId
func getTraceID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if traceID, ok := ctx.Value(traceIDKey).(string); ok {
		return traceID
	}
	return ""
}

// Client HTTP客户端
type Client struct {
	// httpClient 底层HTTP客户端
	httpClient *http.Client
	// baseURL API基础URL
	baseURL string
	// logger 日志记录器
	logger logger.Logger
	// tokenManager Token管理器
	tokenManager *auth.TokenManager
	// retryExecutor 重试执行器
	retryExecutor *retry.Executor
	// debug 是否打印请求和响应详情
	debug bool
}

// New 创建HTTP客户端
func New(baseURL string, timeout time.Duration, log logger.Logger, tm *auth.TokenManager, re *retry.Executor) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		baseURL:       baseURL,
		logger:        log,
		tokenManager:  tm,
		retryExecutor: re,
		debug:         false,
	}
}

// SetDebug 设置是否打印请求和响应详情
func (c *Client) SetDebug(debug bool) *Client {
	c.debug = debug
	return c
}

// Do 执行HTTP请求（带自动 token 和重试）
func (c *Client) Do(ctx context.Context, req *Request) (*Response, error) {
	var resp *Response

	// 使用重试策略执行请求
	err := c.retryExecutor.Do(ctx, func() error {
		// 1. 获取 access_token
		token, err := c.tokenManager.GetToken(ctx)
		if err != nil {
			return fmt.Errorf("failed to get access token: %w", err)
		}

		// 2. 添加 token 到请求
		req.AddQuery("access_token", token)

		// 3. 构建 HTTP 请求
		httpReq, err := req.BuildHTTPRequest(ctx, c.baseURL)
		if err != nil {
			return fmt.Errorf("failed to build http request: %w", err)
		}

		// 4. 记录请求日志
		startTime := time.Now()
		c.logger.Debug("API Request", withTraceID(ctx,
			logger.F("method", httpReq.Method),
			logger.F("url", httpReq.URL.String()))...)

		// 4.1. Debug模式：打印请求详情
		if c.debug {
			c.logRequestDetails(ctx, httpReq, req.Body)
		}

		// 5. 发送请求
		httpResp, err := c.httpClient.Do(httpReq)
		if err != nil {
			duration := time.Since(startTime)
			c.logger.Error("Request failed", withTraceID(ctx,
				logger.F("error", err),
				logger.F("duration", duration))...)
			return fmt.Errorf("http request failed: %w", err)
		}
		defer httpResp.Body.Close()

		// 6. 解析响应
		resp, err = ParseResponse(httpResp)
		duration := time.Since(startTime)

		// 6.1. Debug模式：打印响应详情
		if c.debug {
			c.logResponseDetails(ctx, httpResp.StatusCode, resp)
		}

		if err != nil {
			c.logger.Error("Request failed", withTraceID(ctx,
				logger.F("url", httpReq.URL.String()),
				logger.F("errcode", resp.ErrCode),
				logger.F("errmsg", resp.ErrMsg),
				logger.F("duration", duration))...)

			// 7. Token 失效，刷新后重试
			if errors.IsTokenExpired(err) {
				c.logger.Warn("Token expired, refreshing", withTraceID(ctx,
					logger.F("errcode", resp.ErrCode))...)
				if refreshErr := c.tokenManager.RefreshToken(ctx); refreshErr != nil {
					c.logger.Error("Failed to refresh token", withTraceID(ctx,
						logger.F("error", refreshErr))...)
				}
			}

			return err
		}

		// 8. 记录成功日志
		c.logger.Info("API Request successful", withTraceID(ctx,
			logger.F("url", httpReq.URL.String()),
			logger.F("duration", duration))...)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Get 发送GET请求
func (c *Client) Get(ctx context.Context, path string, query url.Values) (*Response, error) {
	req := NewRequest(MethodGet, path)
	if query != nil {
		req.Query = query
	}
	return c.Do(ctx, req)
}

// Post 发送POST请求
func (c *Client) Post(ctx context.Context, path string, body any) (*Response, error) {
	req := NewRequest(MethodPost, path).SetBody(body)
	return c.Do(ctx, req)
}

// DoAndUnmarshal 执行请求并自动解析响应到指定类型
func DoAndUnmarshal[T any](c *Client, ctx context.Context, req *Request) (*T, error) {
	resp, err := c.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	var result T
	if err := resp.Unmarshal(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAndUnmarshal 发送GET请求并自动解析响应
func GetAndUnmarshal[T any](c *Client, ctx context.Context, path string, query url.Values) (*T, error) {
	req := NewRequest(MethodGet, path)
	if query != nil {
		req.Query = query
	}
	return DoAndUnmarshal[T](c, ctx, req)
}

// PostAndUnmarshal 发送POST请求并自动解析响应
func PostAndUnmarshal[T any](c *Client, ctx context.Context, path string, body any) (*T, error) {
	req := NewRequest(MethodPost, path).SetBody(body)
	return DoAndUnmarshal[T](c, ctx, req)
}

// logRequestDetails 打印请求详情
func (c *Client) logRequestDetails(ctx context.Context, httpReq *http.Request, body any) {
	c.logger.Info("==> Request Details", withTraceID(ctx,
		logger.F("method", httpReq.Method),
		logger.F("url", httpReq.URL.String()))...)

	if body != nil {
		bodyJSON, err := json.MarshalIndent(body, "", "  ")
		if err != nil {
			c.logger.Warn("Failed to marshal request body", withTraceID(ctx,
				logger.F("error", err))...)
		} else {
			c.logger.Info("Request Body", withTraceID(ctx,
				logger.F("body", string(bodyJSON)))...)
		}
	}
}

// logResponseDetails 打印响应详情
func (c *Client) logResponseDetails(ctx context.Context, statusCode int, resp *Response) {
	c.logger.Info("<== Response Details", withTraceID(ctx,
		logger.F("status_code", statusCode),
		logger.F("errcode", resp.ErrCode),
		logger.F("errmsg", resp.ErrMsg))...)

	if len(resp.Body) > 0 {
		var prettyJSON map[string]any
		if err := json.Unmarshal(resp.Body, &prettyJSON); err == nil {
			bodyJSON, _ := json.MarshalIndent(prettyJSON, "", "  ")
			c.logger.Info("Response Body", withTraceID(ctx,
				logger.F("body", string(bodyJSON)))...)
		} else {
			c.logger.Info("Response Body", withTraceID(ctx,
				logger.F("body", string(resp.Body)))...)
		}
	}
}

// withTraceID 为日志字段添加 TraceId
func withTraceID(ctx context.Context, fields ...logger.Field) []logger.Field {
	traceID := getTraceID(ctx)
	if traceID == "" {
		return fields
	}
	// 将 TraceId 添加到字段列表的开头
	result := make([]logger.Field, 0, len(fields)+1)
	result = append(result, logger.F("trace_id", traceID))
	result = append(result, fields...)
	return result
}
