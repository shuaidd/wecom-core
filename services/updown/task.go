package updown

import (
	"context"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/updown"
)

// GetTaskResult 获取异步任务结果
// 文档: https://developer.work.weixin.qq.com/document/path/95814
func (s *Service) GetTaskResult(ctx context.Context, jobID string) (*updown.GetTaskResultResponse, error) {
	query := url.Values{}
	query.Set("jobid", jobID)

	return client.GetAndUnmarshal[updown.GetTaskResultResponse](s.client, ctx, "/cgi-bin/corpgroup/getresult", query)
}
