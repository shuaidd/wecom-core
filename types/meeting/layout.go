package meeting

import "github.com/shuaidd/wecom-core/types/common"

// ==================== 布局相关类型 ====================

// ListLayoutRequest 获取会议布局列表请求
type ListLayoutRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
}

// ListLayoutResponse 获取会议布局列表响应
type ListLayoutResponse struct {
	common.Response
	// SelectedLayoutID 会议应用的布局ID
	SelectedLayoutID string `json:"selected_layout_id,omitempty"`
	// LayoutList 布局对象列表
	LayoutList []Layout `json:"layout_list,omitempty"`
}

// Layout 布局对象
type Layout struct {
	// LayoutID 布局ID
	LayoutID string `json:"layout_id,omitempty"`
	// LayoutName 布局名称
	LayoutName string `json:"layout_name,omitempty"`
	// PageList 布局单页对象列表
	PageList []LayoutPage `json:"page_list,omitempty"`
}

// LayoutPage 布局单页对象
type LayoutPage struct {
	// LayoutTemplateID 布局模板ID
	LayoutTemplateID string `json:"layout_template_id"`
	// EnablePolling 开启或关闭轮询，默认关闭
	EnablePolling bool `json:"enable_polling,omitempty"`
	// PollingSetting 轮询参数设置对象
	PollingSetting *PollingSetting `json:"polling_setting,omitempty"`
	// UserSeatList 用户座次对象列表
	UserSeatList []UserSeat `json:"user_seat_list,omitempty"`
}

// PollingSetting 轮询参数设置
type PollingSetting struct {
	// PollingIntervalUnit 轮询间隔时间类型：1-秒；2-分钟
	PollingIntervalUnit uint32 `json:"polling_interval_unit,omitempty"`
	// PollingInterval 轮询间隔时长，允许取值范围1～999999
	PollingInterval uint32 `json:"polling_interval,omitempty"`
	// IgnoreUserNoVideo 设置是否忽略没开启视频成员
	IgnoreUserNoVideo bool `json:"ignore_user_novideo,omitempty"`
	// IgnoreUserAbsence 设置是否忽略未入会成员
	IgnoreUserAbsence bool `json:"ignore_user_absence,omitempty"`
}

// UserSeat 用户座次对象
type UserSeat struct {
	// GridID 宫格ID
	GridID string `json:"grid_id"`
	// GridType 宫格类型：1-视频画面；2-共享画面；3-拓展应用
	GridType uint32 `json:"grid_type"`
	// VideoType 视频画面来源：1-演讲者；2-自动填充；3-指定人员
	VideoType uint32 `json:"video_type,omitempty"`
	// UserList 宫格中的用户列表
	UserList []GridUser `json:"user_list,omitempty"`
	// UserID 当场会议的企业成员的userid
	UserID string `json:"userid,omitempty"`
	// TmpOpenID 当场会议的用户临时ID
	TmpOpenID string `json:"tmp_openid,omitempty"`
	// NickName 视频画面展示昵称
	NickName string `json:"nick_name,omitempty"`
	// ToolSDKID 拓展应用ID
	ToolSDKID string `json:"tool_sdkid,omitempty"`
}

// GridUser 宫格中的用户
type GridUser struct {
	// UserID 本场会议企业成员的userid
	UserID string `json:"userid,omitempty"`
	// TmpOpenID 用户当前会议临时身份ID，单场会议唯一
	TmpOpenID string `json:"tmp_openid,omitempty"`
	// NickName 用于视频画面展示的昵称(base64)
	NickName string `json:"nick_name,omitempty"`
}

// AddBasicLayoutRequest 添加会议基础布局请求
type AddBasicLayoutRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// LayoutList 布局对象列表
	LayoutList []BasicLayoutInput `json:"layout_list"`
	// DefaultLayoutOrder 布局列表中会议需要应用的布局序号，从1开始计数
	DefaultLayoutOrder uint32 `json:"default_layout_order,omitempty"`
}

// BasicLayoutInput 基础布局输入
type BasicLayoutInput struct {
	// PageList 布局单页对象列表
	PageList []BasicLayoutPageInput `json:"page_list"`
}

