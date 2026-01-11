package calendar

import "github.com/shuaidd/wecom-core/internal/client"

// Attendee 日程参与者
type Attendee struct {
	UserID         string `json:"userid"`                    // 日程参与者ID
	ResponseStatus int    `json:"response_status,omitempty"` // 日程参与者的接受状态。0-未处理；1-待定；2-全部接受；3-仅接受一次；4-拒绝
}

// Reminders 提醒相关信息
type Reminders struct {
	IsRemind              int           `json:"is_remind,omitempty"`                // 是否需要提醒。0-否；1-是
	RemindBeforeEventSecs int           `json:"remind_before_event_secs,omitempty"` // 日程开始前多少秒提醒
	RemindTimeDiffs       []int         `json:"remind_time_diffs,omitempty"`        // 提醒时间与日程开始时间的差值
	IsRepeat              int           `json:"is_repeat,omitempty"`                // 是否重复日程。0-否；1-是
	RepeatType            int           `json:"repeat_type,omitempty"`              // 重复类型。0-每日；1-每周；2-每月；5-每年；7-工作日
	RepeatUntil           int           `json:"repeat_until,omitempty"`             // 重复结束时刻，Unix时间戳
	IsCustomRepeat        int           `json:"is_custom_repeat,omitempty"`         // 是否自定义重复。0-否；1-是
	RepeatInterval        int           `json:"repeat_interval,omitempty"`          // 重复间隔
	RepeatDayOfWeek       []int         `json:"repeat_day_of_week,omitempty"`       // 每周周几重复
	RepeatDayOfMonth      []int         `json:"repeat_day_of_month,omitempty"`      // 每月哪几天重复
	Timezone              int           `json:"timezone,omitempty"`                 // 时区。UTC偏移量表示
	ExcludeTimeList       []ExcludeTime `json:"exclude_time_list,omitempty"`        // 重复日程不包含的日期列表
}

// ExcludeTime 不包含的日期
type ExcludeTime struct {
	StartTime int `json:"start_time"` // 不包含的日期时间戳
}

// Schedule 日程信息
type Schedule struct {
	ScheduleID  string     `json:"schedule_id,omitempty"`  // 日程ID
	Admins      []string   `json:"admins,omitempty"`       // 日程的管理员userid列表
	Attendees   []Attendee `json:"attendees,omitempty"`    // 日程参与者列表
	Summary     string     `json:"summary,omitempty"`      // 日程标题
	Description string     `json:"description,omitempty"`  // 日程描述
	Reminders   *Reminders `json:"reminders,omitempty"`    // 提醒相关信息
	Location    string     `json:"location,omitempty"`     // 日程地址
	CalID       string     `json:"cal_id,omitempty"`       // 日程所属日历ID
	StartTime   int        `json:"start_time"`             // 日程开始时间，Unix时间戳
	EndTime     int        `json:"end_time"`               // 日程结束时间，Unix时间戳
	Status      int        `json:"status,omitempty"`       // 日程状态。0-正常；1-已取消
	IsWholeDay  int        `json:"is_whole_day,omitempty"` // 是否全天日程。0-否；1-是
	Sequence    int64      `json:"sequence,omitempty"`     // 日程编号，是一个自增数字
}

// CreateScheduleRequest 创建日程请求
type CreateScheduleRequest struct {
	Schedule Schedule `json:"schedule"`          // 日程信息
	AgentID  int      `json:"agentid,omitempty"` // 授权方安装的应用agentid
}

// CreateScheduleResponse 创建日程响应
type CreateScheduleResponse struct {
	client.CommonResponse
	ScheduleID string `json:"schedule_id"` // 日程ID
}

// GetScheduleRequest 获取日程详情请求
type GetScheduleRequest struct {
	ScheduleIDList []string `json:"schedule_id_list"` // 日程ID列表，一次最多拉取1000条
}

// GetScheduleResponse 获取日程详情响应
type GetScheduleResponse struct {
	client.CommonResponse
	ScheduleList []Schedule `json:"schedule_list"` // 日程列表
}

// UpdateScheduleRequest 更新日程请求
type UpdateScheduleRequest struct {
	SkipAttendees int      `json:"skip_attendees,omitempty"` // 是否不更新参与人。0-否；1-是
	OpMode        int      `json:"op_mode,omitempty"`        // 操作模式。0-默认全部修改；1-仅修改此日程；2-修改将来的所有日程
	OpStartTime   int      `json:"op_start_time,omitempty"`  // 操作起始时间
	Schedule      Schedule `json:"schedule"`                 // 日程信息
}

// UpdateScheduleResponse 更新日程响应
type UpdateScheduleResponse struct {
	client.CommonResponse
	ScheduleID string `json:"schedule_id,omitempty"` // 修改重复日程新产生的日程ID
}

// DeleteScheduleRequest 取消日程请求
type DeleteScheduleRequest struct {
	ScheduleID  string `json:"schedule_id"`             // 日程ID
	OpMode      int    `json:"op_mode,omitempty"`       // 操作模式。0-默认删除所有日程；1-仅删除此日程；2-删除本次及后续日程
	OpStartTime int    `json:"op_start_time,omitempty"` // 操作起始时间
}

// AddAttendeesRequest 新增日程参与者请求
type AddAttendeesRequest struct {
	ScheduleID string     `json:"schedule_id"` // 日程ID
	Attendees  []Attendee `json:"attendees"`   // 日程参与者列表
}

// DeleteAttendeesRequest 删除日程参与者请求
type DeleteAttendeesRequest struct {
	ScheduleID string     `json:"schedule_id"` // 日程ID
	Attendees  []Attendee `json:"attendees"`   // 日程参与者列表
}

// GetScheduleByCalendarRequest 获取日历下的日程列表请求
type GetScheduleByCalendarRequest struct {
	CalID  string `json:"cal_id"`           // 日历ID
	Offset int    `json:"offset,omitempty"` // 分页，偏移量, 默认为0
	Limit  int    `json:"limit,omitempty"`  // 分页，预期请求的数据量，默认为500，取值范围 1 ~ 1000
}

// GetScheduleByCalendarResponse 获取日历下的日程列表响应
type GetScheduleByCalendarResponse struct {
	client.CommonResponse
	ScheduleList []Schedule `json:"schedule_list"` // 日程列表
}
