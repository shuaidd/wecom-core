package externalcontact

// MomentText 朋友圈文本内容
type MomentText struct {
	Content string `json:"content,omitempty"`
}

// MomentImage 朋友圈图片附件
type MomentImage struct {
	MediaID string `json:"media_id"`
}

// MomentVideo 朋友圈视频附件
type MomentVideo struct {
	MediaID      string `json:"media_id"`
	ThumbMediaID string `json:"thumb_media_id,omitempty"`
}

// MomentLink 朋友圈链接附件
type MomentLink struct {
	Title   string `json:"title,omitempty"`
	URL     string `json:"url"`
	MediaID string `json:"media_id"`
}

// MomentLocation 朋友圈地理位置
type MomentLocation struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Name      string `json:"name"`
}

// MomentAttachment 朋友圈附件
type MomentAttachment struct {
	MsgType string       `json:"msgtype"`
	Image   *MomentImage `json:"image,omitempty"`
	Video   *MomentVideo `json:"video,omitempty"`
	Link    *MomentLink  `json:"link,omitempty"`
}

// SenderList 发表任务的执行者列表
type SenderList struct {
	UserList       []string `json:"user_list,omitempty"`
	DepartmentList []int    `json:"department_list,omitempty"`
}

// ExternalContactList 可见到该朋友圈的客户列表
type ExternalContactList struct {
	TagList []string `json:"tag_list,omitempty"`
}

// VisibleRange 可见范围
type VisibleRange struct {
	SenderList          *SenderList          `json:"sender_list,omitempty"`
	ExternalContactList *ExternalContactList `json:"external_contact_list,omitempty"`
}

// AddMomentTaskRequest 创建发表任务请求
type AddMomentTaskRequest struct {
	Text         *MomentText         `json:"text,omitempty"`
	Attachments  []MomentAttachment  `json:"attachments,omitempty"`
	VisibleRange *VisibleRange       `json:"visible_range,omitempty"`
}

// AddMomentTaskResponse 创建发表任务响应
type AddMomentTaskResponse struct {
	JobID string `json:"jobid"`
}

// InvalidSenderList 不合法的执行者列表
type InvalidSenderList struct {
	UserList       []string `json:"user_list,omitempty"`
	DepartmentList []int    `json:"department_list,omitempty"`
}

// InvalidExternalContactList 不合法的客户列表
type InvalidExternalContactList struct {
	TagList []string `json:"tag_list,omitempty"`
}

// MomentTaskResult 朋友圈任务结果
type MomentTaskResult struct {
	ErrCode                    int                         `json:"errcode"`
	ErrMsg                     string                      `json:"errmsg"`
	MomentID                   string                      `json:"moment_id,omitempty"`
	InvalidSenderList          *InvalidSenderList          `json:"invalid_sender_list,omitempty"`
	InvalidExternalContactList *InvalidExternalContactList `json:"invalid_external_contact_list,omitempty"`
}

// GetMomentTaskResultResponse 获取任务创建结果响应
type GetMomentTaskResultResponse struct {
	Status int               `json:"status"`
	Type   string            `json:"type"`
	Result *MomentTaskResult `json:"result,omitempty"`
}

// GetMomentListRequest 获取企业全部的发表列表请求
type GetMomentListRequest struct {
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	Creator    string `json:"creator,omitempty"`
	FilterType int    `json:"filter_type,omitempty"`
	Cursor     string `json:"cursor,omitempty"`
	Limit      int    `json:"limit,omitempty"`
}

// MomentItem 朋友圈项
type MomentItem struct {
	MomentID    string          `json:"moment_id"`
	Creator     string          `json:"creator,omitempty"`
	CreateTime  int64           `json:"create_time"`
	CreateType  int             `json:"create_type"`
	VisibleType int             `json:"visible_type"`
	Text        *MomentText     `json:"text,omitempty"`
	Image       []MomentImage   `json:"image,omitempty"`
	Video       *MomentVideo    `json:"video,omitempty"`
	Link        *MomentLink     `json:"link,omitempty"`
	Location    *MomentLocation `json:"location,omitempty"`
}

// GetMomentListResponse 获取企业全部的发表列表响应
type GetMomentListResponse struct {
	NextCursor string       `json:"next_cursor,omitempty"`
	MomentList []MomentItem `json:"moment_list"`
}

// GetMomentTaskRequest 获取客户朋友圈企业发表的列表请求
type GetMomentTaskRequest struct {
	MomentID string `json:"moment_id"`
	Cursor   string `json:"cursor,omitempty"`
	Limit    int    `json:"limit,omitempty"`
}

// MomentTask 朋友圈发表任务
type MomentTask struct {
	UserID        string `json:"userid"`
	PublishStatus int    `json:"publish_status"`
}

// GetMomentTaskResponse 获取客户朋友圈企业发表的列表响应
type GetMomentTaskResponse struct {
	NextCursor string       `json:"next_cursor,omitempty"`
	TaskList   []MomentTask `json:"task_list"`
}

// GetMomentCustomerListRequest 获取客户朋友圈发表时选择的可见范围请求
type GetMomentCustomerListRequest struct {
	MomentID string `json:"moment_id"`
	UserID   string `json:"userid"`
	Cursor   string `json:"cursor,omitempty"`
	Limit    int    `json:"limit,omitempty"`
}

// MomentCustomer 朋友圈可见客户
type MomentCustomer struct {
	UserID         string `json:"userid"`
	ExternalUserID string `json:"external_userid"`
}

