package config

import (
	"time"

	"github.com/shuaidd/wecom-core/pkg/cache"
	"github.com/shuaidd/wecom-core/pkg/logger"
)

// Option 配置选项函数
type Option func(*Config)

// WithCorpID 设置企业ID
func WithCorpID(corpID string) Option {
	return func(c *Config) {
		c.CorpID = corpID
	}
}

// WithCorpSecret 设置应用凭证密钥
func WithCorpSecret(secret string) Option {
	return func(c *Config) {
		c.CorpSecret = secret
	}
}

// WithBaseURL 设置API基础URL
func WithBaseURL(baseURL string) Option {
	return func(c *Config) {
		c.BaseURL = baseURL
	}
}

// WithTimeout 设置HTTP请求超时时间
func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

// WithRetry 设置最大重试次数
func WithRetry(maxRetries int) Option {
	return func(c *Config) {
		c.MaxRetries = maxRetries
	}
}

// WithBackoff 设置退避时间
func WithBackoff(initial, max time.Duration) Option {
	return func(c *Config) {
		c.InitialBackoff = initial
		c.MaxBackoff = max
	}
}

// WithLogger 设置日志记录器
func WithLogger(logger logger.Logger) Option {
	return func(c *Config) {
		c.Logger = logger
	}
}

// WithCache 设置缓存
func WithCache(cache cache.Cache) Option {
	return func(c *Config) {
		c.Cache = cache
	}
}

// WithDebug 设置debug模式
func WithDebug(debug bool) Option {
	return func(c *Config) {
		c.Debug = debug
	}
}
