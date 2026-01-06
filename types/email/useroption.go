package email

// UserOptionItem 用户功能属性项
type UserOptionItem struct {
	Type  int32  `json:"type"`  // 功能设置属性类型 1: 强制启用安全登录 2: IMAP/SMTP服务 3: POP/SMTP服务 4: 是否启用安全登录
	Value string `json:"value"` // 1表示启用，0表示关闭
}

// UserOption 用户功能属性
type UserOption struct {
	List []UserOptionItem `json:"list"` // 功能属性列表
}

// GetUserOptionRequest 获取用户功能属性请求
type GetUserOptionRequest struct {
	UserID string   `json:"userid"` // 用户UserID
	Type   []uint32 `json:"type"`   // 功能设置属性类型 1: 强制启用安全登录 2: IMAP/SMTP服务 3: POP/SMTP服务 4: 是否启用安全登录
}

// GetUserOptionResponse 获取用户功能属性响应
type GetUserOptionResponse struct {
	ErrCode int32      `json:"errcode"`
	ErrMsg  string     `json:"errmsg"`
	Option  UserOption `json:"option"` // 用户功能属性
}

// UpdateUserOptionRequest 更改用户功能属性请求
type UpdateUserOptionRequest struct {
	UserID string     `json:"userid"` // 用户UserID
	Option UserOption `json:"option"` // 用户功能属性
}

// UpdateUserOptionResponse 更改用户功能属性响应
type UpdateUserOptionResponse struct {
	ErrCode int32  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
