package config

import (
	"fmt"
	"time"

	"github.com/shuaidd/wecom-core/pkg/cache"
	"github.com/shuaidd/wecom-core/pkg/interceptor"
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

// WithAgent 添加单个应用配置
func WithAgent(agentName string, agentID int64, secret string, agentDesc ...string) Option {
	return func(c *Config) {
		if c.Agents == nil {
			c.Agents = make(map[string]*AgentConfig)
		}
		agent := &AgentConfig{
			AgentID:   agentID,
			Secret:    secret,
			AgentName: agentName,
		}
		if len(agentDesc) > 0 {
			agent.AgentDesc = agentDesc[0]
		}
		// 同时支持通过名称和ID查找
		c.Agents[agentName] = agent
		c.Agents[fmt.Sprintf("%d", agentID)] = agent
	}
}

// WithAgents 批量添加应用配置
func WithAgents(agents ...*AgentConfig) Option {
	return func(c *Config) {
		if c.Agents == nil {
			c.Agents = make(map[string]*AgentConfig)
		}
		for _, agent := range agents {
			// 同时支持通过名称和ID查找
			if agent.AgentName != "" {
				c.Agents[agent.AgentName] = agent
			}
			if agent.AgentID > 0 {
				c.Agents[fmt.Sprintf("%d", agent.AgentID)] = agent
			}
		}
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

// WithRequestInterceptor 添加请求拦截器
func WithRequestInterceptor(interceptors ...interceptor.RequestInterceptor) Option {
	return func(c *Config) {
		c.RequestInterceptors = append(c.RequestInterceptors, interceptors...)
	}
}

// WithResponseInterceptor 添加响应拦截器（解析前）
func WithResponseInterceptor(interceptors ...interceptor.ResponseInterceptor) Option {
	return func(c *Config) {
		c.ResponseInterceptors = append(c.ResponseInterceptors, interceptors...)
	}
}

// WithAfterResponseInterceptor 添加响应后拦截器（解析后）
func WithAfterResponseInterceptor(interceptors ...interceptor.AfterResponseInterceptor) Option {
	return func(c *Config) {
		c.AfterResponseInterceptors = append(c.AfterResponseInterceptors, interceptors...)
	}
}
