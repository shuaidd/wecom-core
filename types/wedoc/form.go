package wedoc

// CreateFormRequest 创建收集表请求
type CreateFormRequest struct {
	SpaceID  string   `json:"spaceid,omitempty"`  // 空间spaceid
	FatherID string   `json:"fatherid,omitempty"` // 父目录fileid, 在根目录时为空间spaceid
	FormInfo FormInfo `json:"form_info"`          // 收集表信息
}

// FormInfo 收集表信息
type FormInfo struct {
	FormTitle    string       `json:"form_title,omitempty"`    // 收集表标题
	FormDesc     string       `json:"form_desc,omitempty"`     // 收集表描述
	FormHeader   string       `json:"form_header,omitempty"`   // 收集表表头背景图链接
	FormQuestion FormQuestion `json:"form_question,omitempty"` // 收集表的问题列表
	FormSetting  FormSetting  `json:"form_setting,omitempty"`  // 收集表设置
}

// FormQuestion 收集表问题列表
type FormQuestion struct {
	Items []QuestionItem `json:"items,omitempty"` // 问题数组，不超过200个
}

// QuestionItem 问题项
type QuestionItem struct {
	QuestionID              uint32                  `json:"question_id"`                         // 问题id，从1开始
	Title                   string                  `json:"title"`                               // 问题描述
	Pos                     uint32                  `json:"pos"`                                 // 问题序号，从1开始
	Status                  uint32                  `json:"status"`                              // 问题状态。1：正常；2：被删除
	ReplyType               uint32                  `json:"reply_type"`                          // 问题类型。1：文本；2：单选；3：多选；5：位置；9：图片；10：文件；11：日期；14：时间；15：下拉列表；16：体温；17：签名；18：部门；19：成员 22：时长
	MustReply               bool                    `json:"must_reply"`                          // 是否必答
	Note                    string                  `json:"note,omitempty"`                      // 问题备注
	Placeholder             string                  `json:"placeholder,omitempty"`               // 编辑提示
	QuestionExtendSetting   QuestionExtendSetting   `json:"question_extend_setting,omitempty"`   // 问题的额外设置
	OptionItem              []OptionItem            `json:"option_item,omitempty"`               // 单选/多选/下拉列表题的选项列表
}

// OptionItem 选项
type OptionItem struct {
	Key    uint32 `json:"key"`    // 选项key（1，2，3...）
	Value  string `json:"value"`  // 选项内容
	Status uint32 `json:"status"` // 选项状态。1：正常；2：被删除
}

// QuestionExtendSetting 问题的额外设置
type QuestionExtendSetting struct {
	TextSetting        *TextSetting        `json:"text_setting,omitempty"`        // 文本题设置
	RadioSetting       *RadioSetting       `json:"radio_setting,omitempty"`       // 单选题设置
	CheckboxSetting    *CheckboxSetting    `json:"checkbox_setting,omitempty"`    // 多选题设置
	LocationSetting    *LocationSetting    `json:"location_setting,omitempty"`    // 位置题设置
	ImageSetting       *ImageSetting       `json:"image_setting,omitempty"`       // 图片题设置
	FileSetting        *FileSetting        `json:"file_setting,omitempty"`        // 文件题设置
	DateSetting        *DateSetting        `json:"date_setting,omitempty"`        // 日期题设置
	TimeSetting        *TimeSetting        `json:"time_setting,omitempty"`        // 时间题设置
	DurationSetting    *DurationSetting    `json:"duration_setting,omitempty"`    // 时长题设置
	TemperatureSetting *TemperatureSetting `json:"temperature_setting,omitempty"` // 体温题设置
	DepartmentSetting  *DepartmentSetting  `json:"department_setting,omitempty"`  // 部门题设置
	MemberSetting      *MemberSetting      `json:"member_setting,omitempty"`      // 成员题设置
}

// TextSetting 文本题设置
type TextSetting struct {
	ValidationType uint32  `json:"validation_type,omitempty"`   // 校验类型。0: 字符个数 1: 数字 2: 电子邮箱 3: 网址 4: 身份证 5: 手机号（大陆地区） 6: 固定电话
	ValidationDetail uint32  `json:"validation_detail,omitempty"` // 校验详情
	CharLen        uint32  `json:"char_len,omitempty"`          // 字符长度
	NumberMin      float64 `json:"number_min,omitempty"`        // 数字的区间左端
	NumberMax      float64 `json:"number_max,omitempty"`        // 数字的区间右端
}

