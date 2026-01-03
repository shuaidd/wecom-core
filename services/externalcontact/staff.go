package externalcontact

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/externalcontact"
)

// GetFollowUserList 获取配置了客户联系功能的成员列表
// 企业和第三方服务商可通过此接口获取配置了客户联系功能的成员列表
// 文档: https://developer.work.weixin.qq.com/document/path/92571
func (s *Service) GetFollowUserList(ctx context.Context) (*externalcontact.GetFollowUserListResponse, error) {
	return client.GetAndUnmarshal[externalcontact.GetFollowUserListResponse](s.client, ctx, "/cgi-bin/externalcontact/get_follow_user_list", nil)
}
