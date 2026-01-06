package email

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/email"
)

// GetUserOption 获取用户功能属性
// 该接口用于获取用户的功能属性
//
// 文档: https://developer.work.weixin.qq.com/document/path/95851
func (s *Service) GetUserOption(ctx context.Context, req *email.GetUserOptionRequest) (*email.GetUserOptionResponse, error) {
	return client.PostAndUnmarshal[email.GetUserOptionResponse](s.client, ctx, "/cgi-bin/exmail/useroption/get", req)
}

// UpdateUserOption 更改用户功能属性
// 该接口用于更新用户的功能属性
//
// 文档: https://developer.work.weixin.qq.com/document/path/95852
func (s *Service) UpdateUserOption(ctx context.Context, req *email.UpdateUserOptionRequest) (*email.UpdateUserOptionResponse, error) {
	return client.PostAndUnmarshal[email.UpdateUserOptionResponse](s.client, ctx, "/cgi-bin/exmail/useroption/update", req)
}