// RadioSetting 单选题设置
type RadioSetting struct {
	AddOtherOption bool `json:"add_other_option,omitempty"` // 是否增加"其他"选项
}

// CheckboxSetting 多选题设置
type CheckboxSetting struct {
	AddOtherOption bool   `json:"add_other_option,omitempty"` // 是否增加"其他"选项
	Type           uint32 `json:"type,omitempty"`             // 多选类型。0: 不限制可选数量 1: 至少选择 2: 最多选择 3: 固定选择
	Number         uint32 `json:"number,omitempty"`           // 多选题可勾选的数量的限制
}

// LocationSetting 位置题设置
type LocationSetting struct {
	LocationType uint32 `json:"location_type,omitempty"` // 位置类型。0: 省/市/区/街道+详细地址 1: 省/市 2: 省/市/区 3: 省/市/区/街道 4: 自动定位
	DistanceType uint32 `json:"distance_type,omitempty"` // 允许定位范围。0: 当前位置 1: 附近100米 2: 附近200米 3: 附近300米
}

// ImageSetting 图片题设置
type ImageSetting struct {
	CameraOnly       bool             `json:"camera_only,omitempty"`        // 是否仅限手机拍照
	UploadImageLimit *UploadFileLimit `json:"upload_image_limit,omitempty"` // 数量和大小限制信息
}

// FileSetting 文件题设置
type FileSetting struct {
	UploadFileLimit *UploadFileLimit `json:"upload_file_limit,omitempty"` // 数量和大小限制信息
}

// UploadFileLimit 上传文件/图片限制
type UploadFileLimit struct {
	CountLimitType uint32 `json:"count_limit_type,omitempty"` // 数量限制类型。0: 等于count数量 1: 小于等于count数量
	Count          uint32 `json:"count,omitempty"`            // 限制的数量
	MaxSize        uint64 `json:"max_size,omitempty"`         // 单个文件大小限制MB
}

// DateSetting 日期题设置
type DateSetting struct {
	DateFormatType uint32 `json:"date_format_type,omitempty"` // 日期格式。0: 年/月/日/时/分 1: 年/月/日 2: 年/月
}

// TimeSetting 时间题设置
type TimeSetting struct {
	TimeFormatType uint32 `json:"time_format_type,omitempty"` // 时间格式。0: 时分 1: 时分秒
}

// DurationSetting 时长题设置
type DurationSetting struct {
	TimeScale uint32 `json:"time_scale,omitempty"` // 时间刻度。1: 按天 2: 按小时
	DateType  uint32 `json:"date_type,omitempty"`  // 日期类型。1: 自然日 2: 工作日（跳过双休和法定节假日）
	DayRange  uint32 `json:"day_range,omitempty"`  // 单位换算，多少小时/天
}

// TemperatureSetting 体温题设置
type TemperatureSetting struct {
	UnitType uint32 `json:"unit_type,omitempty"` // 温度单位。0: 摄氏度 1: 华氏度
}

// DepartmentSetting 部门题设置
type DepartmentSetting struct {
	AllowMultipleSelection bool `json:"allow_multiple_selection,omitempty"` // 是否允许多选
}

// MemberSetting 成员题设置
type MemberSetting struct {
	AllowMultipleSelection bool `json:"allow_multiple_selection,omitempty"` // 是否允许多选
}

// FormSetting 收集表设置
type FormSetting struct {
	FillOutAuth          uint32              `json:"fill_out_auth,omitempty"`           // 填写权限。0：所有人；1：企业内指定人/部门；4:家校所有范围
	FillInRange          *FillInRange        `json:"fill_in_range,omitempty"`           // 指定的可填写的人/部门
	SettingManagerRange  *SettingManagerRange `json:"setting_manager_range,omitempty"`   // 收集表管理员
	TimedRepeatInfo      *TimedRepeatInfo    `json:"timed_repeat_info,omitempty"`       // 定时重复设置项
	AllowMultiFill       bool                `json:"allow_multi_fill,omitempty"`        // 是否允许每人提交多份
	TimedFinish          uint32              `json:"timed_finish,omitempty"`            // 定时关闭
	CanAnonymous         bool                `json:"can_anonymous,omitempty"`           // 是否支持匿名填写
	CanNotifySubmit      bool                `json:"can_notify_submit,omitempty"`       // 是否有回复时提醒
}

