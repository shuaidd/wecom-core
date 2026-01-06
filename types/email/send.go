package email

// EmailRecipient 邮件收件人
type EmailRecipient struct {
	Emails  []string `json:"emails,omitempty"`  // 邮箱地址列表
	UserIDs []string `json:"userids,omitempty"` // 企业内成员的userid列表
}

// Attachment 邮件附件
type Attachment struct {
	FileName string `json:"file_name"` // 文件名
	Content  string `json:"content"`   // 文件内容(base64编码)
}

// ScheduleReminders 日程/会议提醒和重复设置
type ScheduleReminders struct {
	IsRemind              int   `json:"is_remind,omitempty"`                // 是否有提醒 0-不提醒 1-提醒
	RemindBeforeEventMins int   `json:"remind_before_event_mins,omitempty"` // 开始前多少分钟提醒
	Timezone              int   `json:"timezone,omitempty"`                 // 时区(UTC偏移量)
	IsRepeat              int   `json:"is_repeat,omitempty"`                // 是否重复 0-否 1-是
	IsCustomRepeat        int   `json:"is_custom_repeat,omitempty"`         // 是否自定义重复 0-否 1-是
	RepeatType            int   `json:"repeat_type,omitempty"`              // 重复类型 0-每日 1-每周 2-每月 5-每年
	RepeatInterval        int   `json:"repeat_interval,omitempty"`          // 重复间隔
	RepeatDayOfWeek       []int `json:"repeat_day_of_week,omitempty"`       // 每周周几重复(1-7)
	RepeatDayOfMonth      []int `json:"repeat_day_of_month,omitempty"`      // 每月哪几天重复(1-31)
	RepeatWeekOfMonth     []int `json:"repeat_week_of_month,omitempty"`     // 每月第几周重复
	RepeatMonthOfYear     []int `json:"repeat_month_of_year,omitempty"`     // 每年哪几个月重复(1-12)
	RepeatUntil           int64 `json:"repeat_until,omitempty"`             // 重复结束时刻(Unix时间戳)
}

// Schedule 日程/会议信息
type Schedule struct {
	ScheduleID     string             `json:"schedule_id,omitempty"`     // 日程/会议ID(修改/取消时必填)
	Method         string             `json:"method,omitempty"`          // 方法: request-请求 cancel-取消
	Location       string             `json:"location,omitempty"`        // 地点
	StartTime      int64              `json:"start_time"`                // 开始时间(Unix时间戳)
	EndTime        int64              `json:"end_time"`                  // 结束时间(Unix时间戳)
	Reminders      *ScheduleReminders `json:"reminders,omitempty"`       // 提醒和重复设置
	ScheduleAdmins *EmailRecipient    `json:"schedule_admins,omitempty"` // 日程管理员(仅日程邮件)
}

// MeetingOption 会议选项
type MeetingOption struct {
	Password              string `json:"password,omitempty"`                 // 入会密码(4-6位纯数字)
	AutoRecord            int    `json:"auto_record,omitempty"`              // 是否自动录制 0-未开启 1-本地录制 2-云录制
	EnableWaitingRoom     bool   `json:"enable_waiting_room,omitempty"`      // 是否开启等候室
	AllowEnterBeforeHost  bool   `json:"allow_enter_before_host,omitempty"`  // 是否允许成员在主持人进会前加入
	EnterRestraint        int    `json:"enter_restraint,omitempty"`          // 是否限制成员入会 0-所有人 2-仅企业内部
	EnableScreenWatermark bool   `json:"enable_screen_watermark,omitempty"`  // 是否开启屏幕水印
	EnableEnterMute       int    `json:"enable_enter_mute,omitempty"`        // 成员入会时是否静音 0-关闭 1-开启 2-超过6人自动开启
	RemindScope           int    `json:"remind_scope,omitempty"`             // 会议开始时是否提醒 1-不提醒 2-仅主持人 3-所有成员
	WaterMarkType         int    `json:"water_mark_type,omitempty"`          // 水印类型 0-单排 1-多排
}

// Meeting 会议信息
type Meeting struct {
	Hosts         *EmailRecipient `json:"hosts,omitempty"`          // 会议主持人列表(最多10个)
	MeetingAdmins *EmailRecipient `json:"meeting_admins,omitempty"` // 会议管理员(仅可指定1人)
	Option        *MeetingOption  `json:"option,omitempty"`         // 会议选项
}

// SendEmailRequest 发送邮件请求
type SendEmailRequest struct {
	To             *EmailRecipient `json:"to"`                        // 收件人(必填)
	CC             *EmailRecipient `json:"cc,omitempty"`              // 抄送
	BCC            *EmailRecipient `json:"bcc,omitempty"`             // 密送
	Subject        string          `json:"subject"`                   // 标题(必填)
	Content        string          `json:"content"`                   // 内容(必填)
	AttachmentList []*Attachment   `json:"attachment_list,omitempty"` // 附件列表
	ContentType    string          `json:"content_type,omitempty"`    // 内容类型 html/text(默认html)
	Schedule       *Schedule       `json:"schedule,omitempty"`        // 日程/会议信息
	Meeting        *Meeting        `json:"meeting,omitempty"`         // 会议信息(发送会议邮件时必填)
	EnableIDTrans  uint32          `json:"enable_id_trans,omitempty"` // 是否开启id转译 0-否 1-是
}

// SendEmailResponse 发送邮件响应
type SendEmailResponse struct {
	ErrCode int32  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
