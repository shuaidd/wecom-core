package agent

// GetAgentRequest 获取应用详情请求
type GetAgentRequest struct {
	AgentID int `json:"agentid" url:"agentid"` // 应用id
}

// GetAgentResponse 获取应用详情响应
type GetAgentResponse struct {
	ErrCode                 int             `json:"errcode"`
	ErrMsg                  string          `json:"errmsg"`
	AgentID                 int             `json:"agentid"`                   // 企业应用id
	Name                    string          `json:"name"`                      // 企业应用名称
	SquareLogoURL           string          `json:"square_logo_url"`           // 企业应用方形头像
	Description             string          `json:"description"`               // 企业应用详情
	AllowUserInfos          *AllowUserInfos `json:"allow_userinfos"`           // 企业应用可见范围（人员）
	AllowPartys             *AllowPartys    `json:"allow_partys"`              // 企业应用可见范围（部门）
	AllowTags               *AllowTags      `json:"allow_tags"`                // 企业应用可见范围（标签）
	Close                   int             `json:"close"`                     // 企业应用是否被停用。0：未被停用；1：被停用
	RedirectDomain          string          `json:"redirect_domain"`           // 企业应用可信域名
	ReportLocationFlag      int             `json:"report_location_flag"`      // 企业应用是否打开地理位置上报 0：不上报；1：进入会话上报
	IsReportEnter           int             `json:"isreportenter"`             // 是否上报用户进入应用事件。0：不接收；1：接收
	HomeURL                 string          `json:"home_url"`                  // 应用主页url
	CustomizedPublishStatus int             `json:"customized_publish_status"` // 代开发自建应用返回该字段
}

// AllowUserInfos 可见范围（人员）
type AllowUserInfos struct {
	User []UserInfo `json:"user"`
}

// UserInfo 用户信息
type UserInfo struct {
	UserID string `json:"userid"`
}

// AllowPartys 可见范围（部门）
type AllowPartys struct {
	PartyID []int `json:"partyid"`
}

// AllowTags 可见范围（标签）
type AllowTags struct {
	TagID []int `json:"tagid"`
}

// ListAgentResponse 获取应用列表响应
type ListAgentResponse struct {
	ErrCode   int         `json:"errcode"`
	ErrMsg    string      `json:"errmsg"`
	AgentList []AgentItem `json:"agentlist"` // 应用列表
}

// AgentItem 应用项
type AgentItem struct {
	AgentID       int    `json:"agentid"`         // 企业应用id
	Name          string `json:"name"`            // 企业应用名称
	SquareLogoURL string `json:"square_logo_url"` // 企业应用方形头像url
}

// SetAgentRequest 设置应用请求
type SetAgentRequest struct {
	AgentID            int    `json:"agentid"`                        // 企业应用的id
	ReportLocationFlag *int   `json:"report_location_flag,omitempty"` // 企业应用是否打开地理位置上报 0：不上报；1：进入会话上报
	LogoMediaID        string `json:"logo_mediaid,omitempty"`         // 企业应用头像的mediaid
	Name               string `json:"name,omitempty"`                 // 企业应用名称，长度不超过32个utf8字符
	Description        string `json:"description,omitempty"`          // 企业应用详情，长度为4至120个utf8字符
	RedirectDomain     string `json:"redirect_domain,omitempty"`      // 企业应用可信域名
	IsReportEnter      *int   `json:"isreportenter,omitempty"`        // 是否上报用户进入应用事件。0：不接收；1：接收
	HomeURL            string `json:"home_url,omitempty"`             // 应用主页url
}

// SetAgentResponse 设置应用响应
type SetAgentResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
