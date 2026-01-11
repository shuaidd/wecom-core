package email

// CreateGroupRequest 创建邮件群组请求
type CreateGroupRequest struct {
	GroupID             string      `json:"groupid"`                        // 邮件群组ID，邮箱格式(必填)
	GroupName           string      `json:"groupname"`                      // 邮件群组名称(必填)
	EmailList           *StringList `json:"email_list,omitempty"`           // 群组内成员邮箱地址
	TagList             *IDList     `json:"tag_list,omitempty"`             // 群组内包含的标签ID
	DepartmentList      *IDList     `json:"department_list,omitempty"`      // 群组内包含的部门ID
	GroupList           *StringList `json:"group_list,omitempty"`           // 群组内包含的群组邮箱
	AllowType           uint32      `json:"allow_type,omitempty"`           // 群组使用权限 0:企业成员 1:任何人 2:组内成员 3:自定义成员
	AllowEmailList      *StringList `json:"allow_emaillist,omitempty"`      // 允许使用群组群发的成员邮箱地址
	AllowDepartmentList *IDList     `json:"allow_departmentlist,omitempty"` // 允许使用群组群发的部门ID
	AllowTagList        *IDList     `json:"allow_taglist,omitempty"`        // 允许使用群组群发的标签ID
}

// CreateGroupResponse 创建邮件群组响应
type CreateGroupResponse struct {
	ErrCode int32  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// GetGroupRequest 获取邮件群组详情请求(通过URL参数传递)
type GetGroupRequest struct {
	GroupID string `json:"groupid"` // 邮件群组ID，邮箱格式(必填)
}

// GetGroupResponse 获取邮件群组详情响应
type GetGroupResponse struct {
	ErrCode             int32       `json:"errcode"`
	ErrMsg              string      `json:"errmsg"`
	GroupID             string      `json:"groupid"`                        // 邮件群组ID
	GroupName           string      `json:"groupname"`                      // 邮件群组名称
	EmailList           *StringList `json:"email_list,omitempty"`           // 群组内成员邮箱地址
	TagList             *IDList     `json:"tag_list,omitempty"`             // 群组内包含的标签ID
	DepartmentList      *IDList     `json:"department_list,omitempty"`      // 群组内包含的部门ID
	GroupList           *StringList `json:"group_list,omitempty"`           // 群组内包含的群组邮箱ID
	AllowType           uint32      `json:"allow_type,omitempty"`           // 群组使用权限
	AllowEmailList      *StringList `json:"allow_emaillist,omitempty"`      // 允许使用群组群发的成员邮箱地址
	AllowDepartmentList *IDList     `json:"allow_departmentlist,omitempty"` // 允许使用群组群发的部门ID
	AllowTagList        *IDList     `json:"allow_taglist,omitempty"`        // 允许使用群组群发的标签ID
}

// UpdateGroupRequest 更新邮件群组请求
type UpdateGroupRequest struct {
	GroupID             string      `json:"groupid"`                        // 邮件群组ID，邮箱格式(必填)
	GroupName           string      `json:"groupname,omitempty"`            // 邮件群组名称
	EmailList           *StringList `json:"email_list,omitempty"`           // 群组内成员邮箱地址
	TagList             *IDList     `json:"tag_list,omitempty"`             // 群组内包含的标签ID
	DepartmentList      *IDList     `json:"department_list,omitempty"`      // 群组内包含的部门ID
	GroupList           *StringList `json:"group_list,omitempty"`           // 群组内包含的群组邮箱
	AllowType           *uint32     `json:"allow_type,omitempty"`           // 群组使用权限 0:企业成员 1:任何人 2:组内成员 3:自定义成员
	AllowEmailList      *StringList `json:"allow_emaillist,omitempty"`      // 允许使用群组群发的成员邮箱地址
	AllowDepartmentList *IDList     `json:"allow_departmentlist,omitempty"` // 允许使用群组群发的部门ID
	AllowTagList        *IDList     `json:"allow_taglist,omitempty"`        // 允许使用群组群发的标签ID
}

// UpdateGroupResponse 更新邮件群组响应
type UpdateGroupResponse struct {
	ErrCode int32  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// SearchGroupRequest 模糊搜索邮件群组请求(通过URL参数传递)
type SearchGroupRequest struct {
	Fuzzy   uint32 `json:"fuzzy"`             // 1:开启模糊搜索 0:获取全部邮件群组(必填)
	GroupID string `json:"groupid,omitempty"` // 邮件群组ID，邮箱格式
}

// GroupInfo 邮件群组基本信息
type GroupInfo struct {
	GroupID   string `json:"groupid"`   // 邮件群组ID
	GroupName string `json:"groupname"` // 邮件群组名称
}

// SearchGroupResponse 模糊搜索邮件群组响应
type SearchGroupResponse struct {
	ErrCode int32        `json:"errcode"`
	ErrMsg  string       `json:"errmsg"`
	Count   int          `json:"count"`            // 返回条数
	Groups  []*GroupInfo `json:"groups,omitempty"` // 邮件群组列表
}

// DeleteGroupRequest 删除邮件群组请求
type DeleteGroupRequest struct {
	GroupID string `json:"groupid"` // 邮件群组ID，邮箱格式(必填)
}

// DeleteGroupResponse 删除邮件群组响应
type DeleteGroupResponse struct {
	ErrCode int32  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
