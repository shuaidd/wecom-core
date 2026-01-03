package externalcontact

import (
	"context"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/externalcontact"
)

// ListAcquisitionLink 获取获客链接列表
// 企业可通过此接口获取当前仍然有效且是当前应用创建的获客链接
// 文档: https://developer.work.weixin.qq.com/document/path/97297
func (s *Service) ListAcquisitionLink(ctx context.Context, req *externalcontact.ListAcquisitionLinkRequest) (*externalcontact.ListAcquisitionLinkResponse, error) {
	return client.PostAndUnmarshal[externalcontact.ListAcquisitionLinkResponse](s.client, ctx, "/cgi-bin/externalcontact/customer_acquisition/list_link", req)
}

// GetAcquisitionLink 获取获客链接详情
// 企业可通过此接口根据获客链接id获取链接配置详情
// 文档: https://developer.work.weixin.qq.com/document/path/97297
func (s *Service) GetAcquisitionLink(ctx context.Context, linkID string) (*externalcontact.GetAcquisitionLinkResponse, error) {
	req := &externalcontact.GetAcquisitionLinkRequest{
		LinkID: linkID,
	}
	return client.PostAndUnmarshal[externalcontact.GetAcquisitionLinkResponse](s.client, ctx, "/cgi-bin/externalcontact/customer_acquisition/get", req)
}

// CreateAcquisitionLink 创建获客链接
// 企业可通过此接口创建新的获客链接
// 文档: https://developer.work.weixin.qq.com/document/path/97297
func (s *Service) CreateAcquisitionLink(ctx context.Context, req *externalcontact.CreateAcquisitionLinkRequest) (*externalcontact.CreateAcquisitionLinkResponse, error) {
	return client.PostAndUnmarshal[externalcontact.CreateAcquisitionLinkResponse](s.client, ctx, "/cgi-bin/externalcontact/customer_acquisition/create_link", req)
}

// UpdateAcquisitionLink 编辑获客链接
// 企业可通过此接口编辑获客链接，修改获客链接的关联范围或修改获客链接的名称
// 文档: https://developer.work.weixin.qq.com/document/path/97297
func (s *Service) UpdateAcquisitionLink(ctx context.Context, req *externalcontact.UpdateAcquisitionLinkRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/customer_acquisition/update_link", req)
	return err
}

// DeleteAcquisitionLink 删除获客链接
// 企业可通过此接口删除获客链接，删除后的获客链接将无法继续使用
// 文档: https://developer.work.weixin.qq.com/document/path/97297
func (s *Service) DeleteAcquisitionLink(ctx context.Context, linkID string) error {
	req := &externalcontact.DeleteAcquisitionLinkRequest{
		LinkID: linkID,
	}
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/customer_acquisition/delete_link", req)
	return err
}

// GetAcquisitionQuota 查询剩余使用量
// 企业可通过此接口查询当前剩余的使用量
// 文档: https://developer.work.weixin.qq.com/document/path/97375
func (s *Service) GetAcquisitionQuota(ctx context.Context) (*externalcontact.GetAcquisitionQuotaResponse, error) {
	return client.GetAndUnmarshal[externalcontact.GetAcquisitionQuotaResponse](s.client, ctx, "/cgi-bin/externalcontact/customer_acquisition_quota", url.Values{})
}

// GetAcquisitionStatistic 查询链接使用详情
// 企业可通过此接口查询指定获客链接在指定时间范围内的访问情况
// 文档: https://developer.work.weixin.qq.com/document/path/97375
func (s *Service) GetAcquisitionStatistic(ctx context.Context, req *externalcontact.GetAcquisitionStatisticRequest) (*externalcontact.GetAcquisitionStatisticResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetAcquisitionStatisticResponse](s.client, ctx, "/cgi-bin/externalcontact/customer_acquisition/statistic", req)
}

// GetAcquisitionChatInfo 获取成员多次收消息详情
// 企业和服务商可通过此接口获取成员多次收消息情况，如次数、客户id等信息
// 文档: https://developer.work.weixin.qq.com/document/path/97382
func (s *Service) GetAcquisitionChatInfo(ctx context.Context, chatKey string) (*externalcontact.GetAcquisitionChatInfoResponse, error) {
	req := &externalcontact.GetAcquisitionChatInfoRequest{
		ChatKey: chatKey,
	}
	return client.PostAndUnmarshal[externalcontact.GetAcquisitionChatInfoResponse](s.client, ctx, "/cgi-bin/externalcontact/customer_acquisition/get_chat_info", req)
}

// ListAcquisitionCustomer 获取获客客户列表
// 企业可通过此接口获取到由指定的获客链接添加的客户列表
// 文档: https://developer.work.weixin.qq.com/document/path/97443
func (s *Service) ListAcquisitionCustomer(ctx context.Context, req *externalcontact.ListAcquisitionCustomerRequest) (*externalcontact.ListAcquisitionCustomerResponse, error) {
	return client.PostAndUnmarshal[externalcontact.ListAcquisitionCustomerResponse](s.client, ctx, "/cgi-bin/externalcontact/customer_acquisition/customer", req)
}
