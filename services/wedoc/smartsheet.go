package wedoc

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/wedoc"
)

// Smartsheet API URLs
const (
	// 记录管理
	addRecordsURL    = "/cgi-bin/wedoc/smartsheet/add_records"
	getRecordsURL    = "/cgi-bin/wedoc/smartsheet/get_records"
	updateRecordsURL = "/cgi-bin/wedoc/smartsheet/update_records"
	deleteRecordsURL = "/cgi-bin/wedoc/smartsheet/delete_records"

	// 字段管理
	addFieldsURL    = "/cgi-bin/wedoc/smartsheet/add_fields"
	getFieldsURL    = "/cgi-bin/wedoc/smartsheet/get_fields"
	updateFieldsURL = "/cgi-bin/wedoc/smartsheet/update_fields"
	deleteFieldsURL = "/cgi-bin/wedoc/smartsheet/delete_fields"

	// 视图管理
	addViewURL    = "/cgi-bin/wedoc/smartsheet/add_view"
	getViewsURL   = "/cgi-bin/wedoc/smartsheet/get_views"
	updateViewURL = "/cgi-bin/wedoc/smartsheet/update_view"
	deleteViewURL = "/cgi-bin/wedoc/smartsheet/delete_view"

	// 子表管理
	addSheetURL    = "/cgi-bin/wedoc/smartsheet/add_sheet"
	getSheetURL    = "/cgi-bin/wedoc/smartsheet/get_sheet"
	updateSheetURL = "/cgi-bin/wedoc/smartsheet/update_sheet"
	deleteSheetURL = "/cgi-bin/wedoc/smartsheet/delete_sheet"

	// 编组管理
	addFieldGroupURL    = "/cgi-bin/wedoc/smartsheet/add_field_group"
	getFieldGroupURL    = "/cgi-bin/wedoc/smartsheet/get_field_group"
	updateFieldGroupURL = "/cgi-bin/wedoc/smartsheet/update_field_group"
	deleteFieldGroupURL = "/cgi-bin/wedoc/smartsheet/delete_field_group"
)

// ==================== 记录管理 ====================

// AddRecords 添加记录
// 本接口用于在 Smartsheet 中的某个子表里添加一行或多行新记录
func (s *Service) AddRecords(ctx context.Context, req *wedoc.AddRecordsRequest) (*wedoc.AddRecordsResponse, error) {
	return client.PostAndUnmarshal[wedoc.AddRecordsResponse](s.client, ctx, addRecordsURL, req)
}

// GetRecords 查询记录
// 本接口用于获取 Smartsheet 中某个子表下记录信息
func (s *Service) GetRecords(ctx context.Context, req *wedoc.GetRecordsRequest) (*wedoc.GetRecordsResponse, error) {
	return client.PostAndUnmarshal[wedoc.GetRecordsResponse](s.client, ctx, getRecordsURL, req)
}

// UpdateRecords 更新记录
// 本接口用于更新 Smartsheet 中的某个子表里的一行或多行记录
func (s *Service) UpdateRecords(ctx context.Context, req *wedoc.UpdateRecordsRequest) (*wedoc.UpdateRecordsResponse, error) {
	return client.PostAndUnmarshal[wedoc.UpdateRecordsResponse](s.client, ctx, updateRecordsURL, req)
}

// DeleteRecords 删除记录
// 本接口用于删除 Smartsheet 的某个子表中的一行或多行记录
func (s *Service) DeleteRecords(ctx context.Context, req *wedoc.DeleteRecordsRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, deleteRecordsURL, req)
	return err
}

// ==================== 字段管理 ====================

// AddFields 添加字段
// 本接口用于在智能表中的某个子表里添加一列或多列新字段
func (s *Service) AddFields(ctx context.Context, req *wedoc.AddFieldsRequest) (*wedoc.AddFieldsResponse, error) {
	return client.PostAndUnmarshal[wedoc.AddFieldsResponse](s.client, ctx, addFieldsURL, req)
}

// GetFields 查询字段
// 本接口用于获取智能表中某个子表下字段信息
func (s *Service) GetFields(ctx context.Context, req *wedoc.GetFieldsRequest) (*wedoc.GetFieldsResponse, error) {
	return client.PostAndUnmarshal[wedoc.GetFieldsResponse](s.client, ctx, getFieldsURL, req)
}

// UpdateFields 更新字段
// 本接口用于更新智能表中某个子表里的字段
func (s *Service) UpdateFields(ctx context.Context, req *wedoc.UpdateFieldsRequest) (*wedoc.UpdateFieldsResponse, error) {
	return client.PostAndUnmarshal[wedoc.UpdateFieldsResponse](s.client, ctx, updateFieldsURL, req)
}