// FillInRange 可填写的人/部门
type FillInRange struct {
	UserIDs       []string `json:"userids,omitempty"`       // 企业成员userid列表
	DepartmentIDs []uint64 `json:"departmentids,omitempty"` // 部门id列表
}

// SettingManagerRange 收集表管理员
type SettingManagerRange struct {
	UserIDs []string `json:"userids,omitempty"` // 企业成员userid列表
}

// TimedRepeatInfo 定时重复设置项
type TimedRepeatInfo struct {
	Enable         bool   `json:"enable,omitempty"`           // 是否开启定时重复
	RemindTime     uint32 `json:"remind_time,omitempty"`      // 提醒时间
	RepeatType     uint32 `json:"repeat_type,omitempty"`      // 重复类型。0：每周；1：每天；2：每月
	WeekFlag       uint32 `json:"week_flag,omitempty"`        // 每周几重复
	SkipHoliday    bool   `json:"skip_holiday,omitempty"`     // 自动跳过节假日
	DayOfMonth     uint32 `json:"day_of_month,omitempty"`     // 每月的第几天（1 - 31）
	ForkFinishType uint32 `json:"fork_finish_type,omitempty"` // 是否允许补填。0：允许；1：仅当天；2：最后五天内；3：一个月内；4：下一次生成前
	RuleCTime      uint32 `json:"rule_ctime,omitempty"`       // 规则创建时间
	RuleMTime      uint32 `json:"rule_mtime,omitempty"`       // 规则修改时间
}

// CreateFormResponse 创建收集表响应
type CreateFormResponse struct {
	FormID string `json:"formid"` // 收集表id
}

// GetFormInfoRequest 获取收集表信息请求
type GetFormInfoRequest struct {
	FormID string `json:"formid"` // 收集表ID
}

// GetFormInfoResponse 获取收集表信息响应
type GetFormInfoResponse struct {
	FormInfo   FormInfoDetail `json:"form_info"`   // 收集表信息
}

// FormInfoDetail 收集表详细信息
type FormInfoDetail struct {
	FormID       string       `json:"formid"`                  // 收集表id
	FormTitle    string       `json:"form_title"`              // 收集表标题
	FormDesc     string       `json:"form_desc,omitempty"`     // 收集表描述
	FormHeader   string       `json:"form_header,omitempty"`   // 收集表表头背景图链接
	FormQuestion FormQuestion `json:"form_question,omitempty"` // 收集表的问题列表
	FormSetting  FormSetting  `json:"form_setting,omitempty"`  // 收集表的设置
	RepeatedID   []string     `json:"repeated_id,omitempty"`   // 收集表的周期id
}

// ModifyFormRequest 编辑收集表请求
type ModifyFormRequest struct {
	Oper     uint32   `json:"oper"`               // 操作类型。1：全量修改问题；2：全量修改设置
	FormID   string   `json:"formid"`             // 收集表id
	FormInfo FormInfo `json:"form_info,omitempty"` // 收集表信息
}

// GetFormAnswerRequest 读取收集表答案请求
type GetFormAnswerRequest struct {
	RepeatedID string   `json:"repeated_id"` // 收集表周期id
	AnswerIDs  []uint64 `json:"answer_ids"`  // 需要拉取的答案列表，批次大小最大100
}

// GetFormAnswerResponse 读取收集表答案响应
type GetFormAnswerResponse struct {
	Answer Answer `json:"answer"` // 答案
}

// Answer 答案
type Answer struct {
	AnswerList []AnswerItem `json:"answer_list"` // 答案列表
}

// AnswerItem 答案项
type AnswerItem struct {
	AnswerID          uint64      `json:"answer_id"`                    // 答案id
	UserName          string      `json:"user_name,omitempty"`          // 用户名
	UserID            string      `json:"userid,omitempty"`             // 用户id，匿名填写不返回
	CTime             uint64      `json:"ctime"`                        // 创建时间
	MTime             uint64      `json:"mtime"`                        // 修改时间
	Reply             ReplyDetail `json:"reply"`                        // 该用户的答案明细
	AnswerStatus      uint32      `json:"answer_status"`                // 答案状态 1:正常 3:统计者移除此答案或删除
	TmpExternalUserID string      `json:"tmp_external_userid,omitempty"` // 外部用户临时id
}

// ReplyDetail 回答明细
type ReplyDetail struct {
	Items []ReplyItem `json:"items"` // 每个问题的答案
}

