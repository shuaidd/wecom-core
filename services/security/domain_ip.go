package security

import (
	"context"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/security"
)

// GetServerDomainIP 获取企业微信域名IP信息
// 文档: https://developer.work.weixin.qq.com/document/path/97084
func (s *Service) GetServerDomainIP(ctx context.Context) (*security.GetServerDomainIPResponse, error) {
	query := url.Values{}
	return client.GetAndUnmarshal[security.GetServerDomainIPResponse](s.client, ctx, "/cgi-bin/security/get_server_domain_ip", query)
}
