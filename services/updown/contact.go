package updown

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/updown"
)

// ImportChainContact 批量导入上下游联系人
// 文档: https://developer.work.weixin.qq.com/document/path/95813
func (s *Service) ImportChainContact(ctx context.Context, req *updown.ImportChainContactRequest) (string, error) {
	result, err := client.PostAndUnmarshal[updown.ImportChainContactResponse](s.client, ctx, "/cgi-bin/corpgroup/import_chain_contact", req)
	if err != nil {
		return "", err
	}
	return result.JobID, nil
}
