package contact

import (
	"context"
	"fmt"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/contact"
)

// CreateUser 创建成员
// 文档: https://developer.work.weixin.qq.com/document/path/90195
func (s *Service) CreateUser(ctx context.Context, req *contact.CreateUserRequest) (*contact.CreateUserResponse, error) {
	return client.PostAndUnmarshal[contact.CreateUserResponse](s.client, ctx, "/cgi-bin/user/create", req)
}

// GetUser 读取成员
// 文档: https://developer.work.weixin.qq.com/document/path/90196
func (s *Service) GetUser(ctx context.Context, userID string) (*contact.User, error) {
	query := url.Values{}
	query.Set("userid", userID)

	result, err := client.GetAndUnmarshal[contact.GetUserResponse](s.client, ctx, "/cgi-bin/user/get", query)
	if err != nil {
		return nil, err
	}

	return &result.User, nil
}

// UpdateUser 更新成员
// 文档: https://developer.work.weixin.qq.com/document/path/90197
func (s *Service) UpdateUser(ctx context.Context, req *contact.UpdateUserRequest) error {
	_, err := client.PostAndUnmarshal[contact.UpdateUserResponse](s.client, ctx, "/cgi-bin/user/update", req)
	return err
}

// DeleteUser 删除成员
// 文档: https://developer.work.weixin.qq.com/document/path/90198
func (s *Service) DeleteUser(ctx context.Context, userID string) error {
	query := url.Values{}
	query.Set("userid", userID)

	_, err := client.GetAndUnmarshal[contact.DeleteUserResponse](s.client, ctx, "/cgi-bin/user/delete", query)
	return err
}

// ListUsers 获取部门成员
// 文档: https://developer.work.weixin.qq.com/document/path/90200
func (s *Service) ListUsers(ctx context.Context, departmentID int, fetchChild bool) ([]contact.SimpleUser, error) {
	query := url.Values{}
	query.Set("department_id", fmt.Sprintf("%d", departmentID))
	if fetchChild {
		query.Set("fetch_child", "1")
	} else {
		query.Set("fetch_child", "0")
	}

	result, err := client.GetAndUnmarshal[contact.ListUsersResponse](s.client, ctx, "/cgi-bin/user/simplelist", query)
	if err != nil {
		return nil, err
	}

	return result.UserList, nil
}

// ListUsersDetail 获取部门成员详情
// 文档: https://developer.work.weixin.qq.com/document/path/90201
func (s *Service) ListUsersDetail(ctx context.Context, departmentID int, fetchChild bool) ([]contact.User, error) {
	query := url.Values{}
	query.Set("department_id", fmt.Sprintf("%d", departmentID))
	if fetchChild {
		query.Set("fetch_child", "1")
	} else {
		query.Set("fetch_child", "0")
	}

	result, err := client.GetAndUnmarshal[contact.ListUsersDetailResponse](s.client, ctx, "/cgi-bin/user/list", query)
	if err != nil {
		return nil, err
	}

	return result.UserList, nil
}

// ListUserIDs 获取成员ID列表
// 文档: https://developer.work.weixin.qq.com/document/path/96067
func (s *Service) ListUserIDs(ctx context.Context, req *contact.ListUserIDsRequest) (*contact.ListUserIDsResponse, error) {
	return client.PostAndUnmarshal[contact.ListUserIDsResponse](s.client, ctx, "/cgi-bin/user/list_id", req)
}

// AuthSuccess 二次验证
// 文档: https://developer.work.weixin.qq.com/document/path/90203
func (s *Service) AuthSuccess(ctx context.Context, userID string) error {
	query := url.Values{}
	query.Set("userid", userID)

	_, err := client.GetAndUnmarshal[contact.AuthSuccessResponse](s.client, ctx, "/cgi-bin/user/authsucc", query)
	return err
}

// ConvertToOpenID userid转openid
// 文档: https://developer.work.weixin.qq.com/document/path/90202
func (s *Service) ConvertToOpenID(ctx context.Context, userID string) (string, error) {
	req := &contact.ConvertToOpenIDRequest{
		UserID: userID,
	}

	result, err := client.PostAndUnmarshal[contact.ConvertToOpenIDResponse](s.client, ctx, "/cgi-bin/user/convert_to_openid", req)
	if err != nil {
		return "", err
	}

	return result.OpenID, nil
}

// ConvertToUserID openid转userid
// 文档: https://developer.work.weixin.qq.com/document/path/90202
func (s *Service) ConvertToUserID(ctx context.Context, openID string) (string, error) {
	req := &contact.ConvertToUserIDRequest{
		OpenID: openID,
	}

	result, err := client.PostAndUnmarshal[contact.ConvertToUserIDResponse](s.client, ctx, "/cgi-bin/user/convert_to_userid", req)
	if err != nil {
		return "", err
	}

	return result.UserID, nil
}

// GetUserIDByEmail 通过邮箱获取userid
// 文档: https://developer.work.weixin.qq.com/document/path/95895
func (s *Service) GetUserIDByEmail(ctx context.Context, email string, emailType int) (string, error) {
	req := &contact.GetUserIDByEmailRequest{
		Email:     email,
		EmailType: emailType,
	}

	result, err := client.PostAndUnmarshal[contact.GetUserIDByEmailResponse](s.client, ctx, "/cgi-bin/user/get_userid_by_email", req)
	if err != nil {
		return "", err
	}

	return result.UserID, nil
}

// GetUserIDByMobile 通过手机号获取userid
// 文档: https://developer.work.weixin.qq.com/document/path/95402
func (s *Service) GetUserIDByMobile(ctx context.Context, mobile string) (string, error) {
	req := &contact.GetUserIDByMobileRequest{
		Mobile: mobile,
	}

	result, err := client.PostAndUnmarshal[contact.GetUserIDByMobileResponse](s.client, ctx, "/cgi-bin/user/getuserid", req)
	if err != nil {
		return "", err
	}

	return result.UserID, nil
}

// BatchDeleteUsers 批量删除成员
// 文档: https://developer.work.weixin.qq.com/document/path/90199
func (s *Service) BatchDeleteUsers(ctx context.Context, userIDList []string) error {
	req := &contact.BatchDeleteUsersRequest{
		UserIDList: userIDList,
	}

	_, err := client.PostAndUnmarshal[contact.BatchDeleteUsersResponse](s.client, ctx, "/cgi-bin/user/batchdelete", req)
	return err
}
