package contact

import "github.com/shuaidd/wecom-core/types/common"

// Callback 回调信息
type Callback struct {
	URL            string `json:"url,omitempty"`
	Token          string `json:"token,omitempty"`
	EncodingAESKey string `json:"encodingaeskey,omitempty"`
}

// SyncUsersRequest 增量更新成员请求
type SyncUsersRequest struct {
	MediaID  string    `json:"media_id"`
	ToInvite bool      `json:"to_invite,omitempty"`
	Callback *Callback `json:"callback,omitempty"`
}

// SyncUsersResponse 增量更新成员响应
type SyncUsersResponse struct {
	common.Response
	JobID string `json:"jobid"`
}

// ReplaceUsersRequest 全量覆盖成员请求
type ReplaceUsersRequest struct {
	MediaID  string    `json:"media_id"`
	ToInvite bool      `json:"to_invite,omitempty"`
	Callback *Callback `json:"callback,omitempty"`
}

// ReplaceUsersResponse 全量覆盖成员响应
type ReplaceUsersResponse struct {
	common.Response
	JobID string `json:"jobid"`
}

// ReplaceDepartmentsRequest 全量覆盖部门请求
type ReplaceDepartmentsRequest struct {
	MediaID  string    `json:"media_id"`
	Callback *Callback `json:"callback,omitempty"`
}

// ReplaceDepartmentsResponse 全量覆盖部门响应
type ReplaceDepartmentsResponse struct {
	common.Response
	JobID string `json:"jobid"`
}

// BatchUserResult 批量成员操作结果
type BatchUserResult struct {
	UserID  string `json:"userid"`
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// BatchDepartmentResult 批量部门操作结果
type BatchDepartmentResult struct {
	Action  int    `json:"action"`
	PartyID int    `json:"partyid"`
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// GetBatchResultResponse 获取异步任务结果响应
type GetBatchResultResponse struct {
	common.Response
	Status     int         `json:"status"`
	Type       string      `json:"type"`
	Total      int         `json:"total"`
	Percentage int         `json:"percentage"`
	Result     interface{} `json:"result,omitempty"`
}
