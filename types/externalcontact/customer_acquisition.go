package externalcontact

// AcquisitionLink 获客链接信息
type AcquisitionLink struct {
	LinkID     string `json:"link_id,omitempty"`     // 获客链接ID
	LinkName   string `json:"link_name"`             // 获客链接名称
	URL        string `json:"url,omitempty"`         // 获客链接实际URL
	CreateTime int64  `json:"create_time,omitempty"` // 创建时间戳
	SkipVerify bool   `json:"skip_verify,omitempty"` // 是否无需验证
	MarkSource bool   `json:"mark_source,omitempty"` // 是否标记客户添加来源
}

// AcquisitionRange 获客链接使用范围
type AcquisitionRange struct {
	UserList       []string `json:"user_list,omitempty"`       // 使用范围成员列表
	DepartmentList []int    `json:"department_list,omitempty"` // 使用范围部门列表
}

// PriorityOption 优先分配选项
type PriorityOption struct {
	PriorityType       int      `json:"priority_type,omitempty"`        // 优先分配类型: 1-全企业范围 2-指定范围
	PriorityUserIDList []string `json:"priority_userid_list,omitempty"` // 指定成员列表
}

// ListAcquisitionLinkRequest 获取获客链接列表请求
type ListAcquisitionLinkRequest struct {
	Limit  int    `json:"limit,omitempty"`  // 返回的最大记录数，最大值100
	Cursor string `json:"cursor,omitempty"` // 分页游标
}

// ListAcquisitionLinkResponse 获取获客链接列表响应
type ListAcquisitionLinkResponse struct {
	LinkIDList []string `json:"link_id_list"` // 获客链接ID列表
	NextCursor string   `json:"next_cursor"`  // 分页游标
}

// GetAcquisitionLinkRequest 获取获客链接详情请求
type GetAcquisitionLinkRequest struct {
	LinkID string `json:"link_id"` // 获客链接ID
}

// GetAcquisitionLinkResponse 获取获客链接详情响应
type GetAcquisitionLinkResponse struct {
	Link           AcquisitionLink  `json:"link"`                      // 获客链接信息
	Range          AcquisitionRange `json:"range"`                     // 使用范围
	PriorityOption *PriorityOption  `json:"priority_option,omitempty"` // 优先分配选项
}

// CreateAcquisitionLinkRequest 创建获客链接请求
type CreateAcquisitionLinkRequest struct {
	LinkName       string           `json:"link_name"`                 // 链接名称,最长30个字符
	Range          AcquisitionRange `json:"range"`                     // 使用范围
	SkipVerify     bool             `json:"skip_verify,omitempty"`     // 是否无需验证
	PriorityOption *PriorityOption  `json:"priority_option,omitempty"` // 优先分配选项
	MarkSource     bool             `json:"mark_source,omitempty"`     // 是否标记客户添加来源
}

// CreateAcquisitionLinkResponse 创建获客链接响应
type CreateAcquisitionLinkResponse struct {
	Link AcquisitionLink `json:"link"` // 创建的获客链接信息
}

// UpdateAcquisitionLinkRequest 编辑获客链接请求
type UpdateAcquisitionLinkRequest struct {
	LinkID         string            `json:"link_id"`                   // 获客链接ID
	LinkName       string            `json:"link_name,omitempty"`       // 链接名称,最长30个字符
	Range          *AcquisitionRange `json:"range,omitempty"`           // 使用范围
	SkipVerify     bool              `json:"skip_verify,omitempty"`     // 是否无需验证
	PriorityOption *PriorityOption   `json:"priority_option,omitempty"` // 优先分配选项
	MarkSource     bool              `json:"mark_source,omitempty"`     // 是否标记客户添加来源
}

// DeleteAcquisitionLinkRequest 删除获客链接请求
type DeleteAcquisitionLinkRequest struct {
	LinkID string `json:"link_id"` // 获客链接ID
}

// QuotaInfo 额度信息
type QuotaInfo struct {
	ExpireDate int64 `json:"expire_date"` // 额度过期时间戳
	Balance    int   `json:"balance"`     // 额度数量
}

// GetAcquisitionQuotaResponse 查询剩余使用量响应
type GetAcquisitionQuotaResponse struct {
	Total     int         `json:"total"`      // 历史累计使用量
	Balance   int         `json:"balance"`    // 剩余使用量
	QuotaList []QuotaInfo `json:"quota_list"` // 额度列表
}

// GetAcquisitionStatisticRequest 查询链接使用详情请求
type GetAcquisitionStatisticRequest struct {
	LinkID    string `json:"link_id"`    // 获客链接ID
	StartTime int64  `json:"start_time"` // 统计起始时间戳
	EndTime   int64  `json:"end_time"`   // 统计结束时间戳
}

// GetAcquisitionStatisticResponse 查询链接使用详情响应
type GetAcquisitionStatisticResponse struct {
	ClickLinkCustomerCnt int `json:"click_link_customer_cnt"` // 点击链接客户数
	NewCustomerCnt       int `json:"new_customer_cnt"`        // 新增客户数
}

// ChatInfo 会话信息
type ChatInfo struct {
	RecvMsgCnt int    `json:"recv_msg_cnt"` // 成员收到的此客户的消息次数
	LinkID     string `json:"link_id"`      // 成员添加客户的获客链接ID
	State      string `json:"state"`        // 成员添加客户的state
}

// GetAcquisitionChatInfoRequest 获取成员多次收消息详情请求
type GetAcquisitionChatInfoRequest struct {
	ChatKey string `json:"chat_key"` // 会话信息凭据ChatKey
}

// GetAcquisitionChatInfoResponse 获取成员多次收消息详情响应
type GetAcquisitionChatInfoResponse struct {
	UserID         string   `json:"userid"`          // 成员的userid
	ExternalUserID string   `json:"external_userid"` // 客户ID
	ChatInfo       ChatInfo `json:"chat_info"`       // 会话信息
}

// AcquisitionCustomer 获客客户信息
type AcquisitionCustomer struct {
	ExternalUserID string `json:"external_userid"` // 客户external_userid
	UserID         string `json:"userid"`          // 添加客户的跟进人userid
	ChatStatus     int    `json:"chat_status"`     // 会话状态: 0-未发消息 1-已发消息 2-未知
	State          string `json:"state"`           // 获客链接的state参数
}

// ListAcquisitionCustomerRequest 获取获客客户列表请求
type ListAcquisitionCustomerRequest struct {
	LinkID string `json:"link_id"`          // 获客链接ID
	Limit  int    `json:"limit,omitempty"`  // 返回的最大记录数，最大值1000
	Cursor string `json:"cursor,omitempty"` // 分页游标
}

// ListAcquisitionCustomerResponse 获取获客客户列表响应
type ListAcquisitionCustomerResponse struct {
	CustomerList []AcquisitionCustomer `json:"customer_list"` // 客户列表
	NextCursor   string                `json:"next_cursor"`   // 分页游标
}
