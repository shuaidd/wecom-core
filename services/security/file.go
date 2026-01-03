package security

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/security"
)

// GetFileOperRecord 获取文件操作记录
// 文档: https://developer.work.weixin.qq.com/document/path/95435
func (s *Service) GetFileOperRecord(ctx context.Context, req *security.GetFileOperRecordRequest) (*security.GetFileOperRecordResponse, error) {
	return client.PostAndUnmarshal[security.GetFileOperRecordResponse](s.client, ctx, "/cgi-bin/security/get_file_oper_record", req)
}
