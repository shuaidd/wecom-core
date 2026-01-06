package email

import (
	"context"
	"fmt"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/email"
)

// CreateGroup 创建邮件群组
// 该接口用于创建新邮件群组，可以指定群组成员，定义群组使用权限范围
//
// 文档: https://developer.work.weixin.qq.com/document/path/95838
func (s *Service) CreateGroup(ctx context.Context, req *email.CreateGroupRequest) (*email.CreateGroupResponse, error) {
	return client.PostAndUnmarshal[email.CreateGroupResponse](s.client, ctx, "/cgi-bin/exmail/group/create", req)
}

// GetGroup 获取邮件群组详情
// 该接口用于获取邮件群组详细信息，包含群组名称、群组成员、群组使用权限等
//
// 文档: https://developer.work.weixin.qq.com/document/path/95839
func (s *Service) GetGroup(ctx context.Context, groupID string) (*email.GetGroupResponse, error) {
	query := url.Values{}
	query.Set("groupid", groupID)
	return client.GetAndUnmarshal[email.GetGroupResponse](s.client, ctx, "/cgi-bin/exmail/group/get", query)
}

// UpdateGroup 更新邮件群组
// 该接口用于更新邮件群组，可以修改群组名称、群组成员、群组使用权限等
// 注意：Json数组类型传空值将会清空其内容，不传则保持不变
//
// 文档: https://developer.work.weixin.qq.com/document/path/95840
func (s *Service) UpdateGroup(ctx context.Context, req *email.UpdateGroupRequest) (*email.UpdateGroupResponse, error) {
	return client.PostAndUnmarshal[email.UpdateGroupResponse](s.client, ctx, "/cgi-bin/exmail/group/update", req)
}

// SearchGroup 模糊搜索邮件群组
// 该接口用于通过群组ID模糊搜索邮件群组
//
// 文档: https://developer.work.weixin.qq.com/document/path/95841
func (s *Service) SearchGroup(ctx context.Context, fuzzy uint32, groupID string) (*email.SearchGroupResponse, error) {
	query := url.Values{}
	query.Set("fuzzy", fmt.Sprintf("%d", fuzzy))
	if groupID != "" {
		query.Set("groupid", groupID)
	}
	return client.GetAndUnmarshal[email.SearchGroupResponse](s.client, ctx, "/cgi-bin/exmail/group/search", query)
}

// DeleteGroup 删除邮件群组
// 该接口用于删除已有的邮件群组
//
// 文档: https://developer.work.weixin.qq.com/document/path/95842
func (s *Service) DeleteGroup(ctx context.Context, groupID string) (*email.DeleteGroupResponse, error) {
	req := &email.DeleteGroupRequest{
		GroupID: groupID,
	}
	return client.PostAndUnmarshal[email.DeleteGroupResponse](s.client, ctx, "/cgi-bin/exmail/group/delete", req)
}
