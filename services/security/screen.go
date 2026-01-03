package security

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/security"
)

// GetScreenOperRecord 获取截屏/录屏操作记录
// 文档: https://developer.work.weixin.qq.com/document/path/95632
func (s *Service) GetScreenOperRecord(ctx context.Context, req *security.GetScreenOperRecordRequest) (*security.GetScreenOperRecordResponse, error) {
	return client.PostAndUnmarshal[security.GetScreenOperRecordResponse](s.client, ctx, "/cgi-bin/security/get_screen_oper_record", req)
}
