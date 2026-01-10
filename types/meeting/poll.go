package meeting

import "github.com/shuaidd/wecom-core/types/common"

// PollQuestion 投票问题
type PollQuestion struct {
	// QuestionDesc 问题描述，最多50个字符
	QuestionDesc string `json:"question_desc"`
	// QuestionType 问题选择类型：0：单选；1：多选
	QuestionType int32 `json:"question_type"`
	// PollOption 每个问题支持添加10个选项，最少为1个选项（创建主题最少2个选项）。每个选项最多支持36个字符
	PollOption []string `json:"poll_option"`
	// QuestionID 问题ID（仅响应）
	QuestionID string `json:"question_id,omitempty"`
}

// PollOptionInfo 投票选项信息（响应）
type PollOptionInfo struct {
	// OptionID 选项ID
	OptionID int32 `json:"option_id"`
	// OptionDesc 选项描述
	OptionDesc string `json:"option_desc"`
	// OptionNum 每个选项的投票数
	OptionNum int32 `json:"option_num"`
	// Rate 投票比率
	Rate int32 `json:"rate"`
	// OptionUser 用户信息数组
	OptionUser []PollOptionUser `json:"option_user,omitempty"`
}

// PollOptionUser 投票选项用户信息
type PollOptionUser struct {
	// UserID 成员userid。非匿名投票并且是同企业的用户才返回该字段
	UserID string `json:"userid,omitempty"`
	// TmpOpenID 会议临时ID。非匿名投票才返回该字段
	TmpOpenID string `json:"tmp_openid,omitempty"`
}

// PollQuestionData 投票问题数据（响应）
type PollQuestionData struct {
	// QuestionDesc 问题描述
	QuestionDesc string `json:"question_desc"`
	// QuestionType 问题类型：0：单选；1：多选
	QuestionType int32 `json:"question_type"`
	// QuestionID 问题ID
	QuestionID string `json:"question_id"`
	// OptionInfo 选项信息数组
	OptionInfo []PollOptionInfo `json:"option_info"`
}

// CreatePollThemeRequest 创建会议投票主题请求
type CreatePollThemeRequest struct {
	// OperatorUserID 操作者的openid
	OperatorUserID string `json:"operator_userid"`
	// InstanceID 操作者入会所用的设备id
	InstanceID int32 `json:"instance_id"`
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// PollTopic 投票主题，最多50个字符
	PollTopic string `json:"poll_topic"`
	// PollDesc 投票主题描述，最多100个字符
	PollDesc string `json:"poll_desc"`
	// IsAnony 是否匿名：0：实名，默认值；1：匿名
	IsAnony int32 `json:"is_anony,omitempty"`
	// PollQuestions 投票问题数组，每个投票支持添加10个问题
	PollQuestions []PollQuestion `json:"poll_questions"`
}

// CreatePollThemeResponse 创建会议投票主题响应
type CreatePollThemeResponse struct {
	common.Response
	// PollThemeID 投票主题ID
	PollThemeID string `json:"poll_theme_id"`
}

// UpdatePollThemeRequest 修改会议投票主题请求
type UpdatePollThemeRequest struct {
	// OperatorUserID 操作者openid
	OperatorUserID string `json:"operator_userid"`
	// InstanceID 操作者入会设备对应的id
	InstanceID int32 `json:"instance_id"`
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// PollThemeID 投票主题id
	PollThemeID string `json:"poll_theme_id"`
	// PollTopic 投票主题，最多50个字符
	PollTopic string `json:"poll_topic"`
	// PollDesc 投票主题描述，最多100个字符
	PollDesc string `json:"poll_desc"`
	// IsAnony 是否匿名：0：实名；1：匿名
	IsAnony int32 `json:"is_anony,omitempty"`
	// PollQuestions 投票问题数组，每个投票支持添加10个问题
	PollQuestions []PollQuestion `json:"poll_questions"`
}

// DeletePollRequest 删除会议投票请求
type DeletePollRequest struct {
	// OperatorUserID 操作者的openid
	OperatorUserID string `json:"operator_userid"`
	// InstanceID 操作者入会的设备id
	InstanceID int32 `json:"instance_id"`
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// PollThemeID 投票主题ID，传入则代表删除投票主题，删除投票主题不影响投票实例
	// 投票主题ID和投票ID二选一，如果都传入，会使用投票ID
	PollThemeID string `json:"poll_theme_id,omitempty"`
	// PollID 投票ID，传入则代表删除投票实例。当主题下所有主题实例被删，则投票主题也被删除
	// 投票主题ID和投票ID二选一，如果都传入，会使用投票ID
	PollID string `json:"poll_id,omitempty"`
}

