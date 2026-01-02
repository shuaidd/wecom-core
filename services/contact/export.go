package contact

import (
	"context"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/contact"
)

// ExportSimpleUser 导出成员
// 文档: https://developer.work.weixin.qq.com/document/path/94849
func (s *Service) ExportSimpleUser(ctx context.Context, req *contact.ExportRequest) (string, error) {
	result, err := client.PostAndUnmarshal[contact.ExportResponse](s.client, ctx, "/cgi-bin/export/simple_user", req)
	if err != nil {
		return "", err
	}

	return result.JobID, nil
}

// ExportUser 导出成员详情
// 文档: https://developer.work.weixin.qq.com/document/path/94851
func (s *Service) ExportUser(ctx context.Context, req *contact.ExportRequest) (string, error) {
	result, err := client.PostAndUnmarshal[contact.ExportResponse](s.client, ctx, "/cgi-bin/export/user", req)
	if err != nil {
		return "", err
	}

	return result.JobID, nil
}

// ExportDepartment 导出部门
// 文档: https://developer.work.weixin.qq.com/document/path/94852
func (s *Service) ExportDepartment(ctx context.Context, req *contact.ExportRequest) (string, error) {
	result, err := client.PostAndUnmarshal[contact.ExportResponse](s.client, ctx, "/cgi-bin/export/department", req)
	if err != nil {
		return "", err
	}

	return result.JobID, nil
}

// ExportTagUser 导出标签成员
// 文档: https://developer.work.weixin.qq.com/document/path/94853
func (s *Service) ExportTagUser(ctx context.Context, req *contact.ExportTagUserRequest) (string, error) {
	result, err := client.PostAndUnmarshal[contact.ExportResponse](s.client, ctx, "/cgi-bin/export/taguser", req)
	if err != nil {
		return "", err
	}

	return result.JobID, nil
}

// GetExportResult 获取导出结果
// 文档: https://developer.work.weixin.qq.com/document/path/94854
func (s *Service) GetExportResult(ctx context.Context, jobID string) (*contact.GetExportResultResponse, error) {
	query := url.Values{}
	query.Set("jobid", jobID)

	return client.GetAndUnmarshal[contact.GetExportResultResponse](s.client, ctx, "/cgi-bin/export/get_result", query)
}
