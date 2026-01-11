package agent

// WorkbenchTemplateType 工作台模板类型
type WorkbenchTemplateType string

const (
	// WorkbenchTypeNormal 普通模式
	WorkbenchTypeNormal WorkbenchTemplateType = "normal"
	// WorkbenchTypeKeydata 关键数据型
	WorkbenchTypeKeydata WorkbenchTemplateType = "keydata"
	// WorkbenchTypeImage 图片型
	WorkbenchTypeImage WorkbenchTemplateType = "image"
	// WorkbenchTypeList 列表型
	WorkbenchTypeList WorkbenchTemplateType = "list"
	// WorkbenchTypeWebview webview型
	WorkbenchTypeWebview WorkbenchTemplateType = "webview"
)

// KeydataItem 关键数据项
type KeydataItem struct {
	Key      string `json:"key,omitempty"`      // 关键数据名称
	Data     string `json:"data"`               // 关键数据
	JumpURL  string `json:"jump_url,omitempty"` // 点击跳转url
	PagePath string `json:"pagepath,omitempty"` // 小程序pagepath
}

// KeydataTemplate 关键数据型模板
type KeydataTemplate struct {
	Items []KeydataItem `json:"items"` // 关键数据型数组，不超过4个
}

// ImageTemplate 图片型模板
type ImageTemplate struct {
	URL      string `json:"url"`                // 图片url
	JumpURL  string `json:"jump_url,omitempty"` // 点击跳转url
	PagePath string `json:"pagepath,omitempty"` // 小程序pagepath
}

// ListItem 列表项
type ListItem struct {
	Title    string `json:"title"`              // 列表显示文字
	JumpURL  string `json:"jump_url,omitempty"` // 点击跳转url
	PagePath string `json:"pagepath,omitempty"` // 小程序pagepath
}

// ListTemplate 列表型模板
type ListTemplate struct {
	Items []ListItem `json:"items"` // 列表型数组，不超过3个
}

// WebviewTemplate webview型模板
type WebviewTemplate struct {
	URL                string `json:"url"`                            // 渲染展示的url
	JumpURL            string `json:"jump_url,omitempty"`             // 点击跳转url
	PagePath           string `json:"pagepath,omitempty"`             // 小程序pagepath
	Height             string `json:"height,omitempty"`               // 高度，single_row或double_row
	HideTitle          bool   `json:"hide_title,omitempty"`           // 是否隐藏标题
	EnableWebviewClick bool   `json:"enable_webview_click,omitempty"` // 是否开启webview内的链接跳转能力
}

// SetWorkbenchTemplateRequest 设置应用在工作台展示的模板请求
type SetWorkbenchTemplateRequest struct {
	AgentID         int                   `json:"agentid"`                     // 应用id
	Type            WorkbenchTemplateType `json:"type"`                        // 模板类型
	Keydata         *KeydataTemplate      `json:"keydata,omitempty"`           // 关键数据型模板数据
	Image           *ImageTemplate        `json:"image,omitempty"`             // 图片型模板数据
	List            *ListTemplate         `json:"list,omitempty"`              // 列表型模板数据
	Webview         *WebviewTemplate      `json:"webview,omitempty"`           // webview型模板数据
	ReplaceUserData bool                  `json:"replace_user_data,omitempty"` // 是否覆盖用户工作台的数据
}

// SetWorkbenchTemplateResponse 设置应用在工作台展示的模板响应
type SetWorkbenchTemplateResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// GetWorkbenchTemplateRequest 获取应用在工作台展示的模板请求
type GetWorkbenchTemplateRequest struct {
	AgentID int `json:"agentid"` // 应用id
}

// GetWorkbenchTemplateResponse 获取应用在工作台展示的模板响应
type GetWorkbenchTemplateResponse struct {
	ErrCode         int                   `json:"errcode"`
	ErrMsg          string                `json:"errmsg"`
	Type            WorkbenchTemplateType `json:"type"`                        // 模板类型
	Keydata         *KeydataTemplate      `json:"keydata,omitempty"`           // 关键数据型模板数据
	Image           *ImageTemplate        `json:"image,omitempty"`             // 图片型模板数据
	List            *ListTemplate         `json:"list,omitempty"`              // 列表型模板数据
	Webview         *WebviewTemplate      `json:"webview,omitempty"`           // webview型模板数据
	ReplaceUserData bool                  `json:"replace_user_data,omitempty"` // 是否覆盖用户工作台的数据
}

// SetWorkbenchDataRequest 设置应用在用户工作台展示的数据请求
type SetWorkbenchDataRequest struct {
	AgentID int                   `json:"agentid"`           // 应用id
	UserID  string                `json:"userid"`            // 用户userid
	Type    WorkbenchTemplateType `json:"type"`              // 模板类型
	Keydata *KeydataTemplate      `json:"keydata,omitempty"` // 关键数据型模板数据
	Image   *ImageTemplate        `json:"image,omitempty"`   // 图片型模板数据
	List    *ListTemplate         `json:"list,omitempty"`    // 列表型模板数据
	Webview *WebviewTemplate      `json:"webview,omitempty"` // webview型模板数据
}

// SetWorkbenchDataResponse 设置应用在用户工作台展示的数据响应
type SetWorkbenchDataResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// WorkbenchUserData 用户工作台数据
type WorkbenchUserData struct {
	Type    WorkbenchTemplateType `json:"type"`              // 模板类型
	Keydata *KeydataTemplate      `json:"keydata,omitempty"` // 关键数据型模板数据
	Image   *ImageTemplate        `json:"image,omitempty"`   // 图片型模板数据
	List    *ListTemplate         `json:"list,omitempty"`    // 列表型模板数据
	Webview *WebviewTemplate      `json:"webview,omitempty"` // webview型模板数据
}

// BatchSetWorkbenchDataRequest 批量设置应用在用户工作台展示的数据请求
type BatchSetWorkbenchDataRequest struct {
	AgentID    int                `json:"agentid"`     // 应用id
	UserIDList []string           `json:"userid_list"` // 用户userid列表，最多1000个
	Data       *WorkbenchUserData `json:"data"`        // 用户设置的数据
}

// BatchSetWorkbenchDataResponse 批量设置应用在用户工作台展示的数据响应
type BatchSetWorkbenchDataResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// GetWorkbenchDataRequest 获取应用在用户工作台展示的数据请求
type GetWorkbenchDataRequest struct {
	AgentID int    `json:"agentid"` // 应用id
	UserID  string `json:"userid"`  // 用户userid
}

// GetWorkbenchDataResponse 获取应用在用户工作台展示的数据响应
type GetWorkbenchDataResponse struct {
	ErrCode int                `json:"errcode"`
	ErrMsg  string             `json:"errmsg"`
	Data    *WorkbenchUserData `json:"data"` // 用户设置的数据
}
