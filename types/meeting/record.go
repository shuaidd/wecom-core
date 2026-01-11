package meeting

import "github.com/shuaidd/wecom-core/types/common"

// UpdateSharingConfigRequest 修改会议录制共享设置请求
type UpdateSharingConfigRequest struct {
	MeetingRecordID string         `json:"meeting_record_id"`
	MeetingID       string         `json:"meetingid"`
	SharingConfig   *SharingConfig `json:"sharing_config,omitempty"`
}

type SharingConfig struct {
	EnableSharing       *bool  `json:"enable_sharing,omitempty"`
	SharingAuthType     *int   `json:"sharing_auth_type,omitempty"`
	EnablePassword      *bool  `json:"enable_password,omitempty"`
	Password            string `json:"password,omitempty"`
	EnableSharingExpire *bool  `json:"enable_sharing_expire,omitempty"`
	SharingExpire       *int64 `json:"sharing_expire,omitempty"`
	AllowDownload       *bool  `json:"allow_download,omitempty"`
}

// DeleteMeetingRecordRequest 删除会议录制（删除会议下所有录制文件）
type DeleteMeetingRecordRequest struct {
	MeetingRecordID string `json:"meeting_record_id"`
	MeetingID       string `json:"meetingid"`
}

// DeleteRecordFileRequest 删除单个录制文件
type DeleteRecordFileRequest struct {
	RecordFileID string `json:"record_file_id"`
	MeetingID    string `json:"meetingid"`
}

