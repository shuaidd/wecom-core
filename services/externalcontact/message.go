package externalcontact

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/externalcontact"
)

// AddMsgTemplate 创建企业群发
// 企业跟第三方应用可通过此接口添加企业群发消息的任务并通知成员发送给相关客户或客户群
// 文档: https://developer.work.weixin.qq.com/document/path/92135
func (s *Service) AddMsgTemplate(ctx context.Context, req *externalcontact.AddMsgTemplateRequest) (*externalcontact.AddMsgTemplateResponse, error) {
	return client.PostAndUnmarshal[externalcontact.AddMsgTemplateResponse](s.client, ctx, "/cgi-bin/externalcontact/add_msg_template", req)
}

// GetGroupMsgListV2 获取群发记录列表
// 企业和第三方应用可通过此接口获取企业与成员的群发记录
// 文档: https://developer.work.weixin.qq.com/document/path/93338
func (s *Service) GetGroupMsgListV2(ctx context.Context, req *externalcontact.GetGroupMsgListV2Request) (*externalcontact.GetGroupMsgListV2Response, error) {
	return client.PostAndUnmarshal[externalcontact.GetGroupMsgListV2Response](s.client, ctx, "/cgi-bin/externalcontact/get_groupmsg_list_v2", req)
}

// GetGroupMsgTask 获取群发成员发送任务列表
// 企业和第三方应用可通过此接口获取群发成员发送任务列表
// 文档: https://developer.work.weixin.qq.com/document/path/93338
func (s *Service) GetGroupMsgTask(ctx context.Context, req *externalcontact.GetGroupMsgTaskRequest) (*externalcontact.GetGroupMsgTaskResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetGroupMsgTaskResponse](s.client, ctx, "/cgi-bin/externalcontact/get_groupmsg_task", req)
}

// GetGroupMsgSendResult 获取企业群发成员执行结果
// 企业和第三方应用可通过此接口获取企业群发成员执行结果
// 文档: https://developer.work.weixin.qq.com/document/path/93338
func (s *Service) GetGroupMsgSendResult(ctx context.Context, req *externalcontact.GetGroupMsgSendResultRequest) (*externalcontact.GetGroupMsgSendResultResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetGroupMsgSendResultResponse](s.client, ctx, "/cgi-bin/externalcontact/get_groupmsg_send_result", req)
}

// SendWelcomeMsg 发送新客户欢迎语
// 企业微信在向企业推送添加外部联系人事件时，会额外返回一个welcome_code，
// 企业以此为凭据调用接口，即可通过成员向新添加的客户发送个性化的欢迎语
// 文档: https://developer.work.weixin.qq.com/document/path/92137
func (s *Service) SendWelcomeMsg(ctx context.Context, req *externalcontact.SendWelcomeMsgRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/send_welcome_msg", req)
	return err
}

// CancelGroupMsgSend 停止企业群发
// 企业和第三方应用可调用此接口，停止无需成员继续发送的企业群发
// 文档: https://developer.work.weixin.qq.com/document/path/93341
func (s *Service) CancelGroupMsgSend(ctx context.Context, req *externalcontact.CancelGroupMsgSendRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/cancel_groupmsg_send", req)
	return err
}

// RemindGroupMsgSend 提醒成员群发
// 企业和第三方应用可调用此接口，重新触发群发通知，提醒成员完成群发任务
// 24小时内每个群发最多触发三次提醒
// 文档: https://developer.work.weixin.qq.com/document/path/93340
func (s *Service) RemindGroupMsgSend(ctx context.Context, req *externalcontact.RemindGroupMsgSendRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/remind_groupmsg_send", req)
	return err
}

// AddGroupWelcomeTemplate 添加入群欢迎语素材
// 企业可通过此API向企业的入群欢迎语素材库中添加素材
// 每个企业的入群欢迎语素材库中，最多容纳100个素材
// 文档: https://developer.work.weixin.qq.com/document/path/92366
func (s *Service) AddGroupWelcomeTemplate(ctx context.Context, req *externalcontact.AddGroupWelcomeTemplateRequest) (*externalcontact.AddGroupWelcomeTemplateResponse, error) {
	return client.PostAndUnmarshal[externalcontact.AddGroupWelcomeTemplateResponse](s.client, ctx, "/cgi-bin/externalcontact/group_welcome_template/add", req)
}

// EditGroupWelcomeTemplate 编辑入群欢迎语素材
// 企业可通过此API编辑入群欢迎语素材库中的素材，且仅能够编辑调用方自己创建的入群欢迎语素材
// 文档: https://developer.work.weixin.qq.com/document/path/92366
func (s *Service) EditGroupWelcomeTemplate(ctx context.Context, req *externalcontact.EditGroupWelcomeTemplateRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/group_welcome_template/edit", req)
	return err
}

// GetGroupWelcomeTemplate 获取入群欢迎语素材
// 企业可通过此API获取入群欢迎语素材
// 文档: https://developer.work.weixin.qq.com/document/path/92366
func (s *Service) GetGroupWelcomeTemplate(ctx context.Context, req *externalcontact.GetGroupWelcomeTemplateRequest) (*externalcontact.GetGroupWelcomeTemplateResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetGroupWelcomeTemplateResponse](s.client, ctx, "/cgi-bin/externalcontact/group_welcome_template/get", req)
}

// DelGroupWelcomeTemplate 删除入群欢迎语素材
// 企业可通过此API删除入群欢迎语素材，且仅能删除调用方自己创建的入群欢迎语素材
// 文档: https://developer.work.weixin.qq.com/document/path/92366
func (s *Service) DelGroupWelcomeTemplate(ctx context.Context, req *externalcontact.DelGroupWelcomeTemplateRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/group_welcome_template/del", req)
	return err
}
