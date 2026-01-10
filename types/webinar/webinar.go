package webinar

import "github.com/shuaidd/wecom-core/types/common"

type UpdateWebinarRequest struct {
	MeetingID                string        `json:"meetingid"`
	Title                    string        `json:"title"`
	Sponsor                  string        `json:"sponsor,omitempty"`
	StartTime                string        `json:"start_time"`
	EndTime                  string        `json:"end_time"`
	AdmissionType            uint32        `json:"admission_type"`
	Hosts                    []HostInfo    `json:"hosts,omitempty"`
	Password                 string        `json:"password,omitempty"`
	CoverURL                 string        `json:"cover_url,omitempty"`
	Description              string        `json:"description,omitempty"`
	EnableGuestInviteLink    bool          `json:"enable_guest_invite_link,omitempty"`
	MediaSetting             *MediaSetting `json:"media_setting,omitempty"`
	EnableQA                 bool          `json:"enable_qa,omitempty"`
	SensitiveWords           []string      `json:"sensitive_words,omitempty"`
	EnableManualCheck        bool          `json:"enable_manual_check,omitempty"`
	ActivityPage             bool          `json:"activity_page,omitempty"`
	DisplayNumberOfAttendees uint32        `json:"display_number_of_attendees,omitempty"`
	PlaybackForAudience      bool          `json:"playback_for_audience"`
	PreparationMode          bool          `json:"preparation_mode,omitempty"`
}

type HostInfo struct {
	UserID string `json:"userid"`
}

type MediaSetting struct {
	EnableEnterMute           bool   `json:"enable_enter_mute,omitempty"`
	AllowUnmuteSelf           bool   `json:"allow_unmute_self,omitempty"`
	AllowEnterBeforeHost      bool   `json:"allow_enter_before_host,omitempty"`
	EnableScreenWatermark     bool   `json:"enable_screen_watermark,omitempty"`
	WatermarkType             uint32 `json:"watermark_type,omitempty"`
	AllowExternalUser         bool   `json:"allow_external_user,omitempty"`
	AutoRecordType            string `json:"auto_record_type,omitempty"`
	AttendeeJoinAutoRecord    bool   `json:"attendee_join_auto_record,omitempty"`
	EnableHostPauseAutoRecord bool   `json:"enable_host_pause_auto_record,omitempty"`
}

type SetWebinarConfigRequest struct {
	MeetingID                    string     `json:"meetingid"`
	ApproveType                  int32      `json:"approve_type,omitempty"`
	IsCollectQuestion            int32      `json:"is_collect_question,omitempty"`
	NoRegistrationNeededForStaff bool       `json:"no_registration_needed_for_staff,omitempty"`
	QuestionList                 []Question `json:"question_list,omitempty"`
}

type SetWebinarConfigResponse struct {
	common.Response
	QuestionCount int32 `json:"question_count"`
}

type Question struct {
	IsRequired    uint32   `json:"is_required"`
	QuestionTitle string   `json:"question_title,omitempty"`
	OptionList    []Option `json:"option_list,omitempty"`
	QuestionType  uint32   `json:"question_type,omitempty"`
	SpecialType   uint32   `json:"special_type,omitempty"`
}

type Option struct {
	Content string `json:"content,omitempty"`
}

type DeleteWebinarEnrollRequest struct {
	MeetingID    string     `json:"meetingid"`
	EnrollIDList []EnrollID `json:"enroll_id_list"`
}

type EnrollID struct {
	EnrollID string `json:"enroll_id"`
}

type DeleteWebinarEnrollResponse struct {
	common.Response
	TotalCount uint32 `json:"total_count"`
}

type CancelWebinarRequest struct {
	MeetingID string `json:"meetingid"`
}

type ApproveWebinarEnrollRequest struct {
	MeetingID    string   `json:"meetingid"`
	EnrollIDList []string `json:"enroll_id_list"`
	Action       uint32   `json:"action"`
}

type ApproveWebinarEnrollResponse struct {
	common.Response
	HandledCount uint32 `json:"handled_count"`
}

type UpdateWebinarWarmUpRequest struct {
	MeetingID                  string `json:"meetingid"`
	WarmUpPicture              string `json:"warm_up_picture,omitempty"`
	WarmUpVideo                string `json:"warm_up_video,omitempty"`
	AllowAttendeesInviteOthers bool   `json:"allow_attendees_invite_others,omitempty"`
}

type ListWebinarGuestRequest struct {
	MeetingID   string `json:"meetingid,omitempty"`
	MeetingCode string `json:"meeting_code,omitempty"`
}

type ListWebinarGuestResponse struct {
	common.Response
	Guests []GuestInfo `json:"guests"`
}