// GetRecordListRequest 获取会议录制列表
type GetRecordListRequest struct {
	MeetingID   string `json:"meetingid,omitempty"`
	MeetingCode string `json:"meeting_code,omitempty"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	UserID      string `json:"userid,omitempty"`
	Cursor      string `json:"cursor,omitempty"`
	Limit       int    `json:"limit,omitempty"`
}

type GetRecordListResponse struct {
	common.Response
	HasMore        bool            `json:"has_more,omitempty"`
	NextCursor     string          `json:"next_cursor,omitempty"`
	RecordMeetings []RecordMeeting `json:"record_meetings,omitempty"`
}

type RecordMeeting struct {
	MeetingRecordID  string       `json:"meeting_record_id,omitempty"`
	MeetingID        string       `json:"meetingid,omitempty"`
	MeetingCode      string       `json:"meeting_code,omitempty"`
	HostUserID       string       `json:"host_user_id,omitempty"`
	MeetingStartTime int64        `json:"meeting_start_time,omitempty"`
	Title            string       `json:"title,omitempty"`
	State            int          `json:"state,omitempty"`
	RecordFiles      []RecordFile `json:"record_files,omitempty"`
}

type RecordFile struct {
	RecordFileID     string `json:"record_file_id,omitempty"`
	RecordStartTime  int64  `json:"record_start_time,omitempty"`
	RecordEndTime    int64  `json:"record_end_time,omitempty"`
	RecordSize       int64  `json:"record_size,omitempty"`
	SharingState     int    `json:"sharing_state,omitempty"`
	SharingURL       string `json:"sharing_url,omitempty"`
	RequiredSameCorp bool   `json:"required_same_corp,omitempty"`
	RequiredAttendee bool   `json:"required_attendee,omitempty"`
	Password         string `json:"password,omitempty"`
	SharingExpire    int64  `json:"sharing_expire,omitempty"`
	AllowDownload    bool   `json:"allow_download,omitempty"`
}

// GetFileListRequest 获取会议录制地址（文件列表）
type GetFileListRequest struct {
	MeetingRecordID string `json:"meeting_record_id"`
	MeetingID       string `json:"meetingid"`
}

type GetFileListResponse struct {
	common.Response
	MeetingRecordID string     `json:"meeting_record_id,omitempty"`
	MeetingID       string     `json:"meetingid,omitempty"`
	MeetingCode     string     `json:"meeting_code,omitempty"`
	Title           string     `json:"title,omitempty"`
	RecordFiles     []FileItem `json:"record_files,omitempty"`
}

type FileItem struct {
	RecordFileID            string        `json:"record_file_id,omitempty"`
	ViewAddress             string        `json:"view_address,omitempty"`
	DownloadAddress         string        `json:"download_address,omitempty"`
	DownloadAddressFileType string        `json:"download_address_file_type,omitempty"`
	AudioAddress            string        `json:"audio_address,omitempty"`
	AudioAddressFileType    string        `json:"audio_address_file_type,omitempty"`
	MeetingSummary          []SummaryItem `json:"meeting_summary,omitempty"`
}

type SummaryItem struct {
	DownloadAddress string `json:"download_address,omitempty"`
	FileType        string `json:"file_type,omitempty"`
}

// GetRecordFileRequest 获取单个录制文件详情
type GetRecordFileRequest struct {
	RecordFileID string `json:"record_file_id"`
	MeetingID    string `json:"meetingid"`
}

type GetRecordFileResponse struct {
	common.Response
	RecordFileID            string        `json:"record_file_id,omitempty"`
	MeetingID               string        `json:"meetingid,omitempty"`
	MeetingCode             string        `json:"meeting_code,omitempty"`
	ViewAddress             string        `json:"view_address,omitempty"`
	DownloadAddress         string        `json:"download_address,omitempty"`
	DownloadAddressFileType string        `json:"download_address_file_type,omitempty"`
	AudioAddress            string        `json:"audio_address,omitempty"`
	AudioAddressFileType    string        `json:"audio_address_file_type,omitempty"`
	MeetingSummary          *SummaryItem  `json:"meeting_summary,omitempty"`
	AIMeetingTranscripts    []SummaryItem `json:"ai_meeting_transcripts,omitempty"`
	RecordName              string        `json:"record_name,omitempty"`
	StartTime               string        `json:"start_time,omitempty"`
	EndTime                 string        `json:"end_time,omitempty"`
	MeetingRecordName       string        `json:"meeting_record_name,omitempty"`
}

// GetStatisticsRequest 获取录制文件访问统计
type GetStatisticsRequest struct {
	MeetingRecordID string `json:"meeting_record_id"`
	MeetingID       string `json:"meetingid"`
	StartTime       int64  `json:"start_time,omitempty"`
	EndTime         int64  `json:"end_time,omitempty"`
}

type GetStatisticsResponse struct {
	common.Response
	Summaries []StatisticsSummary `json:"summaries,omitempty"`
}

type StatisticsSummary struct {
	Date          string `json:"date,omitempty"`
	ViewCount     int    `json:"view_count,omitempty"`
	DownloadCount int    `json:"download_count,omitempty"`
}

// TranscriptSearchRequest 获取录制转写搜索结果
type TranscriptSearchRequest struct {
	RecordFileID string `json:"record_file_id"`
	MeetingID    string `json:"meetingid"`
	Text         string `json:"text"`
}

type TranscriptSearchResponse struct {
	common.Response
	Hits      []TranscriptHit  `json:"hits,omitempty"`
	Timelines []TranscriptTime `json:"timelines,omitempty"`
}

type TranscriptHit struct {
	PID    string `json:"pid,omitempty"`
	SID    string `json:"sid,omitempty"`
	Offset int    `json:"offset,omitempty"`
	Length int    `json:"length,omitempty"`
}

type TranscriptTime struct {
	PID       string `json:"pid,omitempty"`
	SID       string `json:"sid,omitempty"`
	StartTime int64  `json:"start_time,omitempty"`
}

// GetTranscriptParagraphListRequest 获取录制转写段落信息
type GetTranscriptParagraphListRequest struct {
	RecordFileID string `json:"record_file_id"`
	MeetingID    string `json:"meetingid"`
}

type GetTranscriptParagraphListResponse struct {
	common.Response
	AudioDetect int                   `json:"audio_detect,omitempty"`
	Paragraphs  []TranscriptParagraph `json:"paragraphs,omitempty"`
}

type TranscriptParagraph struct {
	PID       string `json:"pid,omitempty"`
	StartTime int64  `json:"start_time,omitempty"`
	EndTime   int64  `json:"end_time,omitempty"`
}

// GetTranscriptDetailRequest 获取录制转写详情
type GetTranscriptDetailRequest struct {
	RecordFileID string `json:"record_file_id"`
	MeetingID    string `json:"meetingid"`
	PID          string `json:"pid,omitempty"`
	Limit        int    `json:"limit,omitempty"`
}

type GetTranscriptDetailResponse struct {
	common.Response
	HasMore     bool             `json:"has_more,omitempty"`
	Transcripts TranscriptDetail `json:"transcripts,omitempty"`
}

type TranscriptDetail struct {
	Paragraphs  []TranscriptParagraphDetail `json:"paragraphs,omitempty"`
	Keywords    []string                    `json:"keywords,omitempty"`
	AudioDetect int                         `json:"audio_detect,omitempty"`
}

type TranscriptParagraphDetail struct {
	PID         string                 `json:"pid,omitempty"`
	StartTime   int64                  `json:"start_time,omitempty"`
	EndTime     int64                  `json:"end_time,omitempty"`
	SpeakerInfo *TranscriptSpeakerInfo `json:"speaker_info,omitempty"`
	Sentences   []TranscriptSentence   `json:"sentences,omitempty"`
}

type TranscriptSpeakerInfo struct {
	UserID string `json:"userid,omitempty"`
}

type TranscriptSentence struct {
	SID       string           `json:"sid,omitempty"`
	StartTime int64            `json:"start_time,omitempty"`
	EndTime   int64            `json:"end_time,omitempty"`
	Words     []TranscriptWord `json:"words,omitempty"`
}

type TranscriptWord struct {
	WID       string `json:"wid,omitempty"`
	StartTime int64  `json:"start_time,omitempty"`
	EndTime   int64  `json:"end_time,omitempty"`
	Text      string `json:"text,omitempty"`
}

// ================ 高级功能账号 (VIP) 接口 ================

// SubmitBatchAddVIPRequest 分配高级功能账号
type SubmitBatchAddVIPRequest struct {
	UserIDList []string `json:"userid_list"`
}

type SubmitBatchAddVIPResponse struct {
	common.Response
	JobID             string   `json:"jobid,omitempty"`
	InvalidUserIDList []string `json:"invalid_userid_list,omitempty"`
}

type BatchAddJobResultRequest struct {
	JobID string `json:"jobid"`
}

type BatchAddJobResultResponse struct {
	common.Response
	JobResult *VIPJobResult `json:"job_result,omitempty"`
}

type VIPJobResult struct {
	SuccUserIDList []string `json:"succ_userid_list,omitempty"`
	FailUserIDList []string `json:"fail_userid_list,omitempty"`
}

// SubmitBatchDelVIPRequest 取消高级功能账号
type SubmitBatchDelVIPRequest struct {
	UserIDList []string `json:"userid_list"`
}

type SubmitBatchDelVIPResponse struct {
	common.Response
	JobID             string   `json:"jobid,omitempty"`
	InvalidUserIDList []string `json:"invalid_userid_list,omitempty"`
}

type BatchDelJobResultRequest struct {
	JobID string `json:"jobid"`
}

type BatchDelJobResultResponse struct {
	common.Response
	JobResult *VIPJobResult `json:"job_result,omitempty"`
}

// GetVIPListRequest 获取高级功能账号列表
type GetVIPListRequest struct {
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type GetVIPListResponse struct {
	common.Response
	HasMore    bool     `json:"has_more,omitempty"`
	NextCursor string   `json:"next_cursor,omitempty"`
	UserIDList []string `json:"userid_list,omitempty"`
}
