package config

import (
	"time"

	"github.com/shuaidd/wecom-core/pkg/cache"
	"github.com/shuaidd/wecom-core/pkg/logger"
)

// Config 企业微信SDK配置
type Config struct {
	// CorpID 企业ID
	CorpID string

	// CorpSecret 应用凭证密钥
	CorpSecret string

	// BaseURL API基础URL，默认为 https://qyapi.weixin.qq.com
	BaseURL string

	// Timeout HTTP请求超时时间，默认为 30 秒
	Timeout time.Duration

	// MaxRetries 最大重试次数，默认为 3 次
	MaxRetries int

	// InitialBackoff 初始退避时间，默认为 1 秒
	InitialBackoff time.Duration

	// MaxBackoff 最大退避时间，默认为 30 秒
	MaxBackoff time.Duration

	// Logger 日志记录器，默认为 NoopLogger
	Logger logger.Logger

	// Cache Token缓存，默认为内存缓存
	Cache cache.Cache
}

// New 创建配置对象
func New(opts ...Option) *Config {
	// 设置默认值
	cfg := &Config{
		BaseURL:        "https://qyapi.weixin.qq.com",
		Timeout:        30 * time.Second,
		MaxRetries:     3,
		InitialBackoff: 1 * time.Second,
		MaxBackoff:     30 * time.Second,
		Logger:         logger.NewNoopLogger(),
		Cache:          nil, // 将在 internal/auth 中使用默认的内存缓存
	}

	// 应用选项
	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

// Validate 验证配置
func (c *Config) Validate() error {
	if c.CorpID == "" {
		return ErrMissingCorpID
	}
	if c.CorpSecret == "" {
		return ErrMissingCorpSecret
	}
	if c.Timeout <= 0 {
		return ErrInvalidTimeout
	}
	if c.MaxRetries < 0 {
		return ErrInvalidMaxRetries
	}
	return nil
}
