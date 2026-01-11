package meeting

import "github.com/shuaidd/wecom-core/types/common"

// ListMeetingRoomsRequest 获取Rooms会议室列表请求
type ListMeetingRoomsRequest struct {
	MeetingRoomName string `json:"meeting_room_name,omitempty"`
	Cursor          string `json:"cursor,omitempty"`
	Limit           int    `json:"limit,omitempty"`
}

// MeetingRoom Rooms 会议室对象
type MeetingRoom struct {
	MeetingRoomID       string `json:"meeting_room_id,omitempty"`
	MeetingRoomName     string `json:"meeting_room_name,omitempty"`
	MeetingRoomLocation string `json:"meeting_room_location,omitempty"`
	AccountType         int    `json:"account_type,omitempty"`
	ActiveCode          string `json:"active_code,omitempty"`
	ParticipantNumber   int    `json:"participant_number,omitempty"`
	MeetingRoomStatus   int    `json:"meeting_room_status,omitempty"`
	ScheduledStatus     int    `json:"scheduled_status,omitempty"`
	IsAllowCall         bool   `json:"is_allow_call,omitempty"`
}

// ListMeetingRoomsResponse 获取Rooms会议室列表响应
type ListMeetingRoomsResponse struct {
	common.Response
	HasMore         bool          `json:"has_more,omitempty"`
	NextCursor      string        `json:"next_cursor,omitempty"`
	MeetingRoomList []MeetingRoom `json:"meeting_room_list,omitempty"`
}

// ListMeetingsRequest 获取Rooms会议室下的会议列表请求
type ListMeetingsRequest struct {
	MeetingRoomID string `json:"meeting_room_id,omitempty"`
	RoomsID       string `json:"rooms_id,omitempty"`
	StartTime     int64  `json:"start_time,omitempty"`
	EndTime       int64  `json:"end_time,omitempty"`
	Cursor        string `json:"cursor,omitempty"`
	Limit         int    `json:"limit,omitempty"`
}

// MeetingInfo Rooms 下的单个会议信息
type MeetingInfo struct {
	MeetingID   string `json:"meetingid,omitempty"`
	MeetingCode string `json:"meeting_code,omitempty"`
	Subject     string `json:"subject,omitempty"`
	Status      string `json:"status,omitempty"`
	MeetingType int    `json:"meeting_type,omitempty"`
	StartTime   int64  `json:"start_time,omitempty"`
	EndTime     int64  `json:"end_time,omitempty"`
}

// ListMeetingsResponse 获取Rooms会议室下的会议列表响应
type ListMeetingsResponse struct {
	common.Response
	HasMore         bool          `json:"has_more,omitempty"`
	NextCursor      string        `json:"next_cursor,omitempty"`
	MeetingInfoList []MeetingInfo `json:"meeting_info_list,omitempty"`
}

// GetRoomInfoRequest 获取Rooms会议室详情请求
type GetRoomInfoRequest struct {
	MeetingRoomID string `json:"meeting_room_id"`
}

// BasicInfo Rooms 基本信息
type BasicInfo struct {
	RoomsIDList       []string `json:"rooms_id_list,omitempty"`
	MeetingRoomName   string   `json:"meeting_room_name,omitempty"`
	City              string   `json:"city,omitempty"`
	Building          string   `json:"building,omitempty"`
	Floor             string   `json:"floor,omitempty"`
	ParticipantNumber int      `json:"participant_number,omitempty"`
	Device            string   `json:"device,omitempty"`
	Desc              string   `json:"desc,omitempty"`
	Password          string   `json:"password,omitempty"`
}

// AccountInfo Rooms 账号信息
type AccountInfo struct {
	AccountType int    `json:"account_type,omitempty"`
	ValidPeriod string `json:"valid_period,omitempty"`
}

