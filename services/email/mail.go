package email

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/email"
)

// GetNewCount 获取邮件未读数
// 获取指定成员邮箱当前未读邮件数量
//
// 文档: https://developer.work.weixin.qq.com/document/path/95845
func (s *Service) GetNewCount(ctx context.Context, req *email.GetNewCountRequest) (*email.GetNewCountResponse, error) {
	return client.PostAndUnmarshal[email.GetNewCountResponse](s.client, ctx, "/cgi-bin/exmail/mail/get_newcount", req)
}

// GetMailList 获取收件箱邮件列表
// 由于用户可向应用所绑定的邮箱回复邮件，故支持应用获取应用收件箱的邮件列表。
// 该接口用于分页获取收件箱下，指定时间范围内的邮件id列表。
//
// 文档: https://developer.work.weixin.qq.com/document/path/95846
func (s *Service) GetMailList(ctx context.Context, req *email.GetMailListRequest) (*email.GetMailListResponse, error) {
	return client.PostAndUnmarshal[email.GetMailListResponse](s.client, ctx, "/cgi-bin/exmail/app/get_mail_list", req)
}

// ReadMail 获取邮件内容
// 支持应用获取邮件内容。指定单个邮件id，获取邮件eml数据。
//
// 文档: https://developer.work.weixin.qq.com/document/path/95847
func (s *Service) ReadMail(ctx context.Context, req *email.ReadMailRequest) (*email.ReadMailResponse, error) {
	return client.PostAndUnmarshal[email.ReadMailResponse](s.client, ctx, "/cgi-bin/exmail/app/read_mail", req)
}
