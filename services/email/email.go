package email

import (
	"github.com/shuaidd/wecom-core/internal/client"
)

// Service 邮件服务
type Service struct {
	client *client.Client
}

// NewService 创建邮件服务
func NewService(c *client.Client) *Service {
	return &Service{
		client: c,
	}
}