// FinishPollRequest 结束会议投票请求
type FinishPollRequest struct {
	// OperatorUserID 操作者openid
	OperatorUserID string `json:"operator_userid"`
	// InstanceID 操作者入会设备id
	InstanceID int32 `json:"instance_id"`
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// PollThemeID 投票主题ID
	PollThemeID string `json:"poll_theme_id"`
	// PollID 投票ID
	PollID string `json:"poll_id"`
}

// GetPollThemeInfoRequest 获取会议投票主题信息请求
type GetPollThemeInfoRequest struct {
	// OperatorUserID 操作者openid
	OperatorUserID string `json:"operator_userid"`
	// InstanceID 操作者入会设备对应的id
	InstanceID int32 `json:"instance_id"`
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// PollThemeID 投票主题id
	PollThemeID string `json:"poll_theme_id"`
}

// GetPollThemeInfoResponse 获取会议投票主题信息响应
type GetPollThemeInfoResponse struct {
	common.Response
	// PollTopic 投票主题
	PollTopic string `json:"poll_topic"`
	// PollDesc 投票描述
	PollDesc string `json:"poll_desc"`
	// IsAnony 是否匿名：0：实名；1：匿名
	IsAnony int32 `json:"is_anony"`
	// PollQuestionData 投票问题数组
	PollQuestionData []PollQuestionData `json:"poll_question_data"`
}

// GetPollDetailRequest 获取会议投票详情请求
type GetPollDetailRequest struct {
	// OperatorUserID 操作者openid
	OperatorUserID string `json:"operator_userid"`
	// InstanceID 操作者入会设备对应的id
	InstanceID int32 `json:"instance_id"`
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// PollID 投票ID
	PollID string `json:"poll_id"`
}

// GetPollDetailResponse 获取会议投票详情响应
type GetPollDetailResponse struct {
	common.Response
	// PollThemeID 投票主题id
	PollThemeID string `json:"poll_theme_id"`
	// PollTopic 投票主题
	PollTopic string `json:"poll_topic"`
	// PollDesc 投票描述
	PollDesc string `json:"poll_desc"`
	// IsAnony 是否匿名：0：实名；1：匿名
	IsAnony int32 `json:"is_anony"`
	// Status 投票状态：1：投票中；2：已结束
	Status int32 `json:"status"`
	// IsShared 是否共享：0：未共享；1：已共享
	IsShared int32 `json:"is_shared"`
	// VoteTotalNum 投票人数
	VoteTotalNum int32 `json:"vote_total_num"`
	// PollQuestionData 投票结果数组
	PollQuestionData []PollQuestionData `json:"poll_question_data"`
}

// PollInfo 投票信息
type PollInfo struct {
	// PollID 投票ID
	PollID string `json:"poll_id"`
	// PollTopic 投票标题
	PollTopic string `json:"poll_topic"`
	// Status 投票状态：0：未开始；1：投票中；2：已结束
	Status int32 `json:"status"`
	// IsShared 共享状态：0：未共享；1：已共享
	IsShared int32 `json:"is_shared"`
	// IsAnony 是否匿名：0：实名；1：匿名
	IsAnony int32 `json:"is_anony"`
}

// PollThemeInfo 投票主题信息
type PollThemeInfo struct {
	// PollThemeID 投票主题ID
	PollThemeID string `json:"poll_theme_id"`
	// PollsInfo 投票信息列表
	PollsInfo []PollInfo `json:"polls_info"`
}

// GetPollListRequest 获取会议投票列表请求
type GetPollListRequest struct {
	// OperatorUserID 操作者openid
	OperatorUserID string `json:"operator_userid"`
	// InstanceID 操作者入会设备对应的id
	InstanceID int32 `json:"instance_id"`
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
}

// GetPollListResponse 获取会议投票列表响应
type GetPollListResponse struct {
	common.Response
	// PollsThemeInfo 投票主题信息列表
	PollsThemeInfo []PollThemeInfo `json:"polls_theme_info"`
}

// StartPollRequest 发起会议投票请求
type StartPollRequest struct {
	// OperatorUserID 操作者openid
	OperatorUserID string `json:"operator_userid"`
	// InstanceID 操作者入会的设备id
	InstanceID int32 `json:"instance_id"`
	// MeetingID 会议ID
	MeetingID string `json:"meetingid"`
	// PollThemeID 投票主题ID
	PollThemeID string `json:"poll_theme_id"`
}

// StartPollResponse 发起会议投票响应
type StartPollResponse struct {
	common.Response
	// PollID 投票ID
	PollID string `json:"poll_id"`
}
