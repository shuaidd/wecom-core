package message

import (
	"github.com/shuaidd/wecom-core/internal/client"
)

// Service 消息服务
type Service struct {
	client *client.Client
}

// NewService 创建消息服务
func NewService(c *client.Client) *Service {
	return &Service{
		client: c,
	}
}
