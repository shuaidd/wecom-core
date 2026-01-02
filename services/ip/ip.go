package ip

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/ip"
)

// Service IP相关服务
type Service struct {
	client *client.Client
}

// NewService 创建IP服务
func NewService(c *client.Client) *Service {
	return &Service{
		client: c,
	}
}

// GetCallbackIP 获取企业微信回调IP段
// 文档: https://developer.work.weixin.qq.com/document/path/90930
func (s *Service) GetCallbackIP(ctx context.Context) ([]string, error) {
	result, err := client.GetAndUnmarshal[ip.GetCallbackIPResponse](s.client, ctx, "/cgi-bin/getcallbackip", nil)
	if err != nil {
		return nil, err
	}

	return result.IPList, nil
}

// GetAPIDomainIP 获取企业微信接口IP段
// 文档: https://developer.work.weixin.qq.com/document/path/92520
func (s *Service) GetAPIDomainIP(ctx context.Context) ([]string, error) {
	result, err := client.GetAndUnmarshal[ip.GetAPIDomainIPResponse](s.client, ctx, "/cgi-bin/get_api_domain_ip", nil)
	if err != nil {
		return nil, err
	}

	return result.IPList, nil
}
