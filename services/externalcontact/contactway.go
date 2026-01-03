package externalcontact

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/externalcontact"
)

// AddContactWay 配置客户联系「联系我」方式
// 企业可通过此接口为具有客户联系功能的成员生成专属的「联系我」二维码或者「联系我」按钮
// 文档: https://developer.work.weixin.qq.com/document/path/92572
func (s *Service) AddContactWay(ctx context.Context, req *externalcontact.AddContactWayRequest) (*externalcontact.AddContactWayResponse, error) {
	return client.PostAndUnmarshal[externalcontact.AddContactWayResponse](s.client, ctx, "/cgi-bin/externalcontact/add_contact_way", req)
}

// GetContactWay 获取企业已配置的「联系我」方式
// 获取企业配置的「联系我」二维码和「联系我」小程序按钮
// 文档: https://developer.work.weixin.qq.com/document/path/92572
func (s *Service) GetContactWay(ctx context.Context, configID string) (*externalcontact.GetContactWayResponse, error) {
	req := &externalcontact.GetContactWayRequest{
		ConfigID: configID,
	}
	return client.PostAndUnmarshal[externalcontact.GetContactWayResponse](s.client, ctx, "/cgi-bin/externalcontact/get_contact_way", req)
}

// ListContactWay 获取企业已配置的「联系我」列表
// 获取企业配置的「联系我」二维码和「联系我」小程序插件列表。不包含临时会话
// 注意：该接口仅可获取2021年7月10日以后创建的「联系我」
// 文档: https://developer.work.weixin.qq.com/document/path/92572
func (s *Service) ListContactWay(ctx context.Context, req *externalcontact.ListContactWayRequest) (*externalcontact.ListContactWayResponse, error) {
	return client.PostAndUnmarshal[externalcontact.ListContactWayResponse](s.client, ctx, "/cgi-bin/externalcontact/list_contact_way", req)
}

// UpdateContactWay 更新企业已配置的「联系我」方式
// 更新企业配置的「联系我」二维码和「联系我」小程序按钮中的信息，如使用人员和备注等
// 文档: https://developer.work.weixin.qq.com/document/path/92572
func (s *Service) UpdateContactWay(ctx context.Context, req *externalcontact.UpdateContactWayRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/update_contact_way", req)
	return err
}

// DeleteContactWay 删除企业已配置的「联系我」方式
// 删除一个已配置的「联系我」二维码或者「联系我」小程序按钮
// 文档: https://developer.work.weixin.qq.com/document/path/92572
func (s *Service) DeleteContactWay(ctx context.Context, configID string) error {
	req := &externalcontact.DeleteContactWayRequest{
		ConfigID: configID,
	}
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/del_contact_way", req)
	return err
}

// CloseTempChat 结束临时会话
// 将指定的企业成员和客户之前的临时会话断开，断开前会自动下发已配置的结束语
// 文档: https://developer.work.weixin.qq.com/document/path/92572
func (s *Service) CloseTempChat(ctx context.Context, userID, externalUserID string) error {
	req := &externalcontact.CloseTempChatRequest{
		UserID:         userID,
		ExternalUserID: externalUserID,
	}
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/close_temp_chat", req)
	return err
}
