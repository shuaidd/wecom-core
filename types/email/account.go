package email

// GetAppEmailAliasResponse 查询应用邮箱账号响应
type GetAppEmailAliasResponse struct {
	ErrCode   int32    `json:"errcode"`
	ErrMsg    string   `json:"errmsg"`
	Email     string   `json:"email"`      // 当前发信账号的主邮箱地址
	AliasList []string `json:"alias_list"` // 别名邮箱地址列表
}

// UpdateAppEmailAliasRequest 更新应用邮箱账号请求
type UpdateAppEmailAliasRequest struct {
	NewEmail string `json:"new_email"` // 修改后的应用邮箱账号(必填)
}

// UpdateAppEmailAliasResponse 更新应用邮箱账号响应
type UpdateAppEmailAliasResponse struct {
	ErrCode int32  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// ActEmailRequest 禁用/启用邮箱账号请求
type ActEmailRequest struct {
	UserID        string `json:"userid,omitempty"`         // 成员UserID,userid与publicemail_id至少应该传一项，同时传则只操作userid。不可禁用超管与企业创建人
	PublicEmailID int32  `json:"publicemail_id,omitempty"` // 业务邮箱ID,userid与publicemail_id至少应该传一项，同时传则只操作userid
	Type          int32  `json:"type"`                     // 1启用，2禁用
}

// ActEmailResponse 禁用/启用邮箱账号响应
type ActEmailResponse struct {
	ErrCode int32  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
