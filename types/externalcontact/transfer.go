package externalcontact

// OnJobTransferCustomerRequest 分配在职成员的客户请求
type OnJobTransferCustomerRequest struct {
	HandoverUserID     string   `json:"handover_userid"`
	TakeoverUserID     string   `json:"takeover_userid"`
	ExternalUserID     []string `json:"external_userid"`
	TransferSuccessMsg string   `json:"transfer_success_msg,omitempty"`
}

// OnJobCustomerTransferResult 在职客户转接结果
type OnJobCustomerTransferResult struct {
	ExternalUserID string `json:"external_userid"`
	ErrCode        int    `json:"errcode"`
}

// OnJobTransferCustomerResponse 分配在职成员的客户响应
type OnJobTransferCustomerResponse struct {
	Customer []OnJobCustomerTransferResult `json:"customer"`
}

// OnJobTransferGroupChatRequest 分配在职成员的客户群请求
type OnJobTransferGroupChatRequest struct {
	ChatIDList []string `json:"chat_id_list"`
	NewOwner   string   `json:"new_owner"`
}

// FailedChatItem 转接失败的客户群
type FailedChatItem struct {
	ChatID  string `json:"chat_id"`
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// OnJobTransferGroupChatResponse 分配在职成员的客户群响应
type OnJobTransferGroupChatResponse struct {
	FailedChatList []FailedChatItem `json:"failed_chat_list,omitempty"`
}

// TransferResultRequest 查询客户接替状态请求
type TransferResultRequest struct {
	HandoverUserID string `json:"handover_userid"`
	TakeoverUserID string `json:"takeover_userid"`
	Cursor         string `json:"cursor,omitempty"`
}

// CustomerTransferStatus 客户接替状态
type CustomerTransferStatus struct {
	ExternalUserID string `json:"external_userid"`
	Status         int    `json:"status"`
	TakeoverTime   int64  `json:"takeover_time"`
}

// TransferResultResponse 查询客户接替状态响应
type TransferResultResponse struct {
	Customer   []CustomerTransferStatus `json:"customer"`
	NextCursor string                   `json:"next_cursor,omitempty"`
}
