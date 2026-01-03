package security

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/security"
)

// ListAdminOperLog 获取管理端操作日志
// 文档: https://developer.work.weixin.qq.com/document/path/101711
func (s *Service) ListAdminOperLog(ctx context.Context, req *security.ListAdminOperLogRequest) (*security.ListAdminOperLogResponse, error) {
	return client.PostAndUnmarshal[security.ListAdminOperLogResponse](s.client, ctx, "/cgi-bin/security/admin_oper_log/list", req)
}