// HardwareInfo Rooms 硬件信息
type HardwareInfo struct {
	Factory           string `json:"factory,omitempty"`
	DeviceModel       string `json:"device_model,omitempty"`
	SN                string `json:"sn,omitempty"`
	IP                string `json:"ip,omitempty"`
	MAC               string `json:"mac,omitempty"`
	RoomsVersion      string `json:"rooms_version,omitempty"`
	FirmwareVersion   string `json:"firmware_version,omitempty"`
	HealthStatus      string `json:"health_status,omitempty"`
	SystemType        string `json:"system_type,omitempty"`
	MeetingRoomStatus int    `json:"meeting_room_status,omitempty"`
	ActiveTime        string `json:"active_time,omitempty"`
	CPUInfo           string `json:"cpu_info,omitempty"`
	CPUUsage          string `json:"cpu_usage,omitempty"`
	GPUInfo           string `json:"gpu_info,omitempty"`
	NetType           string `json:"net_type,omitempty"`
	MemoryInfo        string `json:"memory_info,omitempty"`
	MonitorFrequency  int    `json:"monitor_frequency,omitempty"`
	CameraModel       string `json:"camera_model,omitempty"`
	EnableVideoMirror bool   `json:"enable_video_mirror,omitempty"`
	MicrophoneInfo    string `json:"microphone_info,omitempty"`
	SpeakerInfo       string `json:"speaker_info,omitempty"`
}

// PMIInfo Rooms PMI 信息
type PMIInfo struct {
	PMICode string `json:"pmi_code,omitempty"`
	PMIPwd  string `json:"pmi_pwd,omitempty"`
}

// GetRoomInfoResponse 获取Rooms会议室详情响应
type GetRoomInfoResponse struct {
	common.Response
	BasicInfo       *BasicInfo    `json:"basic_info,omitempty"`
	AccountInfo     *AccountInfo  `json:"account_info,omitempty"`
	HardwareInfo    *HardwareInfo `json:"hardware_info,omitempty"`
	PMIInfo         *PMIInfo      `json:"pmi_info,omitempty"`
	MonitorStatus   int           `json:"monitor_status,omitempty"`
	IsAllowCall     bool          `json:"is_allow_call,omitempty"`
	ScheduledStatus int           `json:"scheduled_status,omitempty"`
}

// GetRoomConfigRequest 获取Rooms会议室配置项请求
type GetRoomConfigRequest struct {
	MeetingRoomID string `json:"meeting_room_id"`
}

// MeetingSettings Rooms 会议配置项
type MeetingSettings struct {
	WaterMark        int  `json:"water_mark,omitempty"`
	AutoResponse     int  `json:"auto_response,omitempty"`
	Caption          bool `json:"caption,omitempty"`
	RoomPMI          bool `json:"room_pmi,omitempty"`
	RoomNotification bool `json:"room_notification,omitempty"`
}

// RecordSettings Rooms 录制配置
type RecordSettings struct {
	ShareRecord    int  `json:"share_record,omitempty"`
	DownloadRecord bool `json:"download_record,omitempty"`
}

// GetRoomConfigResponse 获取Rooms会议室配置项响应
type GetRoomConfigResponse struct {
	common.Response
	MeetingSettings *MeetingSettings `json:"meeting_settings,omitempty"`
	RecordSettings  *RecordSettings  `json:"record_settings,omitempty"`
}

// GetInventoryResponse 获取Rooms会议室资源响应
type GetInventoryResponse struct {
	common.Response
	NormalCount         int `json:"normal_count,omitempty"`
	SpecialCount        int `json:"special_count,omitempty"`
	NormalUsedCount     int `json:"normal_used_count,omitempty"`
	SpecialUsedCount    int `json:"special_used_count,omitempty"`
	NormalExpiredCount  int `json:"normal_expired_count,omitempty"`
	SpecialExpiredCount int `json:"special_expired_count,omitempty"`
}

// ListControllersRequest 获取控制器列表请求
type ListControllersRequest struct {
	ControllerName string `json:"controller_name,omitempty"`
	Cursor         string `json:"cursor,omitempty"`
	Limit          int    `json:"limit,omitempty"`
}

// ControllerInfo 控制器信息
type ControllerInfo struct {
	RoomsID             string `json:"rooms_id,omitempty"`
	MeetingRoomName     string `json:"meeting_room_name,omitempty"`
	MeetingRoomLocation string `json:"meeting_room_location,omitempty"`
	ControllerName      string `json:"controller_name,omitempty"`
	ManufactureName     string `json:"manufacture_name,omitempty"`
	ControllerModel     string `json:"controller_model,omitempty"`
	AppVersion          string `json:"app_version,omitempty"`
	Status              string `json:"status,omitempty"`
	FrameworkVersion    string `json:"framework_version,omitempty"`
	IPAddress           string `json:"ip_address,omitempty"`
	MACAddress          string `json:"mac_address,omitempty"`
	CPUType             string `json:"cpu_type,omitempty"`
	CPUUsage            string `json:"cpu_usage,omitempty"`
	NetworkType         string `json:"network_type,omitempty"`
	MemUsage            string `json:"mem_usage,omitempty"`
}

