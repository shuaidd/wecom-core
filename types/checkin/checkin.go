package checkin

import (
	"encoding/json"

	"github.com/shuaidd/wecom-core/types/common"
)

// SetCheckinScheduleListRequest 为打卡人员排班
type SetCheckinScheduleListRequest struct {
	GroupID   int64                        `json:"groupid"`
	Items     []SetCheckinScheduleListItem `json:"items"`
	YearMonth int                          `json:"yearmonth"`
}

type SetCheckinScheduleListItem struct {
	UserID     string `json:"userid"`
	Day        int    `json:"day"`
	ScheduleID int64  `json:"schedule_id"`
}

// PunchCorrectionRequest 为打卡人员补卡
type PunchCorrectionRequest struct {
	UserID              string `json:"userid"`
	ScheduleDateTime    int64  `json:"schedule_date_time"`
	ScheduleCheckinTime int64  `json:"schedule_checkin_time,omitempty"`
	CheckinTime         int64  `json:"checkin_time"`
	Remark              string `json:"remark,omitempty"`
}

// AddCheckinUserFaceRequest 录入打卡人员人脸信息
type AddCheckinUserFaceRequest struct {
	UserID   string `json:"userid,omitempty"`
	UserFace string `json:"userface,omitempty"`
}

// AddCheckinRecordRequest 添加打卡记录
type AddCheckinRecordRequest struct {
	Records []AddCheckinRecordItem `json:"records"`
}

type AddCheckinRecordItem struct {
	UserID         string   `json:"userid"`
	CheckinTime    int64    `json:"checkin_time"`
	LocationTitle  string   `json:"location_title,omitempty"`
	LocationDetail string   `json:"location_detail,omitempty"`
	MediaIDs       []string `json:"mediaids,omitempty"`
	Notes          string   `json:"notes,omitempty"`
	DeviceType     int      `json:"device_type"`
	Lat            int64    `json:"lat,omitempty"`
	Lng            int64    `json:"lng,omitempty"`
	DeviceDetail   string   `json:"device_detail,omitempty"`
	WifiName       string   `json:"wifiname,omitempty"`
	WifiMac        string   `json:"wifimac,omitempty"`
}

// AddCheckinOptionRequest / UpdateCheckinOptionRequest 管理打卡规则（入参较复杂，使用 json.RawMessage 便于上层构造）
type AddCheckinOptionRequest struct {
	EffectiveNow bool            `json:"effective_now,omitempty"`
	Group        json.RawMessage `json:"group,omitempty"`
}

type UpdateCheckinOptionRequest = AddCheckinOptionRequest

type ClearCheckinOptionArrayFieldRequest struct {
	GroupID      int64 `json:"groupid"`
	ClearField   []int `json:"clear_field"`
	EffectiveNow bool  `json:"effective_now,omitempty"`
}

type DeleteCheckinOptionRequest struct {
	GroupID      int64 `json:"groupid"`
	EffectiveNow bool  `json:"effective_now,omitempty"`
}

// GetCorpCheckinOptionResponse 获取企业所有打卡规则（简化）
type GetCorpCheckinOptionResponse struct {
	common.Response
	Group json.RawMessage `json:"group"`
}

// GetCheckinOptionRequest 获取员工打卡规则
type GetCheckinOptionRequest struct {
	Datetime   int64    `json:"datetime"`
	UserIDList []string `json:"useridlist"`
}

type GetCheckinOptionResponse struct {
	common.Response
	Info json.RawMessage `json:"info"`
}

// GetCheckinScheduleListRequest 获取打卡人员排班信息
type GetCheckinScheduleListRequest struct {
	StartTime  int64    `json:"starttime"`
	EndTime    int64    `json:"endtime"`
	UserIDList []string `json:"useridlist"`
}

type GetCheckinScheduleListResponse struct {
	common.Response
	ScheduleList json.RawMessage `json:"schedule_list"`
}

// GetCheckinDayDataRequest / Response
type GetCheckinDayDataRequest struct {
	StartTime  int64    `json:"starttime"`
	EndTime    int64    `json:"endtime"`
	UserIDList []string `json:"useridlist"`
}

type GetCheckinDayDataResponse struct {
	common.Response
	Datas json.RawMessage `json:"datas"`
}

// GetCheckinMonthDataRequest / Response
type GetCheckinMonthDataRequest struct {
	StartTime  int64    `json:"starttime"`
	EndTime    int64    `json:"endtime"`
	UserIDList []string `json:"useridlist"`
}

type GetCheckinMonthDataResponse struct {
	common.Response
	Datas json.RawMessage `json:"datas"`
}

// GetCheckinDataRequest / Response 获取打卡记录数据
type GetCheckinDataRequest struct {
	OpenCheckinDataType int      `json:"opencheckindatatype"`
	StartTime           int64    `json:"starttime"`
	EndTime             int64    `json:"endtime"`
	UserIDList          []string `json:"useridlist"`
}

type GetCheckinDataResponse struct {
	common.Response
	CheckinData json.RawMessage `json:"checkindata"`
}

// GetHardwareCheckinDataRequest / Response 获取设备打卡数据
type GetHardwareCheckinDataRequest struct {
	FilterType int      `json:"filter_type,omitempty"`
	StartTime  int64    `json:"starttime"`
	EndTime    int64    `json:"endtime"`
	UserIDList []string `json:"useridlist"`
}

type GetHardwareCheckinDataResponse struct {
	common.Response
	CheckinData json.RawMessage `json:"checkindata"`
}
