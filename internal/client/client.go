package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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
	// agentNameKey 应用名称的 context key
	agentNameKey contextKey = "agent_name"
	// agentIDKey 应用ID的 context key
	agentIDKey contextKey = "agent_id"
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

// WithAgentName 将应用名称添加到 context
func WithAgentName(ctx context.Context, agentName string) context.Context {
	return context.WithValue(ctx, agentNameKey, agentName)
}

// WithAgentID 将应用ID添加到 context
func WithAgentID(ctx context.Context, agentID int64) context.Context {
	return context.WithValue(ctx, agentIDKey, agentID)
}

// getAgentKey 从 context 中获取应用标识（优先使用名称，其次使用ID）
func getAgentKey(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	// 优先使用 agentName
	if agentName, ok := ctx.Value(agentNameKey).(string); ok && agentName != "" {
		return agentName
	}

	// 其次使用 agentID
	if agentID, ok := ctx.Value(agentIDKey).(int64); ok && agentID > 0 {
		return fmt.Sprintf("%d", agentID)
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
	// interceptors 拦截器
	interceptors *Interceptors
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
		interceptors:  NewInterceptors(),
		debug:         false,
	}
}

// SetDebug 设置是否打印请求和响应详情
func (c *Client) SetDebug(debug bool) *Client {
	c.debug = debug
	return c
}

// AddRequestInterceptor 添加请求拦截器
func (c *Client) AddRequestInterceptor(interceptor RequestInterceptor) *Client {
	c.interceptors.AddRequestInterceptor(interceptor)
	return c
}

// AddResponseInterceptor 添加响应拦截器（解析前）
func (c *Client) AddResponseInterceptor(interceptor ResponseInterceptor) *Client {
	c.interceptors.AddResponseInterceptor(interceptor)
	return c
}

// AddAfterResponseInterceptor 添加响应后拦截器（解析后）
func (c *Client) AddAfterResponseInterceptor(interceptor AfterResponseInterceptor) *Client {
	c.interceptors.AddAfterResponseInterceptor(interceptor)
	return c
}

