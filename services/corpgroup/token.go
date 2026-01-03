package corpgroup

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/corpgroup"
)

// GetToken 获取下级/下游企业的access_token
// 文档: https://developer.work.weixin.qq.com/document/path/93359
func (s *Service) GetToken(ctx context.Context, req *corpgroup.GetTokenRequest) (*corpgroup.GetTokenResponse, error) {
	return client.PostAndUnmarshal[corpgroup.GetTokenResponse](s.client, ctx, "/cgi-bin/corpgroup/corp/gettoken", req)
}
