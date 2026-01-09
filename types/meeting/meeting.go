package meeting

import "github.com/shuaidd/wecom-core/types/common"

// CreateMeetingRequest 创建预约会议请求
type CreateMeetingRequest struct {
	// AdminUserID 会议管理员userid
	AdminUserID string `json:"admin_userid"`
	// Title 会议的标题，最多支持40个字节或者20个utf8字符
	Title string `json:"title"`
	// MeetingStart 会议开始时间的unix时间戳。需大于当前时间
	MeetingStart int64 `json:"meeting_start"`
	// MeetingDuration 会议持续时间（单位秒），最小300秒，最大86399秒
	MeetingDuration int `json:"meeting_duration"`
	// Description 会议的描述，最多支持500个字节或者utf8字符
	Description string `json:"description,omitempty"`
	// Location 会议地点,最多128个字符
	Location string `json:"location,omitempty"`
	// AgentID 授权方安装的应用agentid。仅旧的第三方多应用套件需要填此参数
	AgentID int64 `json:"agentid,omitempty"`
	// Invitees 邀请参会的成员
	Invitees *Invitees `json:"invitees,omitempty"`
	// Settings 会议配置
	Settings *Settings `json:"settings,omitempty"`
	// CalID 会议所属日历ID
	CalID string `json:"cal_id,omitempty"`
	// Reminders 重复会议相关配置
	Reminders *Reminders `json:"reminders,omitempty"`
}

// CreateMeetingResponse 创建预约会议响应
type CreateMeetingResponse struct {
	common.Response
	// MeetingID 会议id
	MeetingID string `json:"meetingid"`
	// ExcessUsers 参会人中包含无效会议账号的userid
	ExcessUsers []string `json:"excess_users,omitempty"`
}

// UpdateMeetingRequest 修改预约会议请求
type UpdateMeetingRequest struct {
	// MeetingID 会议id，仅允许修改预约状态下的会议
	MeetingID string `json:"meetingid"`
	// Title 会议的标题，最多支持40个字节或者20个utf8字符
	Title string `json:"title,omitempty"`
	// MeetingStart 会议开始时间的unix时间戳。需大于当前时间
	MeetingStart int64 `json:"meeting_start,omitempty"`
	// MeetingDuration 会议持续时间（单位秒），最小300秒，最大86399秒
	MeetingDuration int `json:"meeting_duration,omitempty"`
	// Description 会议的描述，最多支持500个字节或者utf8字符
	Description string `json:"description,omitempty"`
	// Location 会议地点,最多128个字符
	Location string `json:"location,omitempty"`
	// AgentID 授权方安装的应用agentid。仅旧的第三方多应用套件需要填此参数
	AgentID int64 `json:"agentid,omitempty"`
	// Invitees 邀请参会的成员
	Invitees *Invitees `json:"invitees,omitempty"`
	// Settings 会议配置
	Settings *Settings `json:"settings,omitempty"`
	// CalID 会议所属日历ID
	CalID string `json:"cal_id,omitempty"`
	// Reminders 重复会议相关配置
	Reminders *Reminders `json:"reminders,omitempty"`
}

// UpdateMeetingResponse 修改预约会议响应
type UpdateMeetingResponse struct {
	common.Response
	// ExcessUsers 参会人中包含无效会议账号的userid
	ExcessUsers []string `json:"excess_users,omitempty"`
}

// CancelMeetingRequest 取消预约会议请求
type CancelMeetingRequest struct {
	// MeetingID 会议id，仅允许取消预约状态下的会议
	MeetingID string `json:"meetingid"`
}

// GetMeetingInfoRequest 获取会议详情请求
type GetMeetingInfoRequest struct {
	// MeetingID 会议id
	MeetingID string `json:"meetingid"`
}

// GetMeetingInfoResponse 获取会议详情响应
type GetMeetingInfoResponse struct {
	common.Response
	// AdminUserID 会议管理员的userId
	AdminUserID string `json:"admin_userid"`
	// Title 会议的标题，最大60个字节
	Title string `json:"title"`
	// MeetingStart 会议开始时间的unix时间戳
	MeetingStart int64 `json:"meeting_start"`
	// MeetingDuration 会议时长
	MeetingDuration int `json:"meeting_duration"`
	// Description 会议的描述，最大600字节
	Description string `json:"description"`
	// Location 会议地点,最多128个字符
	Location string `json:"location"`
	// MainDepartment 发起人所在部门
	MainDepartment int `json:"main_department"`
	// Status 会议的状态，1：待开始，2：会议中，3：已结束，4：已取消，5：已过期
	Status int `json:"status"`
	// AgentID 应用agentid
	AgentID int64 `json:"agentid"`
	// Attendees 会议成员
	Attendees *Attendees `json:"attendees"`
	// Settings 会议配置
	Settings *Settings `json:"settings"`
	// CalID 会议所属日历ID
	CalID string `json:"cal_id"`
	// Reminders 重复会议相关配置
	Reminders *Reminders `json:"reminders"`
	// MeetingCode 会议号
	MeetingCode string `json:"meeting_code"`
	// MeetingLink 入会链接
	MeetingLink string `json:"meeting_link"`
}

