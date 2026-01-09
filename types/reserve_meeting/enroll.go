package reserve_meeting

import "github.com/shuaidd/wecom-core/types/common"

// EnrollQueryByTmpOpenIDRequest 获取会议成员报名ID请求
type EnrollQueryByTmpOpenIDRequest struct {
	MeetingID     string   `json:"meetingid"`
	SortingRules  int32    `json:"sorting_rules,omitempty"`
	TmpOpenIDList []string `json:"tmp_openid_list"`
}

// EnrollQueryByTmpOpenIDResponse 获取会议成员报名ID响应
type EnrollQueryByTmpOpenIDResponse struct {
	common.Response
	EnrollIDList []EnrollID `json:"enroll_id_list,omitempty"`
}

// EnrollID 报名ID
type EnrollID struct {
	TmpOpenID string `json:"tmp_openid"`
	EnrollID  string `json:"enroll_id"`
}

// EnrollListRequest 获取会议报名信息请求
type EnrollListRequest struct {
	MeetingID string `json:"meetingid"`
	Status    int32  `json:"status,omitempty"`
	Cursor    string `json:"cursor,omitempty"`
	Limit     int32  `json:"limit,omitempty"`
}

// EnrollListResponse 获取会议报名信息响应
type EnrollListResponse struct {
	common.Response
	HasMore    bool     `json:"has_more"`
	NextCursor string   `json:"next_cursor,omitempty"`
	EnrollList []Enroll `json:"enroll_list,omitempty"`
}

// Enroll 报名信息
type Enroll struct {
	EnrollID         string   `json:"enroll_id"`
	EnrollTime       string   `json:"enroll_time"`
	EnrollSourceType int32    `json:"enroll_source_type"`
	NickName         string   `json:"nick_name"`
	Status           int32    `json:"status"`
	UserID           string   `json:"userid,omitempty"`
	TmpOpenID        string   `json:"tmp_openid"`
	EnrollCode       string   `json:"enroll_code"`
	AnswerList       []Answer `json:"answer_list,omitempty"`
}

// Answer 答题
type Answer struct {
	AnswerContent []string `json:"answer_content"`
	IsRequired    int32    `json:"is_required"`
	QuestionNum   int32    `json:"question_num"`
	QuestionTitle string   `json:"question_title"`
	QuestionType  int32    `json:"question_type"`
	SpecialType   int32    `json:"special_type"`
}

// EnrollGetConfigRequest 获取会议报名配置请求
type EnrollGetConfigRequest struct {
	MeetingID string `json:"meetingid"`
}

// EnrollGetConfigResponse 获取会议报名配置响应
type EnrollGetConfigResponse struct {
	common.Response
	ApproveType                  int32      `json:"approve_type"`
	IsCollectQuestion            int32      `json:"is_collect_question"`
	NoRegistrationNeededForStaff bool       `json:"no_registration_needed_for_staff"`
	QuestionList                 []Question `json:"question_list,omitempty"`
}

// EnrollSetConfigRequest 修改会议报名配置请求
type EnrollSetConfigRequest struct {
	MeetingID                    string     `json:"meetingid"`
	ApproveType                  int32      `json:"approve_type,omitempty"`
	IsCollectQuestion            int32      `json:"is_collect_question,omitempty"`
	NoRegistrationNeededForStaff bool       `json:"no_registration_needed_for_staff,omitempty"`
	QuestionList                 []Question `json:"question_list,omitempty"`
}

// EnrollSetConfigResponse 修改会议报名配置响应
type EnrollSetConfigResponse struct {
	common.Response
	QuestionCount int32 `json:"question_count"`
}

// Question 报名问题
type Question struct {
	IsRequired    int32    `json:"is_required"`
	QuestionTitle string   `json:"question_title,omitempty"`
	OptionList    []Option `json:"option_list,omitempty"`
	QuestionType  int32    `json:"question_type,omitempty"`
	SpecialType   int32    `json:"special_type"`
}

// Option 问题选项
type Option struct {
	Content string `json:"content"`
}

// EnrollDeleteRequest 删除会议报名信息请求
type EnrollDeleteRequest struct {
	MeetingID    string         `json:"meetingid"`
	EnrollIDList []EnrollIDItem `json:"enroll_id_list"`
}

// EnrollIDItem 报名ID项
type EnrollIDItem struct {
	EnrollID string `json:"enroll_id"`
}

// EnrollDeleteResponse 删除会议报名信息响应
type EnrollDeleteResponse struct {
	common.Response
	TotalCount int32 `json:"total_count"`
}

// EnrollImportRequest 导入会议报名信息请求
type EnrollImportRequest struct {
	MeetingID  string             `json:"meetingid"`
	EnrollList []EnrollImportItem `json:"enroll_list"`
}

// EnrollImportItem 导入报名项
type EnrollImportItem struct {
	UserID      string `json:"userid,omitempty"`
	Area        string `json:"area,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	NickName    string `json:"nick_name,omitempty"`
}

// EnrollImportResponse 导入会议报名信息响应
type EnrollImportResponse struct {
	common.Response
	TotalCount int32                      `json:"total_count"`
	EnrollList []EnrollImportResponseItem `json:"enroll_list,omitempty"`
}

// EnrollImportResponseItem 导入报名响应项
type EnrollImportResponseItem struct {
	EnrollID    string `json:"enroll_id"`
	UserID      string `json:"userid"`
	Area        string `json:"area"`
	PhoneNumber string `json:"phone_number"`
	NickName    string `json:"nick_name"`
	EnrollCode  string `json:"enroll_code"`
}

// EnrollApproveRequest 审批会议报名信息请求
type EnrollApproveRequest struct {
	MeetingID    string   `json:"meetingid"`
	Action       int32    `json:"action"`
	EnrollIDList []string `json:"enroll_id_list"`
}

// EnrollApproveResponse 审批会议报名信息响应
type EnrollApproveResponse struct {
	common.Response
	HandledCount int32 `json:"handled_count"`
}
