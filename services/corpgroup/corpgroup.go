package corpgroup

import (
	"github.com/shuaidd/wecom-core/internal/client"
)

// Service 企业互联服务
type Service struct {
	client *client.Client
}

// NewService 创建企业互联服务
func NewService(c *client.Client) *Service {
	return &Service{
		client: c,
	}
}
