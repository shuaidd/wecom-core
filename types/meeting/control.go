package meeting

// OperatedUser 被操作成员信息
type OperatedUser struct {
	// TmpOpenID 被操作者的会中临时ID
	TmpOpenID string `json:"tmp_openid"`
	// InstanceID 成员的终端设备类型
	InstanceID int32 `json:"instance_id"`
	// Nickname 成员昵称字符串，限制20个字符（仅用于设置昵称）
	Nickname string `json:"nickname,omitempty"`
}

// KickoutUsersRequest 移出成员请求
type KickoutUsersRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// AllowRejoin 是否允许再次入会
	AllowRejoin bool `json:"allow_rejoin"`
	// OperatedUsers 被操作对象列表
	OperatedUsers []OperatedUser `json:"operated_users"`
}

// SetNicknamesRequest 修改成员在会中显示的昵称请求
type SetNicknamesRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// OperatedUsers 被操作对象列表
	OperatedUsers []OperatedUser `json:"operated_users"`
}

// MuteUserRequest 静音成员请求
type MuteUserRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// Option 操作类型：true：静音；false：解除静音
	Option bool `json:"option"`
	// OperatedUser 被操作对象
	OperatedUser OperatedUser `json:"operated_user"`
}

// DismissMeetingRequest 结束会议请求
type DismissMeetingRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// ForceDismiss 是否强制结束会议，默认值为1：
	// 0：不强制结束会议，会议中有参会者，则无法强制结束会议
	// 1：强制结束会议，会议中有参会者，也会强制结束会议
	ForceDismiss int32 `json:"force_dismiss,omitempty"`
	// RetrieveCode 是否回收会议号，默认值为0：
	// 0：不回收会议号，可以重新入会
	// 1：回收会议号，不可重新入会
	RetrieveCode int32 `json:"retrieve_code,omitempty"`
}

// SetCohostRequest 管理联席主持人请求
type SetCohostRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// Action 具体设置动作：true：设置联席主持人；false：撤销联席主持人
	Action bool `json:"action"`
	// OperatedUser 被操作成员
	OperatedUser OperatedUser `json:"operated_user"`
}

// MeetingSettingsRequest 管理会中设置请求
type MeetingSettingsRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// MuteAll 是否全体静音：true：全体静音；false：关闭全体静音
	MuteAll bool `json:"mute_all,omitempty"`
	// AllowUnmuteSelf 是否允许成员自己取消静音，请求参数 MuteAll 必传，且 MuteAll = true 时设置才生效
	AllowUnmuteSelf bool `json:"allow_unmute_self,omitempty"`
	// EnableEnterMute 成员入会静音：0：关闭静音；1：开启静音；2：超过6人自动开启静音
	EnableEnterMute int32 `json:"enable_enter_mute,omitempty"`
	// MeetingLocked 是否锁定会议：true：锁定；false：关闭锁定
	MeetingLocked bool `json:"meeting_locked,omitempty"`
	// HideMeetingCodePassword 隐藏会议号和密码：true：隐藏；false：不隐藏
	HideMeetingCodePassword bool `json:"hide_meeting_code_password,omitempty"`
	// AllowChat 允许参会者聊天设置：0：允许参会者自由聊天；1：仅允许参会者公开聊天；2：仅允许私聊主持人
	AllowChat int32 `json:"allow_chat,omitempty"`
	// AllowShareScreen 是否允许参会者发起屏幕共享：true：允许；false：不允许
	AllowShareScreen bool `json:"allow_share_screen,omitempty"`
	// AllowExternalUser 是否仅企业成员可入会：true：仅企业成员可入会；false：不限制
	AllowExternalUser bool `json:"allow_external_user,omitempty"`
	// PlayIvrOnJoin 成员入会是否播放提示音：true：成员入会播放提示音；false：不播放
	PlayIvrOnJoin bool `json:"play_ivr_on_join,omitempty"`
	// EnableWaitingRoom 是否开启等候室：true：开启；false：关闭
	EnableWaitingRoom bool `json:"enable_waiting_room,omitempty"`
}

// ManageWaitingRoomUsersRequest 管理等候室成员请求
type ManageWaitingRoomUsersRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// OperateType 操作类型：1：主持人将等候室成员移入会议；2：主持人将会中成员移入等候室；3：主持人将等候室成员移出等候室
	OperateType int32 `json:"operate_type"`
	// AllowRejoin 移出成员后是否允许其再次加入会议（operate_type=3时才允许设置）
	AllowRejoin bool `json:"allow_rejoin,omitempty"`
	// OperatedUsers 被操作成员列表
	OperatedUsers []OperatedUser `json:"operated_users"`
}

// SwitchUserVideoRequest 关闭或开启成员视频请求
type SwitchUserVideoRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// Video 操作类型：false：关闭视频（默认值）；true：开启视频，仅支持MRA设备
	Video bool `json:"video,omitempty"`
	// OperatedUser 被操作者
	OperatedUser OperatedUser `json:"operated_user"`
}

// CloseScreenShareRequest 关闭成员屏幕共享请求
type CloseScreenShareRequest struct {
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// OperatedUser 被操作对象
	OperatedUser OperatedUser `json:"operated_user"`
}
