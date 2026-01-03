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

// AddJoinWay 配置客户群进群方式
// 企业可调用此接口来生成并配置「加入群聊」的二维码或者小程序按钮，客户通过扫描二维码或点击小程序上的按钮，即可加入特定的客户群
// 文档: https://developer.work.weixin.qq.com/document/path/92229
func (s *Service) AddJoinWay(ctx context.Context, req *externalcontact.AddJoinWayRequest) (*externalcontact.AddJoinWayResponse, error) {
	return client.PostAndUnmarshal[externalcontact.AddJoinWayResponse](s.client, ctx, "/cgi-bin/externalcontact/groupchat/add_join_way", req)
}

// GetJoinWay 获取客户群进群方式配置
// 获取企业配置的群二维码或小程序按钮
// 文档: https://developer.work.weixin.qq.com/document/path/92229
func (s *Service) GetJoinWay(ctx context.Context, configID string) (*externalcontact.GetJoinWayResponse, error) {
	req := &externalcontact.GetJoinWayRequest{
		ConfigID: configID,
	}
	return client.PostAndUnmarshal[externalcontact.GetJoinWayResponse](s.client, ctx, "/cgi-bin/externalcontact/groupchat/get_join_way", req)
}

// UpdateJoinWay 更新客户群进群方式配置
// 更新进群方式配置信息。注意：使用覆盖的方式更新
// 文档: https://developer.work.weixin.qq.com/document/path/92229
func (s *Service) UpdateJoinWay(ctx context.Context, req *externalcontact.UpdateJoinWayRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/groupchat/update_join_way", req)
	return err
}

// DeleteJoinWay 删除客户群进群方式配置
// 删除一个进群方式配置
// 文档: https://developer.work.weixin.qq.com/document/path/92229
func (s *Service) DeleteJoinWay(ctx context.Context, configID string) error {
	req := &externalcontact.DeleteJoinWayRequest{
		ConfigID: configID,
	}
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/groupchat/del_join_way", req)
	return err
}
