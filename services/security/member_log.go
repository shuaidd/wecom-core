package security

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/security"
)

// ListMemberOperLog 获取成员操作记录
// 文档: https://developer.work.weixin.qq.com/document/path/101710
func (s *Service) ListMemberOperLog(ctx context.Context, req *security.ListMemberOperLogRequest) (*security.ListMemberOperLogResponse, error) {
	return client.PostAndUnmarshal[security.ListMemberOperLogResponse](s.client, ctx, "/cgi-bin/security/member_oper_log/list", req)
}
