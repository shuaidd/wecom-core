package security

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/security"
)

// SubmitBatchAddVIPJob 批量分配高级功能账号
// 文档: https://developer.work.weixin.qq.com/document/path/95946
func (s *Service) SubmitBatchAddVIPJob(ctx context.Context, req *security.SubmitBatchAddVIPJobRequest) (*security.SubmitBatchAddVIPJobResponse, error) {
	return client.PostAndUnmarshal[security.SubmitBatchAddVIPJobResponse](s.client, ctx, "/cgi-bin/security/vip/submit_batch_add_job", req)
}

// BatchAddVIPJobResult 查询分配高级功能账号结果
// 文档: https://developer.work.weixin.qq.com/document/path/95946
func (s *Service) BatchAddVIPJobResult(ctx context.Context, req *security.BatchAddVIPJobResultRequest) (*security.BatchAddVIPJobResultResponse, error) {
	return client.PostAndUnmarshal[security.BatchAddVIPJobResultResponse](s.client, ctx, "/cgi-bin/security/vip/batch_add_job_result", req)
}

// SubmitBatchDelVIPJob 批量取消高级功能账号
// 文档: https://developer.work.weixin.qq.com/document/path/95947
func (s *Service) SubmitBatchDelVIPJob(ctx context.Context, req *security.SubmitBatchDelVIPJobRequest) (*security.SubmitBatchDelVIPJobResponse, error) {
	return client.PostAndUnmarshal[security.SubmitBatchDelVIPJobResponse](s.client, ctx, "/cgi-bin/security/vip/submit_batch_del_job", req)
}

// BatchDelVIPJobResult 查询取消高级功能账号结果
// 文档: https://developer.work.weixin.qq.com/document/path/95947
func (s *Service) BatchDelVIPJobResult(ctx context.Context, req *security.BatchDelVIPJobResultRequest) (*security.BatchDelVIPJobResultResponse, error) {
	return client.PostAndUnmarshal[security.BatchDelVIPJobResultResponse](s.client, ctx, "/cgi-bin/security/vip/batch_del_job_result", req)
}

// ListVIP 获取高级功能账号列表
// 文档: https://developer.work.weixin.qq.com/document/path/95945
func (s *Service) ListVIP(ctx context.Context, req *security.ListVIPRequest) (*security.ListVIPResponse, error) {
	return client.PostAndUnmarshal[security.ListVIPResponse](s.client, ctx, "/cgi-bin/security/vip/list", req)
}
