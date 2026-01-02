package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/shuaidd/wecom-core/internal/auth"
	"github.com/shuaidd/wecom-core/internal/errors"
	"github.com/shuaidd/wecom-core/internal/retry"
	"github.com/shuaidd/wecom-core/pkg/logger"
)

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
	}
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
		c.logger.Debug("API Request",
			logger.F("method", httpReq.Method),
			logger.F("url", httpReq.URL.String()))

		// 5. 发送请求
		httpResp, err := c.httpClient.Do(httpReq)
		if err != nil {
			duration := time.Since(startTime)
			c.logger.Error("Request failed",
				logger.F("error", err),
				logger.F("duration", duration))
			return fmt.Errorf("http request failed: %w", err)
		}
		defer httpResp.Body.Close()

		// 6. 解析响应
		resp, err = ParseResponse(httpResp)
		duration := time.Since(startTime)

		if err != nil {
			c.logger.Error("Request failed",
				logger.F("url", httpReq.URL.String()),
				logger.F("errcode", resp.ErrCode),
				logger.F("errmsg", resp.ErrMsg),
				logger.F("duration", duration))

			// 7. Token 失效，刷新后重试
			if errors.IsTokenExpired(err) {
				c.logger.Warn("Token expired, refreshing",
					logger.F("errcode", resp.ErrCode))
				if refreshErr := c.tokenManager.RefreshToken(ctx); refreshErr != nil {
					c.logger.Error("Failed to refresh token",
						logger.F("error", refreshErr))
				}
			}

			return err
		}

		// 8. 记录成功日志
		c.logger.Info("API Request successful",
			logger.F("url", httpReq.URL.String()),
			logger.F("duration", duration))

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
func (c *Client) Post(ctx context.Context, path string, body interface{}) (*Response, error) {
	req := NewRequest(MethodPost, path).SetBody(body)
	return c.Do(ctx, req)
}
