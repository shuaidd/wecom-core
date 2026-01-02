package updown

import (
	"github.com/shuaidd/wecom-core/internal/client"
)

// Service 上下游服务
type Service struct {
	client *client.Client
}

// NewService 创建上下游服务
func NewService(c *client.Client) *Service {
	return &Service{
		client: c,
	}
}
