package kf

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/kf"
)

// AddAccount 添加客服账号
// 文档: https://developer.work.weixin.qq.com/document/path/94254
func (s *Service) AddAccount(ctx context.Context, req *kf.AddAccountRequest) (*kf.AddAccountResponse, error) {
	return client.PostAndUnmarshal[kf.AddAccountResponse](s.client, ctx, "/cgi-bin/kf/account/add", req)
}

// DeleteAccount 删除客服账号
// 文档: https://developer.work.weixin.qq.com/document/path/94255
func (s *Service) DeleteAccount(ctx context.Context, req *kf.DeleteAccountRequest) error {
	_, err := client.PostAndUnmarshal[kf.DeleteAccountResponse](s.client, ctx, "/cgi-bin/kf/account/del", req)
	return err
}

// UpdateAccount 修改客服账号
// 文档: https://developer.work.weixin.qq.com/document/path/94256
func (s *Service) UpdateAccount(ctx context.Context, req *kf.UpdateAccountRequest) error {
	_, err := client.PostAndUnmarshal[kf.UpdateAccountResponse](s.client, ctx, "/cgi-bin/kf/account/update", req)
	return err
}

// ListAccount 获取客服账号列表
// 文档: https://developer.work.weixin.qq.com/document/path/94257
func (s *Service) ListAccount(ctx context.Context, req *kf.ListAccountRequest) (*kf.ListAccountResponse, error) {
	return client.PostAndUnmarshal[kf.ListAccountResponse](s.client, ctx, "/cgi-bin/kf/account/list", req)
}

// AddContactWay 获取客服账号链接
// 文档: https://developer.work.weixin.qq.com/document/path/94258
func (s *Service) AddContactWay(ctx context.Context, req *kf.AddContactWayRequest) (*kf.AddContactWayResponse, error) {
	return client.PostAndUnmarshal[kf.AddContactWayResponse](s.client, ctx, "/cgi-bin/kf/add_contact_way", req)
}
