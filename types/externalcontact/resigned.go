package externalcontact

// GetUnassignedListRequest 获取待分配的离职成员列表请求
type GetUnassignedListRequest struct {
	Cursor   string `json:"cursor,omitempty"`
	PageSize int    `json:"page_size,omitempty"`
}

// UnassignedInfo 离职成员对应的外部联系人项
type UnassignedInfo struct {
	HandoverUserID string `json:"handover_userid"`
	ExternalUserID string `json:"external_userid"`
	DimissionTime  int64  `json:"dimission_time"`
}

// GetUnassignedListResponse 获取待分配的离职成员列表响应
type GetUnassignedListResponse struct {
	Info       []UnassignedInfo `json:"info,omitempty"`
	IsLast     bool             `json:"is_last"`
	NextCursor string           `json:"next_cursor,omitempty"`
}

// TransferCustomerRequest 分配离职成员的客户请求
type TransferCustomerRequest struct {
	HandoverUserID  string   `json:"handover_userid"`
	TakeoverUserID  string   `json:"takeover_userid"`
	ExternalUserIDs []string `json:"external_userid"`
}

// TransferResult 单个客户分配结果
type TransferResult struct {
	ExternalUserID string `json:"external_userid"`
	ErrCode        int    `json:"errcode"`
}

// TransferCustomerResponse 分配离职成员的客户响应
type TransferCustomerResponse struct {
	Customer []TransferResult `json:"customer"`
}