// ListControllersResponse 获取控制器列表响应
type ListControllersResponse struct {
	common.Response
	HasMore            bool             `json:"has_more,omitempty"`
	NextCursor         string           `json:"next_cursor,omitempty"`
	ControllerInfoList []ControllerInfo `json:"controller_info_list,omitempty"`
}

// ListDevicesRequest 获取设备列表请求
type ListDevicesRequest struct {
	MeetingRoomName string `json:"meeting_room_name,omitempty"`
	Cursor          string `json:"cursor,omitempty"`
	Limit           int    `json:"limit,omitempty"`
}

// DeviceMonitorInfo 设备健康信息
type DeviceMonitorInfo struct {
	CameraStatus     bool `json:"camera_status,omitempty"`
	MicrophoneStatus bool `json:"microphone_status,omitempty"`
	SpeakerStatus    bool `json:"speaker_status,omitempty"`
}

// DeviceInfo 设备信息
type DeviceInfo struct {
	MeetingRoomID       string             `json:"meeting_room_id,omitempty"`
	RoomsID             string             `json:"rooms_id,omitempty"`
	MeetingRoomName     string             `json:"meeting_room_name,omitempty"`
	MeetingRoomLocation string             `json:"meeting_room_location,omitempty"`
	DeviceModel         string             `json:"device_model,omitempty"`
	AppVersion          string             `json:"app_version,omitempty"`
	MeetingRoomStatus   int                `json:"meeting_room_status,omitempty"`
	DeviceMonitorInfo   *DeviceMonitorInfo `json:"device_monitor_info,omitempty"`
}

// ListDevicesResponse 获取设备列表响应
type ListDevicesResponse struct {
	common.Response
	HasMore        bool         `json:"has_more,omitempty"`
	NextCursor     string       `json:"next_cursor,omitempty"`
	DeviceInfoList []DeviceInfo `json:"device_info_list,omitempty"`
}

// BookRoomsRequest 预定Rooms会议室请求
type BookRoomsRequest struct {
	MeetingID         string   `json:"meetingid"`
	MeetingRoomIDList []string `json:"meeting_room_id_list"`
	SubjectVisible    bool     `json:"subject_visible,omitempty"`
}

// BookRoomsResponse 预定Rooms会议室响应
type BookRoomsResponse struct {
	common.Response
	MeetingRoomList []MeetingRoom `json:"meeting_room_list,omitempty"`
}

// ReleaseRoomsRequest 释放Rooms会议室请求
type ReleaseRoomsRequest struct {
	MeetingID         string   `json:"meetingid"`
	MeetingRoomIDList []string `json:"meeting_room_id_list"`
}

// MRAAddress MRA 对象
type MRAAddress struct {
	Protocol   int    `json:"protocol"`
	DialString string `json:"dial_string"`
}

// CallRoomRequest 呼叫 Rooms 会议室请求
type CallRoomRequest struct {
	MeetingID     string      `json:"meetingid"`
	MeetingRoomID string      `json:"meeting_room_id,omitempty"`
	MRAAddress    *MRAAddress `json:"mra_address,omitempty"`
}

// CallRoomResponse 呼叫 Rooms 会议室响应
type CallRoomResponse struct {
	common.Response
	InviteID string `json:"invite_id,omitempty"`
}

// CancelCallRequest 取消呼叫 Rooms 请求
type CancelCallRequest struct {
	MeetingID     string      `json:"meetingid"`
	InviteID      string      `json:"invite_id"`
	MeetingRoomID string      `json:"meeting_room_id,omitempty"`
	MRAAddress    *MRAAddress `json:"mra_address,omitempty"`
}

// GetResponseStatusRequest 获取 Rooms 会议室应答状态请求
type GetResponseStatusRequest struct {
	MeetingID     string      `json:"meetingid"`
	MeetingRoomID string      `json:"meeting_room_id,omitempty"`
	MRAAddress    *MRAAddress `json:"mra_address,omitempty"`
}

// GetResponseStatusResponse 获取 Rooms 会议室应答状态响应
type GetResponseStatusResponse struct {
	common.Response
	Status       int    `json:"status,omitempty"`
	ResponseTime string `json:"response_time,omitempty"`
}
