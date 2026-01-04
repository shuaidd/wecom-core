package invoice

import (
	"github.com/shuaidd/wecom-core/internal/client"
)

// Service 电子发票服务
type Service struct {
	client *client.Client
}

// NewService 创建电子发票服务
func NewService(c *client.Client) *Service {
	return &Service{
		client: c,
	}
}
