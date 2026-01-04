package kf

import (
	"github.com/shuaidd/wecom-core/internal/client"
)

// Service 微信客服服务
type Service struct {
	client *client.Client
}

// NewService 创建微信客服服务
func NewService(c *client.Client) *Service {
	return &Service{
		client: c,
	}
}