// ReplyItem 回答项
type ReplyItem struct {
	QuestionID        uint64                 `json:"question_id"`                   // 问题id
	TextReply         string                 `json:"text_reply,omitempty"`          // 文本答案
	OptionReply       []uint32               `json:"option_reply,omitempty"`        // 选择题答案
	OptionExtendReply []OptionExtendReply    `json:"option_extend_reply,omitempty"` // 选择题，其他选项列表
	FileExtendReply   []FileExtendReply      `json:"file_extend_reply,omitempty"`   // 文件题答案列表
	DepartmentReply   *DepartmentReply       `json:"department_reply,omitempty"`    // 部门题答案
	MemberReply       *MemberReply           `json:"member_reply,omitempty"`        // 成员题答案
	DurationReply     *DurationReply         `json:"duration_reply,omitempty"`      // 时长题答案
}

// OptionExtendReply 其他选项答案
type OptionExtendReply struct {
	OptionReply uint32 `json:"option_reply"` // 其他选项的答案id
	ExtendText  string `json:"extend_text"`  // 其他选项的答案字符串
}

// FileExtendReply 文件答案
type FileExtendReply struct {
	Name   string `json:"name"`   // 文件名
	FileID string `json:"fileid"` // 文件id
}

// DepartmentReply 部门答案
type DepartmentReply struct {
	List []DepartmentItem `json:"list"` // 部门列表
}

// DepartmentItem 部门项
type DepartmentItem struct {
	DepartmentID uint64 `json:"department_id"` // 部门id
}

// MemberReply 成员答案
type MemberReply struct {
	List []MemberItem `json:"list"` // 成员列表
}

// MemberItem 成员项
type MemberItem struct {
	UserID string `json:"userid"` // 成员id
}

// DurationReply 时长答案
type DurationReply struct {
	BeginTime uint32  `json:"begin_time"`        // 开始时间，时间戳
	EndTime   uint32  `json:"end_time"`          // 结束时间，时间戳
	TimeScale uint32  `json:"time_scale"`        // 时间刻度。1: 按天 2: 按小时
	DayRange  uint32  `json:"day_range"`         // 单位换算，多少小时/天
	Days      float64 `json:"days,omitempty"`    // 天数
	Hours     float64 `json:"hours,omitempty"`   // 小时数
}

// GetFormStatisticRequest 收集表的统计信息查询请求
type GetFormStatisticRequest struct {
	RepeatedID string `json:"repeated_id"`       // 收集表的repeated_id
	ReqType    uint32 `json:"req_type"`          // 请求类型 1:只获取统计结果 2:获取已提交列表 3:获取未提交列表
	StartTime  uint64 `json:"start_time,omitempty"` // 筛选开始时间
	EndTime    uint64 `json:"end_time,omitempty"`   // 筛选结束时间
	Limit      uint64 `json:"limit,omitempty"`      // 分页拉取时批次大小，最大10000
	Cursor     uint64 `json:"cursor,omitempty"`     // 分页拉取的游标
}

// GetFormStatisticResponse 收集表的统计信息查询响应
type GetFormStatisticResponse struct {
	FillCnt        uint64       `json:"fill_cnt"`               // 已填写次数
	FillUserCnt    uint64       `json:"fill_user_cnt"`          // 已填写人数
	UnfillUserCnt  uint64       `json:"unfill_user_cnt"`        // 未填写人数
	SubmitUsers    []SubmitUser `json:"submit_users,omitempty"` // 已填写人列表
	UnfillUsers    []UnfillUser `json:"unfill_users,omitempty"` // 未填写人列表
	HasMore        bool         `json:"has_more,omitempty"`     // 是否还有更多
	Cursor         uint64       `json:"cursor,omitempty"`       // 上次分页拉取返回的cursor
}

// SubmitUser 已提交用户
type SubmitUser struct {
	UserID            string `json:"userid,omitempty"`              // 企业内成员的id
	TmpExternalUserID string `json:"tmp_external_userid,omitempty"` // 外部用户临时id
	SubmitTime        uint64 `json:"submit_time"`                   // 提交时间
	AnswerID          uint64 `json:"answer_id"`                     // 答案id
	UserName          string `json:"user_name,omitempty"`           // 名字
}

// UnfillUser 未填写用户
type UnfillUser struct {
	UserID   string `json:"userid"`    // 企业内成员的id
	UserName string `json:"user_name"` // 名字
}
