package externalcontact

import (
	"context"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/externalcontact"
)

// ListExternalContact 获取客户列表
// 企业可通过此接口获取指定成员添加的客户列表
// 文档: https://developer.work.weixin.qq.com/document/path/92113
func (s *Service) ListExternalContact(ctx context.Context, userID string) (*externalcontact.ListExternalContactResponse, error) {
	params := url.Values{}
	params.Set("userid", userID)
	return client.GetAndUnmarshal[externalcontact.ListExternalContactResponse](s.client, ctx, "/cgi-bin/externalcontact/list", params)
}

// GetExternalContact 获取客户详情
// 企业可通过此接口，根据外部联系人的userid，拉取客户详情
// 文档: https://developer.work.weixin.qq.com/document/path/92114
func (s *Service) GetExternalContact(ctx context.Context, externalUserID string, cursor ...string) (*externalcontact.GetExternalContactResponse, error) {
	params := url.Values{}
	params.Set("external_userid", externalUserID)
	if len(cursor) > 0 && cursor[0] != "" {
		params.Set("cursor", cursor[0])
	}
	return client.GetAndUnmarshal[externalcontact.GetExternalContactResponse](s.client, ctx, "/cgi-bin/externalcontact/get", params)
}

// UpdateRemark 修改客户备注信息
// 企业可通过此接口修改指定用户添加的客户的备注信息
// 文档: https://developer.work.weixin.qq.com/document/path/92115
func (s *Service) UpdateRemark(ctx context.Context, req *externalcontact.UpdateRemarkRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/remark", req)
	return err
}

// BatchGetByUser 批量获取客户详情
// 企业可通过此接口获取指定成员添加的客户信息列表
// 文档: https://developer.work.weixin.qq.com/document/path/92994
func (s *Service) BatchGetByUser(ctx context.Context, req *externalcontact.BatchGetByUserRequest) (*externalcontact.BatchGetByUserResponse, error) {
	return client.PostAndUnmarshal[externalcontact.BatchGetByUserResponse](s.client, ctx, "/cgi-bin/externalcontact/batch/get_by_user", req)
}
