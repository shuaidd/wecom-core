package externalcontact

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/externalcontact"
)

// OnJobTransferCustomer 分配在职成员的客户
// 企业可通过此接口，转接在职成员的客户给其他成员
// 文档: https://developer.work.weixin.qq.com/document/path/94081
func (s *Service) OnJobTransferCustomer(ctx context.Context, req *externalcontact.OnJobTransferCustomerRequest) (*externalcontact.OnJobTransferCustomerResponse, error) {
	return client.PostAndUnmarshal[externalcontact.OnJobTransferCustomerResponse](s.client, ctx, "/cgi-bin/externalcontact/transfer_customer", req)
}

// OnJobTransferGroupChat 分配在职成员的客户群
// 企业可通过此接口，将在职成员为群主的群，分配给另一个客服成员
// 文档: https://developer.work.weixin.qq.com/document/path/95703
func (s *Service) OnJobTransferGroupChat(ctx context.Context, req *externalcontact.OnJobTransferGroupChatRequest) (*externalcontact.OnJobTransferGroupChatResponse, error) {
	return client.PostAndUnmarshal[externalcontact.OnJobTransferGroupChatResponse](s.client, ctx, "/cgi-bin/externalcontact/groupchat/onjob_transfer", req)
}

// GetTransferResult 查询客户接替状态
// 企业和第三方可通过此接口查询在职成员的客户转接情况
// 文档: https://developer.work.weixin.qq.com/document/path/94082
func (s *Service) GetTransferResult(ctx context.Context, req *externalcontact.TransferResultRequest) (*externalcontact.TransferResultResponse, error) {
	return client.PostAndUnmarshal[externalcontact.TransferResultResponse](s.client, ctx, "/cgi-bin/externalcontact/transfer_result", req)
}
