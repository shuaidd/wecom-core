package agent

// CreateMenuRequest 创建菜单请求
type CreateMenuRequest struct {
	AgentID int          `json:"-" url:"agentid"` // 企业应用的id
	Button  []MenuButton `json:"button"`          // 一级菜单数组，个数应为1~3个
}

// MenuButton 菜单按钮
type MenuButton struct {
	Type      string       `json:"type,omitempty"`       // 菜单的响应动作类型
	Name      string       `json:"name"`                 // 菜单的名字
	Key       string       `json:"key,omitempty"`        // 菜单KEY值，用于消息接口推送，不超过128字节
	URL       string       `json:"url,omitempty"`        // 网页链接
	PagePath  string       `json:"pagepath,omitempty"`   // 小程序的页面路径
	AppID     string       `json:"appid,omitempty"`      // 小程序的appid
	SubButton []MenuButton `json:"sub_button,omitempty"` // 二级菜单数组，个数应为1~5个
}

// CreateMenuResponse 创建菜单响应
type CreateMenuResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// GetMenuRequest 获取菜单请求
type GetMenuRequest struct {
	AgentID int `json:"-" url:"agentid"` // 应用id
}

// GetMenuResponse 获取菜单响应
type GetMenuResponse struct {
	ErrCode int          `json:"errcode"`
	ErrMsg  string       `json:"errmsg"`
	Button  []MenuButton `json:"button"` // 菜单数组
}

// DeleteMenuRequest 删除菜单请求
type DeleteMenuRequest struct {
	AgentID int `json:"-" url:"agentid"` // 应用id
}

// DeleteMenuResponse 删除菜单响应
type DeleteMenuResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
