package externalcontact

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/externalcontact"
)

// GetContactList 获取已服务的外部联系人
// 企业可通过此接口获取所有已服务的外部联系人，及其添加人和加入的群聊
// 文档: https://developer.work.weixin.qq.com/document/path/97297
func (s *Service) GetContactList(ctx context.Context, req *externalcontact.GetContactListRequest) (*externalcontact.GetContactListResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetContactListResponse](s.client, ctx, "/cgi-bin/externalcontact/contact_list", req)
}
