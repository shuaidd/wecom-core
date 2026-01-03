package security

import (
	"github.com/shuaidd/wecom-core/internal/client"
)

// Service 安全管理服务
type Service struct {
	client *client.Client
}

// NewService 创建安全管理服务
func NewService(c *client.Client) *Service {
	return &Service{
		client: c,
	}
}
