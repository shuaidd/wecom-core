package email

// GetNewCountRequest 获取邮件未读数请求
type GetNewCountRequest struct {
	UserID string `json:"userid"` // 成员UserID
}

// GetNewCountResponse 获取邮件未读数响应
type GetNewCountResponse struct {
	ErrCode int32  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Count   int32  `json:"count"` // 成员邮箱中邮件未读数
}

// GetMailListRequest 获取收件箱邮件列表请求
type GetMailListRequest struct {
	BeginTime uint32 `json:"begin_time"`       // 开始时间，unix时间戳
	EndTime   uint32 `json:"end_time"`         // 结束时间，unix时间戳
	Cursor    string `json:"cursor,omitempty"` // 上一次调用时返回的next_cursor，第一次拉取可以不填
	Limit     uint32 `json:"limit,omitempty"`  // 期望请求的数据量，默认值为100，最大值为1000
}

// MailItem 邮件项
type MailItem struct {
	MailID string `json:"mail_id"` // 邮件id
}

// GetMailListResponse 获取收件箱邮件列表响应
type GetMailListResponse struct {
	ErrCode    int32      `json:"errcode"`
	ErrMsg     string     `json:"errmsg"`
	NextCursor string     `json:"next_cursor"` // 下一次请求的cursor值
	HasMore    uint32     `json:"has_more"`    // 是否还有更多数据。0-没有 1-有
	MailList   []MailItem `json:"mail_list"`   // 邮件列表
}

// ReadMailRequest 获取邮件内容请求
type ReadMailRequest struct {
	MailID string `json:"mail_id"` // 邮件id
}

// ReadMailResponse 获取邮件内容响应
type ReadMailResponse struct {
	ErrCode  int32  `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	MailData string `json:"mail_data"` // 邮件eml内容
}
