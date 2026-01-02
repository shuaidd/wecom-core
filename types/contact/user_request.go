package contact

import "github.com/shuaidd/wecom-core/types/common"

// CreateUserRequest 创建成员请求
type CreateUserRequest struct {
	// UserID 成员UserID，必填
	UserID string `json:"userid"`
	// Name 成员名称，必填
	Name string `json:"name"`
	// Alias 别名
	Alias string `json:"alias,omitempty"`
	// Mobile 手机号码
	Mobile string `json:"mobile,omitempty"`
	// Department 成员所属部门id列表
	Department []int `json:"department,omitempty"`
	// Order 部门内的排序值
	Order []int `json:"order,omitempty"`
	// Position 职务信息
	Position string `json:"position,omitempty"`
	// Gender 性别。1表示男性，2表示女性
	Gender string `json:"gender,omitempty"`
	// Email 邮箱
	Email string `json:"email,omitempty"`
	// BizMail 企业邮箱
	BizMail string `json:"biz_mail,omitempty"`
	// Telephone 座机
	Telephone string `json:"telephone,omitempty"`
	// IsLeaderInDept 是否为部门负责人
	IsLeaderInDept []int `json:"is_leader_in_dept,omitempty"`
	// DirectLeader 直属上级UserID
	DirectLeader []string `json:"direct_leader,omitempty"`
	// AvatarMediaID 成员头像的mediaid
	AvatarMediaID string `json:"avatar_mediaid,omitempty"`
	// Enable 启用/禁用成员。1表示启用成员，0表示禁用成员
	Enable int `json:"enable,omitempty"`
	// ExtAttr 扩展属性
	ExtAttr *ExtAttr `json:"extattr,omitempty"`
	// ToInvite 是否邀请该成员使用企业微信
	ToInvite bool `json:"to_invite,omitempty"`
	// ExternalProfile 成员对外属性
	ExternalProfile *ExternalProfile `json:"external_profile,omitempty"`
	// ExternalPosition 对外职务
	ExternalPosition string `json:"external_position,omitempty"`
	// Address 地址
	Address string `json:"address,omitempty"`
	// MainDepartment 主部门
	MainDepartment int `json:"main_department,omitempty"`
}

// DepartmentInfo 部门信息
type DepartmentInfo struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// CreatedDepartmentList 新增的部门列表
type CreatedDepartmentList struct {
	DepartmentInfo []DepartmentInfo `json:"department_info,omitempty"`
}

// CreateUserResponse 创建成员响应
type CreateUserResponse struct {
	common.Response
	CreatedDepartmentList *CreatedDepartmentList `json:"created_department_list,omitempty"`
}

// GetUserResponse 读取成员响应
type GetUserResponse struct {
	common.Response
	User
}

// UpdateUserRequest 更新成员请求
type UpdateUserRequest struct {
	// UserID 成员UserID，必填
	UserID string `json:"userid"`
	// Name 成员名称
	Name string `json:"name,omitempty"`
	// Alias 别名
	Alias string `json:"alias,omitempty"`
	// Mobile 手机号码
	Mobile string `json:"mobile,omitempty"`
	// Department 成员所属部门id列表
	Department []int `json:"department,omitempty"`
	// Order 部门内的排序值
	Order []int `json:"order,omitempty"`
	// Position 职务信息
	Position string `json:"position,omitempty"`
	// Gender 性别。1表示男性，2表示女性
	Gender string `json:"gender,omitempty"`
	// Email 邮箱
	Email string `json:"email,omitempty"`
	// BizMail 企业邮箱
	BizMail string `json:"biz_mail,omitempty"`
	// Telephone 座机
	Telephone string `json:"telephone,omitempty"`
	// IsLeaderInDept 是否为部门负责人
	IsLeaderInDept []int `json:"is_leader_in_dept,omitempty"`
	// DirectLeader 直属上级UserID
	DirectLeader []string `json:"direct_leader,omitempty"`
	// AvatarMediaID 成员头像的mediaid
	AvatarMediaID string `json:"avatar_mediaid,omitempty"`
	// Enable 启用/禁用成员。1表示启用成员，0表示禁用成员
	Enable *int `json:"enable,omitempty"`
	// ExtAttr 扩展属性
	ExtAttr *ExtAttr `json:"extattr,omitempty"`
	// ExternalProfile 成员对外属性
	ExternalProfile *ExternalProfile `json:"external_profile,omitempty"`
	// ExternalPosition 对外职务
	ExternalPosition string `json:"external_position,omitempty"`
	// Address 地址
	Address string `json:"address,omitempty"`
	// MainDepartment 主部门
	MainDepartment int `json:"main_department,omitempty"`
}

// UpdateUserResponse 更新成员响应
type UpdateUserResponse struct {
	common.Response
}

// DeleteUserResponse 删除成员响应
type DeleteUserResponse struct {
	common.Response
}

// SimpleUser 简化的成员信息
type SimpleUser struct {
	UserID     string `json:"userid"`
	Name       string `json:"name"`
	Department []int  `json:"department"`
}

// ListUsersResponse 获取部门成员列表响应
type ListUsersResponse struct {
	common.Response
	UserList []SimpleUser `json:"userlist"`
}

// ListUsersDetailResponse 获取部门成员详情列表响应
type ListUsersDetailResponse struct {
	common.Response
	UserList []User `json:"userlist"`
}

// AuthSuccessResponse 二次验证响应
type AuthSuccessResponse struct {
	common.Response
}

// ConvertToOpenIDRequest userid转openid请求
type ConvertToOpenIDRequest struct {
	UserID string `json:"userid"`
}

// ConvertToOpenIDResponse userid转openid响应
type ConvertToOpenIDResponse struct {
	common.Response
	OpenID string `json:"openid"`
}

// ConvertToUserIDRequest openid转userid请求
type ConvertToUserIDRequest struct {
	OpenID string `json:"openid"`
}

// ConvertToUserIDResponse openid转userid响应
type ConvertToUserIDResponse struct {
	common.Response
	UserID string `json:"userid"`
}

// GetUserIDByEmailRequest 邮箱获取userid请求
type GetUserIDByEmailRequest struct {
	Email     string `json:"email"`
	EmailType int    `json:"email_type,omitempty"`
}

// GetUserIDByEmailResponse 邮箱获取userid响应
type GetUserIDByEmailResponse struct {
	common.Response
	UserID string `json:"userid"`
}

// GetUserIDByMobileRequest 手机号获取userid请求
type GetUserIDByMobileRequest struct {
	Mobile string `json:"mobile"`
}

// GetUserIDByMobileResponse 手机号获取userid响应
type GetUserIDByMobileResponse struct {
	common.Response
	UserID string `json:"userid"`
}

// BatchDeleteUsersRequest 批量删除成员请求
type BatchDeleteUsersRequest struct {
	UserIDList []string `json:"useridlist"`
}

// BatchDeleteUsersResponse 批量删除成员响应
type BatchDeleteUsersResponse struct {
	common.Response
}

