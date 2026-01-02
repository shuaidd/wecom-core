package contact

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/contact"
)

// ConvertTmpExternalUserID 转换tmp_external_userid
// 文档: https://developer.work.weixin.qq.com/document/path/95195
func (s *Service) ConvertTmpExternalUserID(ctx context.Context, req *contact.ConvertTmpExternalUserIDRequest) (*contact.ConvertTmpExternalUserIDResponse, error) {
	return client.PostAndUnmarshal[contact.ConvertTmpExternalUserIDResponse](s.client, ctx, "/cgi-bin/idconvert/convert_tmp_external_userid", req)
}
