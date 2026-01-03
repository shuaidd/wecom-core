package externalcontact

import (
	"github.com/shuaidd/wecom-core/internal/client"
)

// Service 外部联系人服务
type Service struct {
	client *client.Client
}

// NewService 创建外部联系人服务
func NewService(c *client.Client) *Service {
	return &Service{
		client: c,
	}
}