type GuestInfo struct {
	GuestType   uint32 `json:"guest_type"`
	UserID      string `json:"userid,omitempty"`
	Area        string `json:"area,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	GuestName   string `json:"guest_name,omitempty"`
	Email       string `json:"email,omitempty"`
}

type QueryWebinarEnrollByTmpOpenIDRequest struct {
	MeetingID     string   `json:"meetingid"`
	SortingRules  uint32   `json:"sorting_rules,omitempty"`
	TmpOpenIDList []string `json:"tmp_openid_list"`
}

type QueryWebinarEnrollByTmpOpenIDResponse struct {
	common.Response
	EnrollIDList []EnrollIDInfo `json:"enroll_id_list,omitempty"`
}

type EnrollIDInfo struct {
	TmpOpenID string `json:"tmp_openid"`
	EnrollID  string `json:"enroll_id"`
}

type ListWebinarEnrollRequest struct {
	MeetingID string `json:"meetingid"`
	Status    uint32 `json:"status,omitempty"`
	Cursor    string `json:"cursor,omitempty"`
	Limit     uint32 `json:"limit,omitempty"`
}

type ListWebinarEnrollResponse struct {
	common.Response
	HasMore    bool     `json:"has_more"`
	NextCursor string   `json:"next_cursor,omitempty"`
	EnrollList []Enroll `json:"enroll_list,omitempty"`
}

type Enroll struct {
	EnrollID         string   `json:"enroll_id"`
	EnrollTime       string   `json:"enroll_time"`
	EnrollSourceType uint32   `json:"enroll_source_type"`
	NickName         string   `json:"nick_name"`
	Status           uint32   `json:"status"`
	UserID           string   `json:"userid,omitempty"`
	TmpOpenID        string   `json:"tmp_openid,omitempty"`
	EnrollCode       string   `json:"enroll_code,omitempty"`
	AnswerList       []Answer `json:"answer_list,omitempty"`
}

type Answer struct {
	AnswerContent []string `json:"answer_content"`
	IsRequired    uint32   `json:"is_required"`
	QuestionNum   uint32   `json:"question_num"`
	QuestionTitle string   `json:"question_title"`
	QuestionType  uint32   `json:"question_type"`
	SpecialType   uint32   `json:"special_type"`
}

type GetWebinarConfigRequest struct {
	MeetingID string `json:"meetingid"`
}

type GetWebinarConfigResponse struct {
	common.Response
	ApproveType                  uint32     `json:"approve_type"`
	IsCollectQuestion            uint32     `json:"is_collect_question"`
	NoRegistrationNeededForStaff bool       `json:"no_registration_needed_for_staff"`
	QuestionList                 []Question `json:"question_list,omitempty"`
}

type GetWebinarRequest struct {
	MeetingID   string `json:"meetingid,omitempty"`
	MeetingCode string `json:"meeting_code,omitempty"`
}

type GetWebinarResponse struct {
	common.Response
	MeetingID                  string        `json:"meetingid,omitempty"`
	MeetingCode                string        `json:"meeting_code,omitempty"`
	Title                      string        `json:"title,omitempty"`
	Sponsor                    string        `json:"sponsor,omitempty"`
	StartTime                  string        `json:"start_time,omitempty"`
	EndTime                    string        `json:"end_time,omitempty"`
	AdmissionType              uint32        `json:"admission_type,omitempty"`
	Hosts                      []HostInfo    `json:"hosts,omitempty"`
	Password                   string        `json:"password,omitempty"`
	CoverURL                   string        `json:"cover_url,omitempty"`
	Description                string        `json:"description,omitempty"`
	EnableGuestInviteLink      bool          `json:"enable_guest_invite_link,omitempty"`
	AudienceJoinLink           string        `json:"audience_join_link,omitempty"`
	GuestJoinLink              string        `json:"guest_join_link,omitempty"`
	MediaSetting               *MediaSetting `json:"media_setting,omitempty"`
	EnableQA                   bool          `json:"enable_qa,omitempty"`
	ManualCheckLink            string        `json:"manual_check_link,omitempty"`
	ManualCheckPassword        string        `json:"manual_check_password,omitempty"`
	ActivityPage               bool          `json:"activity_page,omitempty"`
	DisplayNumberOfAttendees   uint32        `json:"display_number_of_attendees,omitempty"`
	PlaybackForAudience        bool          `json:"playback_for_audience,omitempty"`
	PlaybackURL                string        `json:"playback_url,omitempty"`
	PreparationMode            bool          `json:"preparation_mode,omitempty"`
	WarmUpPicture              string        `json:"warm_up_picture,omitempty"`
	WarmUpVideo                string        `json:"warm_up_video,omitempty"`
	AllowAttendeesInviteOthers bool          `json:"allow_attendees_invite_others,omitempty"`
}
