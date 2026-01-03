package message

import (
	"context"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/common"
	"github.com/shuaidd/wecom-core/types/message"
)

// SendAppChat 应用推送消息
// 文档: https://developer.work.weixin.qq.com/document/path/90248
func (s *Service) SendAppChat(ctx context.Context, req *message.AppChatSendRequest) (*common.Response, error) {
	return client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/appchat/send", req)
}

// GetAppChat 获取群聊会话
// 文档: https://developer.work.weixin.qq.com/document/path/90247
func (s *Service) GetAppChat(ctx context.Context, chatID string) (*message.GetAppChatResponse, error) {
	query := url.Values{}
	query.Set("chatid", chatID)
	return client.GetAndUnmarshal[message.GetAppChatResponse](s.client, ctx, "/cgi-bin/appchat/get", query)
}

// UpdateAppChat 修改群聊会话
// 文档: https://developer.work.weixin.qq.com/document/path/90246
func (s *Service) UpdateAppChat(ctx context.Context, req *message.UpdateAppChatRequest) (*common.Response, error) {
	return client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/appchat/update", req)
}

// ListSmartsheetGroupChat 获取群聊列表
// 该接口可以查询通过智能表格自动化创建的群聊列表
// 文档: https://developer.work.weixin.qq.com/document/path/98149
func (s *Service) ListSmartsheetGroupChat(ctx context.Context, req *message.ListSmartsheetGroupChatRequest) (*message.ListSmartsheetGroupChatResponse, error) {
	return client.PostAndUnmarshal[message.ListSmartsheetGroupChatResponse](s.client, ctx, "/cgi-bin/wedoc/smartsheet/groupchat/list", req)
}
