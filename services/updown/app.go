package updown

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/updown"
)

// ListAppShareInfo 获取应用共享信息
// 文档: https://developer.work.weixin.qq.com/document/path/93403
func (s *Service) ListAppShareInfo(ctx context.Context, req *updown.ListAppShareInfoRequest) (*updown.ListAppShareInfoResponse, error) {
	return client.PostAndUnmarshal[updown.ListAppShareInfoResponse](s.client, ctx, "/cgi-bin/corpgroup/corp/list_app_share_info", req)
}
