package corpgroup

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/corpgroup"
)

// TransferSession 获取下级/下游企业小程序session
// 文档: https://developer.work.weixin.qq.com/document/path/95792
func (s *Service) TransferSession(ctx context.Context, req *corpgroup.TransferSessionRequest) (*corpgroup.TransferSessionResponse, error) {
	return client.PostAndUnmarshal[corpgroup.TransferSessionResponse](s.client, ctx, "/cgi-bin/miniprogram/transfer_session", req)
}
