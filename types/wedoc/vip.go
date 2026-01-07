package wedoc

// ==================== 高级功能账号管理相关类型 ====================

// BatchAddVipRequest 分配高级功能账号请求
type BatchAddVipRequest struct {
	UserIDList []string `json:"userid_list"` // 要分配高级功能的企业成员userid列表，单次操作最大限制100个
}

// BatchAddVipResponse 分配高级功能账号响应
type BatchAddVipResponse struct {
	SuccUserIDList []string `json:"succ_userid_list"` // 分配成功的userid列表，包括已经是高级功能账号的userid
	FailUserIDList []string `json:"fail_userid_list"` // 分配失败的userid列表
}

// ListVipRequest 获取高级功能账号列表请求
type ListVipRequest struct {
	Cursor string `json:"cursor,omitempty"` // 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Limit  uint32 `json:"limit,omitempty"`  // 用于分页查询，每次请求返回的数据上限。默认100，最大200
}

// ListVipResponse 获取高级功能账号列表响应
type ListVipResponse struct {
	HasMore    bool     `json:"has_more"`    // 是否还有更多数据未获取
	NextCursor string   `json:"next_cursor"` // 下一次请求的cursor值
	UserIDList []string `json:"userid_list"` // 符合条件的企业成员userid列表
}

// BatchDelVipRequest 取消高级功能账号请求
type BatchDelVipRequest struct {
	UserIDList []string `json:"userid_list"` // 要撤销分配高级功能的企业成员userid列表，单次操作最多限制100个
}

// BatchDelVipResponse 取消高级功能账号响应
type BatchDelVipResponse struct {
	SuccUserIDList []string `json:"succ_userid_list"` // 撤销分配成功的userid列表
	FailUserIDList []string `json:"fail_userid_list"` // 撤销分配失败的userid列表
}
