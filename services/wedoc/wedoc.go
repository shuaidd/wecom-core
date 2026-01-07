package wedoc

import (
	"github.com/shuaidd/wecom-core/internal/client"
)

// Service 微文档服务
type Service struct {
	client *client.Client
}

// New 创建微文档服务实例
func New(c *client.Client) *Service {
	return &Service{client: c}
}