// BasicLayoutPageInput 基础布局单页输入
type BasicLayoutPageInput struct {
	// LayoutTemplateID 布局模板ID
	LayoutTemplateID string `json:"layout_template_id"`
	// UserSeatList 用户座次对象列表
	UserSeatList []UserSeat `json:"user_seat_list,omitempty"`
}

// AddBasicLayoutResponse 添加会议基础布局响应
type AddBasicLayoutResponse struct {
	common.Response
	// SelectedLayoutID 会议应用的布局ID
	SelectedLayoutID string `json:"selected_layout_id,omitempty"`
	// LayoutList 布局对象列表
	LayoutList []BasicLayoutOutput `json:"layout_list,omitempty"`
}

// BasicLayoutOutput 基础布局输出
type BasicLayoutOutput struct {
	// LayoutID 布局ID
	LayoutID string `json:"layout_id,omitempty"`
	// PageList 布局单页对象列表
	PageList []BasicLayoutPageOutput `json:"page_list,omitempty"`
}

// BasicLayoutPageOutput 基础布局单页输出
type BasicLayoutPageOutput struct {
	// LayoutTemplateID 布局模板ID
	LayoutTemplateID string `json:"layout_template_id,omitempty"`
	// UserSeatList 用户座次对象列表
	UserSeatList []UserSeat `json:"user_seat_list,omitempty"`
}

// AddAdvancedLayoutRequest 添加会议高级布局请求
type AddAdvancedLayoutRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// LayoutList 布局对象列表
	LayoutList []AdvancedLayoutInput `json:"layout_list"`
}

// AdvancedLayoutInput 高级布局输入
type AdvancedLayoutInput struct {
	// LayoutName 布局名称
	LayoutName string `json:"layout_name,omitempty"`
	// PageList 布局单页对象列表
	PageList []LayoutPage `json:"page_list"`
}

// AddAdvancedLayoutResponse 添加会议高级布局响应
type AddAdvancedLayoutResponse struct {
	common.Response
	// LayoutList 布局对象列表
	LayoutList []Layout `json:"layout_list,omitempty"`
}

// UpdateBasicLayoutRequest 修改会议基础布局请求
type UpdateBasicLayoutRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// LayoutID 布局ID
	LayoutID string `json:"layout_id"`
	// PageList 布局单页对象列表
	PageList []BasicLayoutPageInput `json:"page_list"`
	// EnableSetDefault 是否设置为会议应用的布局，默认不设置
	EnableSetDefault bool `json:"enable_set_default,omitempty"`
}

// UpdateAdvancedLayoutRequest 修改会议高级布局请求
type UpdateAdvancedLayoutRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// LayoutID 布局ID
	LayoutID string `json:"layout_id"`
	// LayoutName 布局名称
	LayoutName string `json:"layout_name,omitempty"`
	// PageList 布局单页对象列表
	PageList []LayoutPage `json:"page_list"`
}

// BatchDeleteLayoutRequest 批量删除布局请求
type BatchDeleteLayoutRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// LayoutIDList 布局ID列表，要删除的一个或多个布局ID（最多支持20个）
	LayoutIDList []string `json:"layout_id_list"`
}

// SetDefaultLayoutRequest 设置会议默认布局请求
type SetDefaultLayoutRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// SelectedLayoutID 会议应用的布局ID（若送空""，表示恢复成会议自带的默认原始布局）
	SelectedLayoutID string `json:"selected_layout_id"`
}

// ApplyAdvancedLayoutRequest 设置高级布局请求
type ApplyAdvancedLayoutRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// LayoutID 选择应用的布局ID（若传空""，表示恢复成当前会议的默认布局）
	LayoutID string `json:"layout_id"`
	// UserList 用户列表对象数组。如果该字段为空，为会议设置高级自定义布局；如果该字段携带用户，则只为指定用户设置个性布局。单次最多支持20个用户。
	UserList []ApplyLayoutUser `json:"user_list,omitempty"`
}

// ApplyLayoutUser 应用布局的用户
type ApplyLayoutUser struct {
	// TmpOpenID 用户当前会议中的临时身份ID，单场会议唯一。仅对H.323/SIP 会议室终端生效
	TmpOpenID string `json:"tmp_openid"`
}

