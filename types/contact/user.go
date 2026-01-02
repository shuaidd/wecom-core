package contact

// ExtAttr 扩展属性
type ExtAttr struct {
	Attrs []Attr `json:"attrs,omitempty"`
}

// Attr 属性
type Attr struct {
	Type        int          `json:"type"`
	Name        string       `json:"name"`
	Text        *TextAttr    `json:"text,omitempty"`
	Web         *WebAttr     `json:"web,omitempty"`
	MiniProgram *MiniProgram `json:"miniprogram,omitempty"`
}

// TextAttr 文本属性
type TextAttr struct {
	Value string `json:"value"`
}

// WebAttr 网页属性
type WebAttr struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

// MiniProgram 小程序属性
type MiniProgram struct {
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath"`
	Title    string `json:"title"`
}

// ExternalProfile 对外属性
type ExternalProfile struct {
	ExternalCorpName string       `json:"external_corp_name,omitempty"`
	WechatChannels   *WechatChan  `json:"wechat_channels,omitempty"`
	ExternalAttr     []Attr       `json:"external_attr,omitempty"`
}

// WechatChan 视频号
type WechatChan struct {
	Nickname string `json:"nickname,omitempty"`
	Status   int    `json:"status,omitempty"`
}

// User 成员信息
type User struct {
	// UserID 成员UserID
	UserID string `json:"userid"`
	// Name 成员名称
	Name string `json:"name"`
	// Mobile 手机号码
	Mobile string `json:"mobile,omitempty"`
	// Department 成员所属部门id列表
	Department []int `json:"department,omitempty"`
	// Order 部门内的排序值
	Order []int `json:"order,omitempty"`
	// Position 职务信息
	Position string `json:"position,omitempty"`
	// Gender 性别。0表示未定义，1表示男性，2表示女性
	Gender string `json:"gender,omitempty"`
	// Email 邮箱
	Email string `json:"email,omitempty"`
	// BizMail 企业邮箱
	BizMail string `json:"biz_mail,omitempty"`
	// IsLeaderInDept 是否为部门负责人
	IsLeaderInDept []int `json:"is_leader_in_dept,omitempty"`
	// DirectLeader 直属上级UserID
	DirectLeader []string `json:"direct_leader,omitempty"`
	// Avatar 头像url
	Avatar string `json:"avatar,omitempty"`
	// ThumbAvatar 头像缩略图url
	ThumbAvatar string `json:"thumb_avatar,omitempty"`
	// Telephone 座机
	Telephone string `json:"telephone,omitempty"`
	// Alias 别名
	Alias string `json:"alias,omitempty"`
	// Address 地址
	Address string `json:"address,omitempty"`
	// OpenUserID 全局唯一ID
	OpenUserID string `json:"open_userid,omitempty"`
	// MainDepartment 主部门
	MainDepartment int `json:"main_department,omitempty"`
	// ExtAttr 扩展属性
	ExtAttr *ExtAttr `json:"extattr,omitempty"`
	// Status 激活状态: 1=已激活，2=已禁用，4=未激活，5=退出企业
	Status int `json:"status,omitempty"`
	// QRCode 员工个人二维码
	QRCode string `json:"qr_code,omitempty"`
	// ExternalPosition 对外职务
	ExternalPosition string `json:"external_position,omitempty"`
	// ExternalProfile 成员对外属性
	ExternalProfile *ExternalProfile `json:"external_profile,omitempty"`
}
