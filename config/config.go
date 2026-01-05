package config

import (
	"time"

	"github.com/shuaidd/wecom-core/pkg/cache"
	"github.com/shuaidd/wecom-core/pkg/interceptor"
	"github.com/shuaidd/wecom-core/pkg/logger"
)

// AgentConfig 应用配置
type AgentConfig struct {
	// AgentID 应用ID
	AgentID int64

	// Secret 应用凭证密钥
	Secret string

	// AgentName 应用名称(可选，用于通过名称查找应用)
	AgentName string

	// AgentDesc 应用描述(可选)
	AgentDesc string
}

// Config 企业微信SDK配置
type Config struct {
	// CorpID 企业ID
	CorpID string

	// CorpSecret 应用凭证密钥（向后兼容，如果只有一个应用可使用此字段）
	// Deprecated: 推荐使用 Agents 字段支持多应用
	CorpSecret string

	// Agents 多应用配置，key为应用名称或ID
	Agents map[string]*AgentConfig

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

	// RequestInterceptors 请求拦截器列表
	RequestInterceptors []interceptor.RequestInterceptor

	// ResponseInterceptors 响应拦截器列表（解析前）
	ResponseInterceptors []interceptor.ResponseInterceptor

	// AfterResponseInterceptors 响应后拦截器列表（解析后）
	AfterResponseInterceptors []interceptor.AfterResponseInterceptor

	// 是否开启debug模式
	Debug bool
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

	// 检查是否配置了应用凭证（支持单应用和多应用两种模式）
	if c.CorpSecret == "" && len(c.Agents) == 0 {
		return ErrMissingCorpSecret
	}

	// 如果配置了多个应用，验证每个应用的配置
	if len(c.Agents) > 0 {
		for key, agent := range c.Agents {
			if agent.Secret == "" {
				return &ErrInvalidAgentConfig{AgentKey: key, Reason: "secret is required"}
			}
		}
	}

	if c.Timeout <= 0 {
		return ErrInvalidTimeout
	}
	if c.MaxRetries < 0 {
		return ErrInvalidMaxRetries
	}
	return nil
}

// GetAgentByName 根据应用名称获取应用配置
func (c *Config) GetAgentByName(name string) *AgentConfig {
	if c.Agents == nil {
		return nil
	}
	return c.Agents[name]
}

// GetAgentByID 根据应用ID获取应用配置
func (c *Config) GetAgentByID(agentID int64) *AgentConfig {
	if c.Agents == nil {
		return nil
	}
	// 遍历查找匹配的 AgentID
	for _, agent := range c.Agents {
		if agent.AgentID == agentID {
			return agent
		}
	}
	return nil
}

// GetDefaultAgent 获取默认应用配置（用于向后兼容单应用模式）
func (c *Config) GetDefaultAgent() *AgentConfig {
	// 如果配置了 CorpSecret，返回默认应用
	if c.CorpSecret != "" {
		return &AgentConfig{
			Secret: c.CorpSecret,
		}
	}

	// 如果只有一个应用，返回该应用
	if len(c.Agents) == 1 {
		for _, agent := range c.Agents {
			return agent
		}
	}

	return nil
}
