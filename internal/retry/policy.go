package retry

import (
	"time"
)

// Policy 重试策略
type Policy struct {
	// MaxRetries 最大重试次数
	MaxRetries int
	// InitialBackoff 初始退避时间
	InitialBackoff time.Duration
	// MaxBackoff 最大退避时间
	MaxBackoff time.Duration
}

// NewPolicy 创建重试策略
func NewPolicy(maxRetries int, initialBackoff, maxBackoff time.Duration) *Policy {
	return &Policy{
		MaxRetries:     maxRetries,
		InitialBackoff: initialBackoff,
		MaxBackoff:     maxBackoff,
	}
}

// Backoff 计算退避时间（指数退避算法）
// backoff = InitialBackoff * 2^attempt
func (p *Policy) Backoff(attempt int) time.Duration {
	backoff := p.InitialBackoff * (1 << attempt)
	if backoff > p.MaxBackoff {
		return p.MaxBackoff
	}
	return backoff
}
