package contact

import (
	"context"
	"fmt"
	"net/url"

	"github.com/shuaidd/wecom-core/types/contact"
)

// CreateUser 创建成员
// 文档: https://developer.work.weixin.qq.com/document/path/90195
func (s *Service) CreateUser(ctx context.Context, req *contact.CreateUserRequest) (*contact.CreateUserResponse, error) {
	resp, err := s.client.Post(ctx, "/cgi-bin/user/create", req)
	if err != nil {
		return nil, err
	}

	var result contact.CreateUserResponse
	if err := resp.Unmarshal(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetUser 读取成员
// 文档: https://developer.work.weixin.qq.com/document/path/90196
func (s *Service) GetUser(ctx context.Context, userID string) (*contact.User, error) {
	query := url.Values{}
	query.Set("userid", userID)

	resp, err := s.client.Get(ctx, "/cgi-bin/user/get", query)
	if err != nil {
		return nil, err
	}

	var result contact.GetUserResponse
	if err := resp.Unmarshal(&result); err != nil {
		return nil, err
	}

	return &result.User, nil
}

// UpdateUser 更新成员
// 文档: https://developer.work.weixin.qq.com/document/path/90197
func (s *Service) UpdateUser(ctx context.Context, req *contact.UpdateUserRequest) error {
	_, err := s.client.Post(ctx, "/cgi-bin/user/update", req)
	return err
}

// DeleteUser 删除成员
// 文档: https://developer.work.weixin.qq.com/document/path/90198
func (s *Service) DeleteUser(ctx context.Context, userID string) error {
	query := url.Values{}
	query.Set("userid", userID)

	_, err := s.client.Get(ctx, "/cgi-bin/user/delete", query)
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

	resp, err := s.client.Get(ctx, "/cgi-bin/user/simplelist", query)
	if err != nil {
		return nil, err
	}

	var result contact.ListUsersResponse
	if err := resp.Unmarshal(&result); err != nil {
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

	resp, err := s.client.Get(ctx, "/cgi-bin/user/list", query)
	if err != nil {
		return nil, err
	}

	var result contact.ListUsersDetailResponse
	if err := resp.Unmarshal(&result); err != nil {
		return nil, err
	}

	return result.UserList, nil
}
