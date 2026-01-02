package contact

import (
	"context"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/contact"
)

// SyncUsers 增量更新成员
// 文档: https://developer.work.weixin.qq.com/document/path/90980
func (s *Service) SyncUsers(ctx context.Context, req *contact.SyncUsersRequest) (string, error) {
	result, err := client.PostAndUnmarshal[contact.SyncUsersResponse](s.client, ctx, "/cgi-bin/batch/syncuser", req)
	if err != nil {
		return "", err
	}

	return result.JobID, nil
}

// ReplaceUsers 全量覆盖成员
// 文档: https://developer.work.weixin.qq.com/document/path/90981
func (s *Service) ReplaceUsers(ctx context.Context, req *contact.ReplaceUsersRequest) (string, error) {
	result, err := client.PostAndUnmarshal[contact.ReplaceUsersResponse](s.client, ctx, "/cgi-bin/batch/replaceuser", req)
	if err != nil {
		return "", err
	}

	return result.JobID, nil
}

// ReplaceDepartments 全量覆盖部门
// 文档: https://developer.work.weixin.qq.com/document/path/90982
func (s *Service) ReplaceDepartments(ctx context.Context, req *contact.ReplaceDepartmentsRequest) (string, error) {
	result, err := client.PostAndUnmarshal[contact.ReplaceDepartmentsResponse](s.client, ctx, "/cgi-bin/batch/replaceparty", req)
	if err != nil {
		return "", err
	}

	return result.JobID, nil
}

// GetBatchResult 获取异步任务结果
// 文档: https://developer.work.weixin.qq.com/document/path/90983
func (s *Service) GetBatchResult(ctx context.Context, jobID string) (*contact.GetBatchResultResponse, error) {
	query := url.Values{}
	query.Set("jobid", jobID)

	return client.GetAndUnmarshal[contact.GetBatchResultResponse](s.client, ctx, "/cgi-bin/batch/getresult", query)
}