// GetMomentCustomerListResponse 获取客户朋友圈发表时选择的可见范围响应
type GetMomentCustomerListResponse struct {
	NextCursor   string           `json:"next_cursor,omitempty"`
	CustomerList []MomentCustomer `json:"customer_list"`
}

// GetMomentSendResultRequest 获取客户朋友圈发表后的可见客户列表请求
type GetMomentSendResultRequest struct {
	MomentID string `json:"moment_id"`
	UserID   string `json:"userid"`
	Cursor   string `json:"cursor,omitempty"`
	Limit    int    `json:"limit,omitempty"`
}

// MomentSendCustomer 朋友圈发送成功的客户
type MomentSendCustomer struct {
	ExternalUserID string `json:"external_userid"`
}

// GetMomentSendResultResponse 获取客户朋友圈发表后的可见客户列表响应
type GetMomentSendResultResponse struct {
	NextCursor   string               `json:"next_cursor,omitempty"`
	CustomerList []MomentSendCustomer `json:"customer_list"`
}

// GetMomentCommentsRequest 获取客户朋友圈的互动数据请求
type GetMomentCommentsRequest struct {
	MomentID string `json:"moment_id"`
	UserID   string `json:"userid"`
}

// MomentComment 朋友圈评论/点赞
type MomentComment struct {
	ExternalUserID string `json:"external_userid,omitempty"`
	UserID         string `json:"userid,omitempty"`
	CreateTime     int64  `json:"create_time"`
}

// GetMomentCommentsResponse 获取客户朋友圈的互动数据响应
type GetMomentCommentsResponse struct {
	CommentList []MomentComment `json:"comment_list"`
	LikeList    []MomentComment `json:"like_list"`
}

// CancelMomentTaskRequest 停止发表企业朋友圈请求
type CancelMomentTaskRequest struct {
	MomentID string `json:"moment_id"`
}

// MomentStrategyPrivilege 规则组权限配置
type MomentStrategyPrivilege struct {
	ViewMomentList            bool `json:"view_moment_list,omitempty"`
	SendMoment                bool `json:"send_moment,omitempty"`
	ManageMomentCoverAndSign  bool `json:"manage_moment_cover_and_sign,omitempty"`
}

// MomentStrategy 客户朋友圈规则组
type MomentStrategy struct {
	StrategyID   int                     `json:"strategy_id"`
	ParentID     int                     `json:"parent_id,omitempty"`
	StrategyName string                  `json:"strategy_name"`
	CreateTime   int64                   `json:"create_time,omitempty"`
	AdminList    []string                `json:"admin_list,omitempty"`
	Privilege    *MomentStrategyPrivilege `json:"privilege,omitempty"`
}

// ListMomentStrategyRequest 获取规则组列表请求
type ListMomentStrategyRequest struct {
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

// MomentStrategyItem 规则组列表项
type MomentStrategyItem struct {
	StrategyID int `json:"strategy_id"`
}

// ListMomentStrategyResponse 获取规则组列表响应
type ListMomentStrategyResponse struct {
	Strategy   []MomentStrategyItem `json:"strategy"`
	NextCursor string               `json:"next_cursor,omitempty"`
}

// GetMomentStrategyRequest 获取规则组详情请求
type GetMomentStrategyRequest struct {
	StrategyID int `json:"strategy_id"`
}

// GetMomentStrategyResponse 获取规则组详情响应
type GetMomentStrategyResponse struct {
	Strategy MomentStrategy `json:"strategy"`
}

// GetMomentStrategyRangeRequest 获取规则组管理范围请求
type GetMomentStrategyRangeRequest struct {
	StrategyID int    `json:"strategy_id"`
	Cursor     string `json:"cursor,omitempty"`
	Limit      int    `json:"limit,omitempty"`
}

// MomentStrategyRange 规则组管理范围节点
type MomentStrategyRange struct {
	Type    int    `json:"type"`
	UserID  string `json:"userid,omitempty"`
	PartyID int    `json:"partyid,omitempty"`
}

// GetMomentStrategyRangeResponse 获取规则组管理范围响应
type GetMomentStrategyRangeResponse struct {
	Range      []MomentStrategyRange `json:"range"`
	NextCursor string                `json:"next_cursor,omitempty"`
}

// CreateMomentStrategyRequest 创建新的规则组请求
type CreateMomentStrategyRequest struct {
	ParentID     int                      `json:"parent_id,omitempty"`
	StrategyName string                   `json:"strategy_name"`
	AdminList    []string                 `json:"admin_list"`
	Privilege    *MomentStrategyPrivilege `json:"privilege,omitempty"`
	Range        []MomentStrategyRange    `json:"range"`
}

// CreateMomentStrategyResponse 创建新的规则组响应
type CreateMomentStrategyResponse struct {
	StrategyID int `json:"strategy_id"`
}

// EditMomentStrategyRequest 编辑规则组及其管理范围请求
type EditMomentStrategyRequest struct {
	StrategyID   int                      `json:"strategy_id"`
	StrategyName string                   `json:"strategy_name,omitempty"`
	AdminList    []string                 `json:"admin_list,omitempty"`
	Privilege    *MomentStrategyPrivilege `json:"privilege,omitempty"`
	RangeAdd     []MomentStrategyRange    `json:"range_add,omitempty"`
	RangeDel     []MomentStrategyRange    `json:"range_del,omitempty"`
}

// DeleteMomentStrategyRequest 删除规则组请求
type DeleteMomentStrategyRequest struct {
	StrategyID int `json:"strategy_id"`
}
