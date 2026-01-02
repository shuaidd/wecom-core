package contact

import (
	"context"
	"fmt"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/contact"
)

// CreateDepartment 创建部门
// 文档: https://developer.work.weixin.qq.com/document/path/90205
func (s *Service) CreateDepartment(ctx context.Context, req *contact.CreateDepartmentRequest) (int, error) {
	result, err := client.PostAndUnmarshal[contact.CreateDepartmentResponse](s.client, ctx, "/cgi-bin/department/create", req)
	if err != nil {
		return 0, err
	}

	return result.ID, nil
}

// UpdateDepartment 更新部门
// 文档: https://developer.work.weixin.qq.com/document/path/90206
func (s *Service) UpdateDepartment(ctx context.Context, req *contact.UpdateDepartmentRequest) error {
	_, err := client.PostAndUnmarshal[contact.UpdateDepartmentResponse](s.client, ctx, "/cgi-bin/department/update", req)
	return err
}

// DeleteDepartment 删除部门
// 文档: https://developer.work.weixin.qq.com/document/path/90207
func (s *Service) DeleteDepartment(ctx context.Context, id int) error {
	query := url.Values{}
	query.Set("id", fmt.Sprintf("%d", id))

	_, err := client.GetAndUnmarshal[contact.DeleteDepartmentResponse](s.client, ctx, "/cgi-bin/department/delete", query)
	return err
}

// ListDepartments 获取部门列表
// 文档: https://developer.work.weixin.qq.com/document/path/90208
// id: 部门id。获取指定部门及其下的子部门（以及子部门的子部门等等，递归）。如果不填，默认获取全量组织架构
func (s *Service) ListDepartments(ctx context.Context, id int) ([]contact.Department, error) {
	query := url.Values{}
	if id > 0 {
		query.Set("id", fmt.Sprintf("%d", id))
	}

	result, err := client.GetAndUnmarshal[contact.ListDepartmentsResponse](s.client, ctx, "/cgi-bin/department/list", query)
	if err != nil {
		return nil, err
	}

	return result.Department, nil
}

// GetDepartment 获取单个部门详情
// 文档: https://developer.work.weixin.qq.com/document/path/95351
func (s *Service) GetDepartment(ctx context.Context, id int) (*contact.Department, error) {
	query := url.Values{}
	query.Set("id", fmt.Sprintf("%d", id))

	result, err := client.GetAndUnmarshal[contact.GetDepartmentResponse](s.client, ctx, "/cgi-bin/department/get", query)
	if err != nil {
		return nil, err
	}

	return &result.Department, nil
}

// ListSimpleDepartments 获取子部门ID列表
// 文档: https://developer.work.weixin.qq.com/document/path/95350
func (s *Service) ListSimpleDepartments(ctx context.Context, id int) ([]contact.SimpleDepartment, error) {
	query := url.Values{}
	if id > 0 {
		query.Set("id", fmt.Sprintf("%d", id))
	}

	result, err := client.GetAndUnmarshal[contact.ListSimpleDepartmentsResponse](s.client, ctx, "/cgi-bin/department/simplelist", query)
	if err != nil {
		return nil, err
	}

	return result.DepartmentID, nil
}