// GetUserMeetingIDsRequest 获取成员会议ID列表请求
type GetUserMeetingIDsRequest struct {
	// UserID 企业成员的userid
	UserID string `json:"userid"`
	// Cursor 上一次调用时返回的cursor，初次调用可以填"0"
	Cursor string `json:"cursor,omitempty"`
	// Limit 每次拉取的数据量，默认值和最大值都为100
	Limit int `json:"limit,omitempty"`
	// BeginTime 开始时间
	BeginTime int64 `json:"begin_time,omitempty"`
	// EndTime 结束时间，时间跨度不超过180天
	EndTime int64 `json:"end_time,omitempty"`
}

// GetUserMeetingIDsResponse 获取成员会议ID列表响应
type GetUserMeetingIDsResponse struct {
	common.Response
	// NextCursor 当前数据最后一个key值
	NextCursor string `json:"next_cursor,omitempty"`
	// MeetingIDList 会议ID列表
	MeetingIDList []string `json:"meetingid_list,omitempty"`
}

// Invitees 邀请参会的成员
type Invitees struct {
	// UserID 参与会议的企业成员userid
	UserID []string `json:"userid,omitempty"`
}

// Attendees 会议成员
type Attendees struct {
	// Member 企业内部成员
	Member []AttendeeMember `json:"member,omitempty"`
	// TmpExternalUser 会中参会的外部联系人
	TmpExternalUser []AttendeeTmpExternalUser `json:"tmp_external_user,omitempty"`
}

// AttendeeMember 企业内部成员
type AttendeeMember struct {
	// UserID 企业内部成员的userid
	UserID string `json:"userid"`
	// Status 与会状态。1为已参与，2为未参与
	Status int `json:"status"`
	// FirstJoinTime 参会人首次加入会议时间的unix时间戳
	FirstJoinTime int64 `json:"first_join_time"`
	// LastQuitTime 参会人最后一次离开会议时间的unix时间戳
	LastQuitTime int64 `json:"last_quit_time"`
	// TotalJoinCount 参会人入会次数
	TotalJoinCount int `json:"total_join_count"`
	// CumulativeTime 参会人累计参会时长，单位为秒
	CumulativeTime int `json:"cumulative_time"`
}

// AttendeeTmpExternalUser 会中参会的外部联系人
type AttendeeTmpExternalUser struct {
	// TmpExternalUserID 会中入会的外部用户临时id
	TmpExternalUserID string `json:"tmp_external_userid"`
	// Status 与会状态。1为已参与，2为未参与
	Status int `json:"status"`
	// FirstJoinTime 参会人首次进入会议时间的unix时间戳
	FirstJoinTime int64 `json:"first_join_time"`
	// LastQuitTime 参会人最后一次离开会议时间的unix时间戳
	LastQuitTime int64 `json:"last_quit_time"`
	// TotalJoinCount 参会人入会次数
	TotalJoinCount int `json:"total_join_count"`
	// CumulativeTime 参会人累计参会时长，单位为秒
	CumulativeTime int `json:"cumulative_time"`
}

// Settings 会议配置
type Settings struct {
	// RemindScope 会议开始时来电提醒方式，1.不提醒 2.仅提醒主持人 3.提醒所有成员入 4.指定部分人响铃
	RemindScope int `json:"remind_scope,omitempty"`
	// Password 入会密码，仅可输入4-6位纯数字
	Password string `json:"password,omitempty"`
	// EnableWaitingRoom 是否开启等候室
	EnableWaitingRoom bool `json:"enable_waiting_room,omitempty"`
	// AllowEnterBeforeHost 是否允许成员在主持人进会前加入
	AllowEnterBeforeHost bool `json:"allow_enter_before_host,omitempty"`
	// EnableEnterMute 成员入会时静音；1:开启；0:关闭；2:超过6人后自动开启静音
	EnableEnterMute int `json:"enable_enter_mute,omitempty"`
	// EnableScreenWatermark 是否开启屏幕水印
	EnableScreenWatermark bool `json:"enable_screen_watermark,omitempty"`
	// Hosts 会议主持人人列表，主持人员最多10个
	Hosts *Hosts `json:"hosts,omitempty"`
	// RingUsers 指定响铃的成员列表
	RingUsers *RingUsers `json:"ring_users,omitempty"`
}

// Reminders 重复会议相关配置
type Reminders struct {
	// IsRepeat 是否是周期性会议，1：周期性会议 0：非周期性会议
	IsRepeat int `json:"is_repeat,omitempty"`
	// RepeatType 周期性会议重复类型，0.每天；1.每周；2.每月；7.每个工作日
	RepeatType int `json:"repeat_type,omitempty"`
	// RepeatUntil 重复结束时刻
	RepeatUntil int64 `json:"repeat_until,omitempty"`
	// RepeatInterval 重复间隔
	RepeatInterval int `json:"repeat_interval,omitempty"`
	// RemindBefore 指定会议开始前多久提醒成员
	RemindBefore []int `json:"remind_before,omitempty"`
}

// Hosts 会议主持人人列表
type Hosts struct {
	// UserID 企业成员userid
	UserID []string `json:"userid,omitempty"`
}

// RingUsers 指定响铃的成员列表
type RingUsers struct {
	// UserID 指定响铃的成员userid
	UserID []string `json:"userid,omitempty"`
}
