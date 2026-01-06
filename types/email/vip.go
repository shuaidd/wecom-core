package email

// BatchAddVIPRequest 分配高级功能账号请求
type BatchAddVIPRequest struct {
	UserIDList []string `json:"userid_list"` // 要分配高级功能的企业成员userid列表，单次操作最大限制100个
}

// BatchAddVIPResponse 分配高级功能账号响应
type BatchAddVIPResponse struct {
	ErrCode        int32    `json:"errcode"`
	ErrMsg         string   `json:"errmsg"`
	SuccUserIDList []string `json:"succ_userid_list"` // 分配成功的用户列表，包括之前已经分配过的用户
	FailUserIDList []string `json:"fail_userid_list"` // 分配失败的用户列表
}

// BatchDelVIPRequest 取消高级功能账号请求
type BatchDelVIPRequest struct {
	UserIDList []string `json:"userid_list"` // 要撤销分配高级功能的企业成员userid列表，单次操作最多限制100个
}

// BatchDelVIPResponse 取消高级功能账号响应
type BatchDelVIPResponse struct {
	ErrCode        int32    `json:"errcode"`
	ErrMsg         string   `json:"errmsg"`
	SuccUserIDList []string `json:"succ_userid_list"` // 撤销分配成功的用户列表
	FailUserIDList []string `json:"fail_userid_list"` // 撤销分配失败的用户列表
}

// ListVIPRequest 获取高级功能账号列表请求
type ListVIPRequest struct {
	Cursor string `json:"cursor,omitempty"` // 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Limit  uint32 `json:"limit,omitempty"`  // 用于分页查询，每次请求返回的数据上限。默认100，最大200
}

// ListVIPResponse 获取高级功能账号列表响应
type ListVIPResponse struct {
	ErrCode    int32    `json:"errcode"`
	ErrMsg     string   `json:"errmsg"`
	HasMore    bool     `json:"has_more"`    // 是否还有更多数据未获取
	NextCursor string   `json:"next_cursor"` // 下一次请求的cursor值
	UserIDList []string `json:"userid_list"` // 符合条件的企业成员userid列表
}