// Do 执行HTTP请求（带自动 token 和重试）
func (c *Client) Do(ctx context.Context, req *Request) (*Response, error) {
	var resp *Response

	// 使用重试策略执行请求
	err := c.retryExecutor.Do(ctx, func() error {
		// 1. 从 context 获取应用标识
		agentKey := getAgentKey(ctx)

		// 2. 获取 access_token（根据应用标识）
		token, err := c.tokenManager.GetTokenByAgent(ctx, agentKey)
		if err != nil {
			return fmt.Errorf("failed to get access token: %w", err)
		}

		// 3. 添加 token 到请求
		req.AddQuery("access_token", token)

		// 4. 构建 HTTP 请求
		httpReq, err := req.BuildHTTPRequest(ctx, c.baseURL)
		if err != nil {
			return fmt.Errorf("failed to build http request: %w", err)
		}

		// 4.1. 执行请求前拦截器
		if err := c.interceptors.executeRequestInterceptors(ctx, httpReq, req.Body); err != nil {
			c.logger.Error("Request interceptor failed", withTraceID(ctx,
				logger.F("error", err))...)
			return fmt.Errorf("request interceptor failed: %w", err)
		}

		// 5. 记录请求日志
		startTime := time.Now()
		c.logger.Debug("API Request", withTraceID(ctx,
			logger.F("method", httpReq.Method),
			logger.F("url", httpReq.URL.String()),
			logger.F("agent_key", agentKey))...)

		// 5.1. Debug模式：打印请求详情
		if c.debug {
			c.logRequestDetails(ctx, httpReq, req.Body)
		}

		// 6. 发送请求
		httpResp, err := c.httpClient.Do(httpReq)
		if err != nil {
			duration := time.Since(startTime)
			c.logger.Error("Request failed", withTraceID(ctx,
				logger.F("error", err),
				logger.F("duration", duration))...)
			return fmt.Errorf("http request failed: %w", err)
		}
		defer httpResp.Body.Close()

		// 6.1. 执行响应前拦截器（解析前）
		if err := c.interceptors.executeResponseInterceptors(ctx, httpResp); err != nil {
			c.logger.Error("Response interceptor failed", withTraceID(ctx,
				logger.F("error", err))...)
			return fmt.Errorf("response interceptor failed: %w", err)
		}

		// 7. 解析响应
		resp, err = ParseResponse(httpResp)
		duration := time.Since(startTime)

		// 7.1. Debug模式：打印响应详情
		if c.debug {
			c.logResponseDetails(ctx, httpResp.StatusCode, resp)
		}

		if err != nil {
			c.logger.Error("Request failed", withTraceID(ctx,
				logger.F("url", httpReq.URL.String()),
				logger.F("errcode", resp.ErrCode),
				logger.F("errmsg", resp.ErrMsg),
				logger.F("duration", duration))...)

			// 8. Token 失效，刷新后重试
			if errors.IsTokenExpired(err) {
				c.logger.Warn("Token expired, refreshing", withTraceID(ctx,
					logger.F("errcode", resp.ErrCode),
					logger.F("agent_key", agentKey))...)
				if refreshErr := c.tokenManager.RefreshTokenByAgent(ctx, agentKey); refreshErr != nil {
					c.logger.Error("Failed to refresh token", withTraceID(ctx,
						logger.F("error", refreshErr),
						logger.F("agent_key", agentKey))...)
				}
			}

			return err
		}

		// 9. 记录成功日志
		c.logger.Info("API Request successful", withTraceID(ctx,
			logger.F("url", httpReq.URL.String()),
			logger.F("duration", duration))...)

		// 9.1. 执行响应后拦截器（解析后）
		if err := c.interceptors.executeAfterResponseInterceptors(ctx, resp); err != nil {
			c.logger.Error("After response interceptor failed", withTraceID(ctx,
				logger.F("error", err))...)
			return fmt.Errorf("after response interceptor failed: %w", err)
		}

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

// PostMultipart 发送multipart/form-data POST请求
func (c *Client) PostMultipart(ctx context.Context, path string, query url.Values, body []byte, contentType string) (*Response, error) {
	req := NewMultipartRequest(path, body, contentType)
	if query != nil {
		req.Query = query
	}
	return c.Do(ctx, req)
}

// PostMultipartAndUnmarshal 发送multipart/form-data POST请求并自动解析响应
func PostMultipartAndUnmarshal[T any](c *Client, ctx context.Context, path string, body []byte, contentType string) (*T, error) {
	resp, err := c.PostMultipart(ctx, path, nil, body, contentType)
	if err != nil {
		return nil, err
	}

	var result T
	if err := resp.Unmarshal(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// PostMultipartAndUnmarshalWithQuery 发送带查询参数的multipart/form-data POST请求并自动解析响应
func PostMultipartAndUnmarshalWithQuery[T any](c *Client, ctx context.Context, path string, query url.Values, body []byte, contentType string) (*T, error) {
	resp, err := c.PostMultipart(ctx, path, query, body, contentType)
	if err != nil {
		return nil, err
	}

	var result T
	if err := resp.Unmarshal(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetMedia 下载媒体文件
func (c *Client) GetMedia(ctx context.Context, path string, query url.Values, headers map[string]string) ([]byte, error) {
	var result []byte

	// 使用重试策略执行请求
	err := c.retryExecutor.Do(ctx, func() error {
		// 1. 从 context 获取应用标识
		agentKey := getAgentKey(ctx)

		// 2. 获取 access_token（根据应用标识）
		token, err := c.tokenManager.GetTokenByAgent(ctx, agentKey)
		if err != nil {
			return fmt.Errorf("failed to get access token: %w", err)
		}

		// 3. 添加 token 到查询参数
		if query == nil {
			query = url.Values{}
		}
		query.Set("access_token", token)

		// 4. 构建完整URL
		u, err := url.Parse(c.baseURL)
		if err != nil {
			return fmt.Errorf("invalid base URL: %w", err)
		}
		u.Path = path
		u.RawQuery = query.Encode()

		// 5. 创建HTTP请求
		httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
		if err != nil {
			return fmt.Errorf("failed to create http request: %w", err)
		}

		// 6. 添加自定义headers（如Range）
		for key, value := range headers {
			httpReq.Header.Set(key, value)
		}

		// 7. 记录请求日志
		startTime := time.Now()
		c.logger.Debug("API Request (Media)", withTraceID(ctx,
			logger.F("method", httpReq.Method),
			logger.F("url", httpReq.URL.String()),
			logger.F("agent_key", agentKey))...)

		// 8. 发送请求
		httpResp, err := c.httpClient.Do(httpReq)
		if err != nil {
			duration := time.Since(startTime)
			c.logger.Error("Media request failed", withTraceID(ctx,
				logger.F("error", err),
				logger.F("duration", duration))...)
			return fmt.Errorf("http request failed: %w", err)
		}
		defer httpResp.Body.Close()

		duration := time.Since(startTime)

		// 9. 检查HTTP状态码
		if httpResp.StatusCode != http.StatusOK && httpResp.StatusCode != http.StatusPartialContent {
			// 尝试解析错误响应
			var errResp Response
			body, _ := io.ReadAll(httpResp.Body)
			if jsonErr := json.Unmarshal(body, &errResp); jsonErr == nil && errResp.ErrCode != 0 {
				c.logger.Error("Media request failed", withTraceID(ctx,
					logger.F("url", httpReq.URL.String()),
					logger.F("errcode", errResp.ErrCode),
					logger.F("errmsg", errResp.ErrMsg),
					logger.F("duration", duration))...)

				// Token 失效，刷新后重试
				apiErr := errors.New(errResp.ErrCode, errResp.ErrMsg)
				if errors.IsTokenExpired(apiErr) {
					c.logger.Warn("Token expired, refreshing", withTraceID(ctx,
						logger.F("errcode", errResp.ErrCode),
						logger.F("agent_key", agentKey))...)
					if refreshErr := c.tokenManager.RefreshTokenByAgent(ctx, agentKey); refreshErr != nil {
						c.logger.Error("Failed to refresh token", withTraceID(ctx,
							logger.F("error", refreshErr),
							logger.F("agent_key", agentKey))...)
					}
				}

				return apiErr
			}
			return fmt.Errorf("unexpected status code: %d", httpResp.StatusCode)
		}

		// 10. 读取响应体
		result, err = io.ReadAll(httpResp.Body)
		if err != nil {
			c.logger.Error("Failed to read media response", withTraceID(ctx,
				logger.F("error", err),
				logger.F("duration", duration))...)
			return fmt.Errorf("failed to read response body: %w", err)
		}

		// 11. 记录成功日志
		c.logger.Info("Media request successful", withTraceID(ctx,
			logger.F("url", httpReq.URL.String()),
			logger.F("size", len(result)),
			logger.F("duration", duration))...)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
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
