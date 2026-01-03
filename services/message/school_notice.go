package message

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/message"
)

// SendSchoolNotice 发送学校通知
// 文档: https://developer.work.weixin.qq.com/document/path/90236#发送学校通知
func (s *Service) SendSchoolNotice(ctx context.Context, req *message.SendSchoolNoticeRequest) (*message.SendSchoolNoticeResponse, error) {
	return client.PostAndUnmarshal[message.SendSchoolNoticeResponse](s.client, ctx, "/cgi-bin/externalcontact/message/send", req)
}
