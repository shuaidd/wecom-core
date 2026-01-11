package email

// IDList ID列表
type IDList struct {
	List []uint32 `json:"list,omitempty"`
}

// StringList 字符串列表
type StringList struct {
	List []string `json:"list,omitempty"`
}

// AuthCodeInfo 客户端专用密码信息
type AuthCodeInfo struct {
	Remark string `json:"remark,omitempty"` // 客户端专用密码备注
}

// CreatePublicMailRequest 创建公共邮箱请求
type CreatePublicMailRequest struct {
	Email          string        `json:"email"`                      // 公共邮箱地址(必填)
	Name           string        `json:"name"`                       // 公共邮箱名称(必填)
	UserIDList     *StringList   `json:"userid_list,omitempty"`      // 有权限使用的成员UserID列表
	DepartmentList *IDList       `json:"department_list,omitempty"`  // 有权限使用的部门ID列表
	TagList        *IDList       `json:"tag_list,omitempty"`         // 有权限使用的标签ID列表
	CreateAuthCode uint32        `json:"create_auth_code,omitempty"` // 是否创建客户端专用密码 0-否 1-是
	AuthCodeInfo   *AuthCodeInfo `json:"auth_code_info,omitempty"`   // 客户端专用密码信息
}

// CreatePublicMailResponse 创建公共邮箱响应
type CreatePublicMailResponse struct {
	ErrCode    int32  `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	ID         uint32 `json:"id"`                     // 公共邮箱ID
	AuthCodeID uint32 `json:"auth_code_id,omitempty"` // 客户端专用密码ID
	AuthCode   string `json:"auth_code,omitempty"`    // 客户端专用密码
}

// UpdatePublicMailRequest 更新公共邮箱请求
type UpdatePublicMailRequest struct {
	ID             uint32        `json:"id"`                         // 公共邮箱ID(必填)
	Name           string        `json:"name,omitempty"`             // 公共邮箱名称
	UserIDList     *StringList   `json:"userid_list,omitempty"`      // 有权限使用的成员UserID列表
	DepartmentList *IDList       `json:"department_list,omitempty"`  // 有权限使用的部门ID列表
	TagList        *IDList       `json:"tag_list,omitempty"`         // 有权限使用的标签ID列表
	AliasList      *StringList   `json:"alias_list,omitempty"`       // 邮箱别名列表
	CreateAuthCode uint32        `json:"create_auth_code,omitempty"` // 是否创建客户端专用密码 0-否 1-是
	AuthCodeInfo   *AuthCodeInfo `json:"auth_code_info,omitempty"`   // 客户端专用密码信息
}

// UpdatePublicMailResponse 更新公共邮箱响应
type UpdatePublicMailResponse struct {
	ErrCode    int32  `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	AuthCodeID uint32 `json:"auth_code_id,omitempty"` // 客户端专用密码ID
	AuthCode   string `json:"auth_code,omitempty"`    // 客户端专用密码
}

// DeletePublicMailRequest 删除公共邮箱请求
type DeletePublicMailRequest struct {
	ID uint32 `json:"id"` // 公共邮箱ID(必填)
}

// DeletePublicMailResponse 删除公共邮箱响应
type DeletePublicMailResponse struct {
	ErrCode int32  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// GetPublicMailRequest 获取公共邮箱详情请求
type GetPublicMailRequest struct {
	IDList []uint32 `json:"id_list"` // 公共邮箱ID列表(必填)
}

// PublicMailDetail 公共邮箱详情
type PublicMailDetail struct {
	ID             uint32      `json:"id"`                        // 公共邮箱ID
	Email          string      `json:"email"`                     // 公共邮箱地址
	Name           string      `json:"name"`                      // 公共邮箱名称
	UserIDList     *StringList `json:"userid_list,omitempty"`     // 有权限使用的成员UserID列表
	DepartmentList *IDList     `json:"department_list,omitempty"` // 有权限使用的部门ID列表
	TagList        *IDList     `json:"tag_list,omitempty"`        // 有权限使用的标签ID列表
	AliasList      *StringList `json:"alias_list,omitempty"`      // 邮箱别名列表
}

// GetPublicMailResponse 获取公共邮箱详情响应
type GetPublicMailResponse struct {
	ErrCode int32               `json:"errcode"`
	ErrMsg  string              `json:"errmsg"`
	List    []*PublicMailDetail `json:"list,omitempty"` // 公共邮箱详情列表
}

// SearchPublicMailRequest 搜索公共邮箱请求(通过URL参数传递)
type SearchPublicMailRequest struct {
	Fuzzy uint32 `json:"fuzzy"` // 1-开启模糊搜索 0-获取全部公共邮箱(必填)
	Email string `json:"email"` // 公共邮箱名称或邮箱地址
}

// PublicMailInfo 公共邮箱基本信息
type PublicMailInfo struct {
	ID    uint32 `json:"id"`    // 公共邮箱ID
	Email string `json:"email"` // 公共邮箱地址
	Name  string `json:"name"`  // 公共邮箱名称
}

// SearchPublicMailResponse 搜索公共邮箱响应
type SearchPublicMailResponse struct {
	ErrCode int32             `json:"errcode"`
	ErrMsg  string            `json:"errmsg"`
	List    []*PublicMailInfo `json:"list,omitempty"` // 公共邮箱列表
}

// GetAuthCodeListRequest 获取客户端专用密码列表请求
type GetAuthCodeListRequest struct {
	ID uint32 `json:"id"` // 公共邮箱ID(必填)
}

// AuthCodeDetail 客户端专用密码详情
type AuthCodeDetail struct {
	AuthCodeID  uint32 `json:"auth_code_id"`  // 客户端专用密码ID
	CreateTime  uint32 `json:"create_time"`   // 创建时间戳
	LastUseTime uint32 `json:"last_use_time"` // 最后使用时间(未使用过返回0)
	Remark      string `json:"remark"`        // 备注
}

// GetAuthCodeListResponse 获取客户端专用密码列表响应
type GetAuthCodeListResponse struct {
	ErrCode      int32             `json:"errcode"`
	ErrMsg       string            `json:"errmsg"`
	AuthCodeList []*AuthCodeDetail `json:"auth_code_list,omitempty"` // 客户端专用密码列表
}

// DeleteAuthCodeRequest 删除客户端专用密码请求
type DeleteAuthCodeRequest struct {
	ID         uint32 `json:"id"`           // 公共邮箱ID(必填)
	AuthCodeID uint32 `json:"auth_code_id"` // 客户端专用密码ID(必填)
}

// DeleteAuthCodeResponse 删除客户端专用密码响应
type DeleteAuthCodeResponse struct {
	ErrCode int32  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
