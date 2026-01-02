package contact

import "github.com/shuaidd/wecom-core/types/common"

// CreateDepartmentRequest 创建部门请求
type CreateDepartmentRequest struct {
	// Name 部门名称，必填
	Name string `json:"name"`
	// NameEN 部门英文名称
	NameEN string `json:"name_en,omitempty"`
	// ParentID 父部门id，必填
	ParentID int `json:"parentid"`
	// Order 在父部门中的次序值
	Order int `json:"order,omitempty"`
	// ID 部门id，可选。若不填该参数，将自动生成id
	ID int `json:"id,omitempty"`
}

// CreateDepartmentResponse 创建部门响应
type CreateDepartmentResponse struct {
	common.Response
	ID int `json:"id"`
}

// UpdateDepartmentRequest 更新部门请求
type UpdateDepartmentRequest struct {
	// ID 部门id，必填
	ID int `json:"id"`
	// Name 部门名称
	Name string `json:"name,omitempty"`
	// NameEN 部门英文名称
	NameEN string `json:"name_en,omitempty"`
	// ParentID 父部门id
	ParentID int `json:"parentid,omitempty"`
	// Order 在父部门中的次序值
	Order int `json:"order,omitempty"`
}

// UpdateDepartmentResponse 更新部门响应
type UpdateDepartmentResponse struct {
	common.Response
}

// DeleteDepartmentResponse 删除部门响应
type DeleteDepartmentResponse struct {
	common.Response
}

// GetDepartmentResponse 获取部门详情响应
type GetDepartmentResponse struct {
	common.Response
	Department Department `json:"department"`
}

// ListDepartmentsResponse 获取部门列表响应
type ListDepartmentsResponse struct {
	common.Response
	Department []Department `json:"department"`
}

// SimpleDepartment 简化的部门信息
type SimpleDepartment struct {
	ID       int `json:"id"`
	ParentID int `json:"parentid"`
	Order    int `json:"order"`
}

// ListSimpleDepartmentsResponse 获取子部门ID列表响应
type ListSimpleDepartmentsResponse struct {
	common.Response
	DepartmentID []SimpleDepartment `json:"department_id"`
}