// DeleteFields 删除字段
// 本接口用于删除智能表中某个子表里的字段
func (s *Service) DeleteFields(ctx context.Context, req *wedoc.DeleteFieldsRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, deleteFieldsURL, req)
	return err
}

// ==================== 视图管理 ====================

// AddView 添加视图
// 本接口用于在智能表中的某个子表里添加视图
func (s *Service) AddView(ctx context.Context, req *wedoc.AddViewRequest) (*wedoc.AddViewResponse, error) {
	return client.PostAndUnmarshal[wedoc.AddViewResponse](s.client, ctx, addViewURL, req)
}

// GetViews 查询视图
// 本接口用于获取 Smartsheet 中某个子表里全部视图信息
func (s *Service) GetViews(ctx context.Context, req *wedoc.GetViewsRequest) (*wedoc.GetViewsResponse, error) {
	return client.PostAndUnmarshal[wedoc.GetViewsResponse](s.client, ctx, getViewsURL, req)
}

// UpdateView 更新视图
// 本接口用于更新 Smartsheet 中的某个视图
func (s *Service) UpdateView(ctx context.Context, req *wedoc.UpdateViewRequest) (*wedoc.UpdateViewResponse, error) {
	return client.PostAndUnmarshal[wedoc.UpdateViewResponse](s.client, ctx, updateViewURL, req)
}

// DeleteView 删除视图
// 本接口用于删除 Smartsheet 中的某个视图
func (s *Service) DeleteView(ctx context.Context, req *wedoc.DeleteViewRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, deleteViewURL, req)
	return err
}

// ==================== 子表管理 ====================

// AddSheet 添加子表
// 本接口用于在表格的某个位置添加一个智能表
func (s *Service) AddSheet(ctx context.Context, req *wedoc.AddSheetRequest) (*wedoc.AddSheetResponse, error) {
	return client.PostAndUnmarshal[wedoc.AddSheetResponse](s.client, ctx, addSheetURL, req)
}

// GetSheet 查询子表
// 本接口用于查询一篇在线表格中全部智能表信息
func (s *Service) GetSheet(ctx context.Context, req *wedoc.GetSheetRequest) (*wedoc.GetSheetResponse, error) {
	return client.PostAndUnmarshal[wedoc.GetSheetResponse](s.client, ctx, getSheetURL, req)
}

// UpdateSheet 更新子表
// 本接口用于更新智能表中的子表属性
func (s *Service) UpdateSheet(ctx context.Context, req *wedoc.UpdateSheetRequest) (*wedoc.UpdateSheetResponse, error) {
	return client.PostAndUnmarshal[wedoc.UpdateSheetResponse](s.client, ctx, updateSheetURL, req)
}

// DeleteSheet 删除子表
// 本接口用于删除智能表中的某个子表
func (s *Service) DeleteSheet(ctx context.Context, req *wedoc.DeleteSheetRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, deleteSheetURL, req)
	return err
}

// ==================== 编组管理 ====================

// AddFieldGroup 添加编组
// 本接口用于在智能表中的某个子表里添加编组
func (s *Service) AddFieldGroup(ctx context.Context, req *wedoc.AddFieldGroupRequest) (*wedoc.AddFieldGroupResponse, error) {
	return client.PostAndUnmarshal[wedoc.AddFieldGroupResponse](s.client, ctx, addFieldGroupURL, req)
}

// GetFieldGroup 获取编组
// 本接口用于获取智能表中某个子表下的编组信息
func (s *Service) GetFieldGroup(ctx context.Context, req *wedoc.GetFieldGroupRequest) (*wedoc.GetFieldGroupResponse, error) {
	return client.PostAndUnmarshal[wedoc.GetFieldGroupResponse](s.client, ctx, getFieldGroupURL, req)
}

// UpdateFieldGroup 更新编组
// 本接口用于更新智能表中某个子表里的编组
func (s *Service) UpdateFieldGroup(ctx context.Context, req *wedoc.UpdateFieldGroupRequest) (*wedoc.UpdateFieldGroupResponse, error) {
	return client.PostAndUnmarshal[wedoc.UpdateFieldGroupResponse](s.client, ctx, updateFieldGroupURL, req)
}

// DeleteFieldGroup 删除编组
// 本接口用于删除智能表中某个子表里的编组
func (s *Service) DeleteFieldGroup(ctx context.Context, req *wedoc.DeleteFieldGroupRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, deleteFieldGroupURL, req)
	return err
}
