package contact

import (
	"github.com/shuaidd/wecom-core/internal/client"
)

// Service 通讯录服务
type Service struct {
	client *client.Client
}

// NewService 创建通讯录服务
func NewService(c *client.Client) *Service {
	return &Service{
		client: c,
	}
}
