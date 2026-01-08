package calendar

import "github.com/shuaidd/wecom-core/internal/client"

// Calendar 日历信息
type Calendar struct {
	CalID          string        `json:"cal_id,omitempty"`           // 日历ID
	Admins         []string      `json:"admins,omitempty"`           // 日历的管理员userid列表
	SetAsDefault   int           `json:"set_as_default,omitempty"`   // 是否将该日历设置为默认日历。0-否；1-是
	Summary        string        `json:"summary"`                    // 日历标题。1 ~ 128 字符
	Color          string        `json:"color"`                      // 日历颜色，RGB颜色编码16进制表示
	Description    string        `json:"description,omitempty"`      // 日历描述。0 ~ 512 字符
	IsPublic       int           `json:"is_public,omitempty"`        // 是否公共日历。0-否；1-是
	PublicRange    *PublicRange  `json:"public_range,omitempty"`     // 公开范围
	IsCorpCalendar int           `json:"is_corp_calendar,omitempty"` // 是否全员日历。0-否；1-是
	Shares         []Share       `json:"shares,omitempty"`           // 日历通知范围成员列表
}

// PublicRange 公开范围
type PublicRange struct {
	UserIDs  []string `json:"userids,omitempty"`  // 公开的成员列表范围
	PartyIDs []int    `json:"partyids,omitempty"` // 公开的部门列表范围
}

// Share 日历通知范围成员
type Share struct {
	UserID     string `json:"userid"`               // 日历通知范围成员的id
	Permission int    `json:"permission,omitempty"` // 日历通知范围成员权限。1：可查看；3：仅查看闲忙状态
}

// FailedShare 无效的日历通知范围成员
type FailedShare struct {
	ErrCode int    `json:"errcode"` // 错误码
	ErrMsg  string `json:"errmsg"`  // 错误码说明
	UserID  string `json:"userid"`  // 日历通知范围成员的id
}

// FailResult 无效的输入内容
type FailResult struct {
	Shares []FailedShare `json:"shares,omitempty"` // 无效的日历通知范围成员列表
}

// CreateCalendarRequest 创建日历请求
type CreateCalendarRequest struct {
	Calendar Calendar `json:"calendar"` // 日历信息
	AgentID  int      `json:"agentid,omitempty"` // 授权方安装的应用agentid
}

// CreateCalendarResponse 创建日历响应
type CreateCalendarResponse struct {
	client.CommonResponse
	CalID      string     `json:"cal_id"`                // 日历ID
	FailResult FailResult `json:"fail_result,omitempty"` // 无效的输入内容
}

// GetCalendarRequest 获取日历详情请求
type GetCalendarRequest struct {
	CalIDList []string `json:"cal_id_list"` // 日历ID列表，一次最多可获取1000条
}

// GetCalendarResponse 获取日历详情响应
type GetCalendarResponse struct {
	client.CommonResponse
	CalendarList []Calendar `json:"calendar_list"` // 日历列表
}

// UpdateCalendarRequest 更新日历请求
type UpdateCalendarRequest struct {
	SkipPublicRange int      `json:"skip_public_range,omitempty"` // 是否不更新可订阅范围。0-否；1-是
	Calendar        Calendar `json:"calendar"`                    // 日历信息
}

// UpdateCalendarResponse 更新日历响应
type UpdateCalendarResponse struct {
	client.CommonResponse
	FailResult FailResult `json:"fail_result,omitempty"` // 无效的输入内容
}

// DeleteCalendarRequest 删除日历请求
type DeleteCalendarRequest struct {
	CalID string `json:"cal_id"` // 日历ID
}
