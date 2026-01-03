package externalcontact

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/externalcontact"
)

// ListGroupChat 获取客户群列表
// 该接口用于获取配置过客户群管理的客户群列表
// 文档: https://developer.work.weixin.qq.com/document/path/92120
func (s *Service) ListGroupChat(ctx context.Context, req *externalcontact.ListGroupChatRequest) (*externalcontact.ListGroupChatResponse, error) {
	return client.PostAndUnmarshal[externalcontact.ListGroupChatResponse](s.client, ctx, "/cgi-bin/externalcontact/groupchat/list", req)
}

// GetGroupChat 获取客户群详情
// 通过客户群ID，获取详情。包括群名、群成员列表、群成员入群时间、入群方式
// 文档: https://developer.work.weixin.qq.com/document/path/92122
func (s *Service) GetGroupChat(ctx context.Context, req *externalcontact.GetGroupChatRequest) (*externalcontact.GetGroupChatResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetGroupChatResponse](s.client, ctx, "/cgi-bin/externalcontact/groupchat/get", req)
}

// OpenGIDToChatID 客户群opengid转换
// 用户在微信里的客户群里打开小程序时，某些场景下可以获取到群的opengid，
// 如果该群是企业微信的客户群，则企业或第三方可以调用此接口将一个opengid转换为客户群chat_id
// 文档: https://developer.work.weixin.qq.com/document/path/94822
func (s *Service) OpenGIDToChatID(ctx context.Context, req *externalcontact.OpenGIDToChatIDRequest) (*externalcontact.OpenGIDToChatIDResponse, error) {
	return client.PostAndUnmarshal[externalcontact.OpenGIDToChatIDResponse](s.client, ctx, "/cgi-bin/externalcontact/opengid_to_chatid", req)
}

// TransferGroupChat 分配离职成员的客户群给新群主
// 文档: https://developer.work.weixin.qq.com/document/path/93323
func (s *Service) TransferGroupChat(ctx context.Context, req *externalcontact.TransferGroupChatRequest) (*externalcontact.TransferGroupChatResponse, error) {
	return client.PostAndUnmarshal[externalcontact.TransferGroupChatResponse](s.client, ctx, "/cgi-bin/externalcontact/groupchat/transfer", req)
}