// GetUserLayoutRequest 获取用户布局请求
type GetUserLayoutRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// TmpOpenID 被操作用户临时身份ID
	TmpOpenID string `json:"tmp_openid"`
	// InstanceID 被操作用户终端设备类型ID
	InstanceID uint32 `json:"instance_id"`
}

// GetUserLayoutResponse 获取用户布局响应
type GetUserLayoutResponse struct {
	common.Response
	// SelectedLayoutID 会议应用的布局ID
	SelectedLayoutID string `json:"selected_layout_id,omitempty"`
	// LayoutName 布局名称
	LayoutName string `json:"layout_name,omitempty"`
	// LayoutType 布局类型：0-默认布局；2-自定义会议布局；3-个性布局
	LayoutType uint32 `json:"layout_type,omitempty"`
	// PageList 布局单页对象列表
	PageList []LayoutPage `json:"page_list,omitempty"`
}

// ListLayoutTemplateResponse 获取布局模板列表响应
type ListLayoutTemplateResponse struct {
	common.Response
	// LayoutTemplateList 布局模板对象列表
	LayoutTemplateList []LayoutTemplate `json:"layout_template_list,omitempty"`
}

// LayoutTemplate 布局模板对象
type LayoutTemplate struct {
	// LayoutTemplateID 布局模板ID
	LayoutTemplateID string `json:"layout_template_id,omitempty"`
	// ThumbnailURL 缩略图URL
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	// PictureURL 布局图URL
	PictureURL string `json:"picture_url,omitempty"`
	// RenderRule 渲染规则
	RenderRule string `json:"render_rule,omitempty"`
}

// ==================== 背景相关类型 ====================

// ListBackgroundRequest 获取会议背景列表请求
type ListBackgroundRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
}

// ListBackgroundResponse 获取会议背景列表响应
type ListBackgroundResponse struct {
	common.Response
	// SelectedBackgroundID 会议应用的背景ID
	SelectedBackgroundID string `json:"selected_background_id,omitempty"`
	// BackgroundList 背景对象列表
	BackgroundList []Background `json:"background_list,omitempty"`
}

// Background 背景对象
type Background struct {
	// BackgroundID 背景ID
	BackgroundID string `json:"background_id,omitempty"`
	// ImageMD5 背景图片MD5（图片内容MD5的十六进制表示）
	ImageMD5 string `json:"image_md5,omitempty"`
}

// AddBackgroundRequest 添加会议背景请求
type AddBackgroundRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// ImageList 图片对象列表
	ImageList []BackgroundImage `json:"image_list"`
	// DefaultImageOrder 图片列表中会议需要使用的背景图片序号，从1开始计数。不填默认为1
	DefaultImageOrder uint32 `json:"default_image_order,omitempty"`
}

// BackgroundImage 背景图片对象
type BackgroundImage struct {
	// ImageMD5 背景图片MD5（图片内容MD5的十六进制表示）
	ImageMD5 string `json:"image_md5"`
	// ImageURL 背景图片URL
	ImageURL string `json:"image_url"`
}

// AddBackgroundResponse 添加会议背景响应
type AddBackgroundResponse struct {
	common.Response
	// SelectedBackgroundID 会议应用的背景ID
	SelectedBackgroundID string `json:"selected_background_id,omitempty"`
	// BackgroundList 背景对象列表
	BackgroundList []Background `json:"background_list,omitempty"`
}

// DeleteBackgroundRequest 删除会议背景请求
type DeleteBackgroundRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// BackgroundID 背景ID
	BackgroundID string `json:"background_id"`
}

// BatchDeleteBackgroundRequest 批量删除会议背景请求
type BatchDeleteBackgroundRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// BackgroundIDList 背景ID列表
	BackgroundIDList []string `json:"background_id_list"`
}

// SetDefaultBackgroundRequest 设置会议默认背景请求
type SetDefaultBackgroundRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// SelectedBackgroundID 会议应用的背景ID（若送空""，则表示恢复成会议默认的黑色背景）
	SelectedBackgroundID string `json:"selected_background_id"`
}
