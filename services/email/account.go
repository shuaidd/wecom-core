package email

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/email"
)

// GetAppEmailAlias 查询应用邮箱账号
// 应用调用此接口，可以查询自己的应用邮箱账号及别名邮箱
//
// 文档: https://developer.work.weixin.qq.com/document/path/95843
func (s *Service) GetAppEmailAlias(ctx context.Context) (*email.GetAppEmailAliasResponse, error) {
	return client.PostAndUnmarshal[email.GetAppEmailAliasResponse](s.client, ctx, "/cgi-bin/exmail/app/get_email_alias", nil)
}

// UpdateAppEmailAlias 更新应用邮箱账号
// 应用调用此接口，可以更新应用邮箱账号，原有的应用邮箱账号将会作为别名邮箱，具有收信能力
//
// 文档: https://developer.work.weixin.qq.com/document/path/95844
func (s *Service) UpdateAppEmailAlias(ctx context.Context, req *email.UpdateAppEmailAliasRequest) (*email.UpdateAppEmailAliasResponse, error) {
	return client.PostAndUnmarshal[email.UpdateAppEmailAliasResponse](s.client, ctx, "/cgi-bin/exmail/app/update_email_alias", req)
}

// ActEmail 禁用/启用邮箱账号
// 禁用或启用邮箱，可以操作成员邮箱和业务邮箱
//
// 文档: https://developer.work.weixin.qq.com/document/path/95853
func (s *Service) ActEmail(ctx context.Context, req *email.ActEmailRequest) (*email.ActEmailResponse, error) {
	return client.PostAndUnmarshal[email.ActEmailResponse](s.client, ctx, "/cgi-bin/exmail/account/act_email", req)
}
