package kf

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/kf"
)

// SendMsg 发送消息
// 当微信客户处于"新接入待处理"或"由智能助手接待"状态下,可调用该接口给用户发送消息
// 注意仅当微信客户在主动发送消息给客服后的48小时内,企业可发送消息给客户,最多可发送5条消息
// 支持发送消息类型:文本、图片、语音、视频、文件、图文、小程序、菜单消息、地理位置、获客链接
// 文档: https://developer.work.weixin.qq.com/document/path/94677
func (s *Service) SendMsg(ctx context.Context, req *kf.SendMsgRequest) (*kf.SendMsgResponse, error) {
	return client.PostAndUnmarshal[kf.SendMsgResponse](s.client, ctx, "/cgi-bin/kf/send_msg", req)
}

// SendMsgOnEvent 发送欢迎语等事件响应消息
// 当特定的事件回调消息包含code字段,或通过接口变更到特定的会话状态,会返回code字段
// 开发者可以此code为凭证,调用该接口给用户发送相应事件场景下的消息
// 文档: https://developer.work.weixin.qq.com/document/path/94698
func (s *Service) SendMsgOnEvent(ctx context.Context, req *kf.SendMsgOnEventRequest) (*kf.SendMsgOnEventResponse, error) {
	return client.PostAndUnmarshal[kf.SendMsgOnEventResponse](s.client, ctx, "/cgi-bin/kf/send_msg_on_event", req)
}
