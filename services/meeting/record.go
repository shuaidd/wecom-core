package meeting

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/common"
	"github.com/shuaidd/wecom-core/types/meeting"
)

// UpdateSharingConfig 根据会议录制 ID 修改共享等配置
// 文档: docs/录制管理/修改会议录制共享设置.md
func (s *Service) UpdateSharingConfig(ctx context.Context, req *meeting.UpdateSharingConfigRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/meeting/record/update_sharing_config", req)
	return err
}

// SubmitBatchAddVIP 分配高级功能账号
// 文档: docs/录制管理/分配高级功能账号.md
func (s *Service) SubmitBatchAddVIP(ctx context.Context, req *meeting.SubmitBatchAddVIPRequest) (*meeting.SubmitBatchAddVIPResponse, error) {
	return client.PostAndUnmarshal[meeting.SubmitBatchAddVIPResponse](s.client, ctx, "/cgi-bin/meeting/vip/submit_batch_add_job", req)
}

// BatchAddJobResult 查询分配高级功能账号结果
// 文档: docs/录制管理/分配高级功能账号.md
func (s *Service) BatchAddJobResult(ctx context.Context, req *meeting.BatchAddJobResultRequest) (*meeting.BatchAddJobResultResponse, error) {
	return client.PostAndUnmarshal[meeting.BatchAddJobResultResponse](s.client, ctx, "/cgi-bin/meeting/vip/batch_add_job_result", req)
}

// SubmitBatchDelVIP 取消高级功能账号
// 文档: docs/录制管理/取消高级功能账号.md
func (s *Service) SubmitBatchDelVIP(ctx context.Context, req *meeting.SubmitBatchDelVIPRequest) (*meeting.SubmitBatchDelVIPResponse, error) {
	return client.PostAndUnmarshal[meeting.SubmitBatchDelVIPResponse](s.client, ctx, "/cgi-bin/meeting/vip/submit_batch_del_job", req)
}

// BatchDelJobResult 查询取消高级功能账号结果
// 文档: docs/录制管理/取消高级功能账号.md
func (s *Service) BatchDelJobResult(ctx context.Context, req *meeting.BatchDelJobResultRequest) (*meeting.BatchDelJobResultResponse, error) {
	return client.PostAndUnmarshal[meeting.BatchDelJobResultResponse](s.client, ctx, "/cgi-bin/meeting/vip/batch_del_job_result", req)
}

// GetVIPList 获取高级功能账号列表
// 文档: docs/录制管理/获取高级功能账号列表.md
func (s *Service) GetVIPList(ctx context.Context, req *meeting.GetVIPListRequest) (*meeting.GetVIPListResponse, error) {
	return client.PostAndUnmarshal[meeting.GetVIPListResponse](s.client, ctx, "/cgi-bin/meeting/vip/list", req)
}

// DeleteMeetingRecord 删除会议录制（删除会议的所有录制文件）
// 文档: docs/录制管理/删除会议录制.md
func (s *Service) DeleteMeetingRecord(ctx context.Context, req *meeting.DeleteMeetingRecordRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/meeting/record/delete", req)
	return err
}

// DeleteRecordFile 删除单个录制文件
// 文档: docs/录制管理/删除单个录制文件.md
func (s *Service) DeleteRecordFile(ctx context.Context, req *meeting.DeleteRecordFileRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/meeting/record/delete_file", req)
	return err
}

// GetRecordList 获取会议录制列表
// 文档: docs/录制管理/获取会议录制列表.md
func (s *Service) GetRecordList(ctx context.Context, req *meeting.GetRecordListRequest) (*meeting.GetRecordListResponse, error) {
	return client.PostAndUnmarshal[meeting.GetRecordListResponse](s.client, ctx, "/cgi-bin/meeting/record/list", req)
}

// GetFileList 获取会议录制地址（文件列表）
// 文档: docs/录制管理/获取会议录制地址.md
func (s *Service) GetFileList(ctx context.Context, req *meeting.GetFileListRequest) (*meeting.GetFileListResponse, error) {
	return client.PostAndUnmarshal[meeting.GetFileListResponse](s.client, ctx, "/cgi-bin/meeting/record/get_file_list", req)
}

// GetRecordFile 获取单个录制文件详情
// 文档: docs/录制管理/获取单个录制文件详情.md
func (s *Service) GetRecordFile(ctx context.Context, req *meeting.GetRecordFileRequest) (*meeting.GetRecordFileResponse, error) {
	return client.PostAndUnmarshal[meeting.GetRecordFileResponse](s.client, ctx, "/cgi-bin/meeting/record/get_file", req)
}

// GetStatistics 获取录制文件访问统计
// 文档: docs/录制管理/获取录制文件访问统计.md
func (s *Service) GetStatistics(ctx context.Context, req *meeting.GetStatisticsRequest) (*meeting.GetStatisticsResponse, error) {
	return client.PostAndUnmarshal[meeting.GetStatisticsResponse](s.client, ctx, "/cgi-bin/meeting/record/get_statistics", req)
}

// TranscriptSearch 根据指定内容搜索录制转写
// 文档: docs/录制管理/获取录制转写搜索结果.md
func (s *Service) TranscriptSearch(ctx context.Context, req *meeting.TranscriptSearchRequest) (*meeting.TranscriptSearchResponse, error) {
	return client.PostAndUnmarshal[meeting.TranscriptSearchResponse](s.client, ctx, "/cgi-bin/meeting/record/transcript/search", req)
}

// GetTranscriptParagraphList 获取录制转写段落信息
// 文档: docs/录制管理/获取录制转写段落信息.md
func (s *Service) GetTranscriptParagraphList(ctx context.Context, req *meeting.GetTranscriptParagraphListRequest) (*meeting.GetTranscriptParagraphListResponse, error) {
	return client.PostAndUnmarshal[meeting.GetTranscriptParagraphListResponse](s.client, ctx, "/cgi-bin/meeting/record/transcript/get_paragraph_list", req)
}

// GetTranscriptDetail 获取录制转写详情
// 文档: docs/录制管理/获取录制转写详情.md
func (s *Service) GetTranscriptDetail(ctx context.Context, req *meeting.GetTranscriptDetailRequest) (*meeting.GetTranscriptDetailResponse, error) {
	return client.PostAndUnmarshal[meeting.GetTranscriptDetailResponse](s.client, ctx, "/cgi-bin/meeting/record/transcript/get_detail", req)
}
