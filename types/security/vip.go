package security

// SubmitBatchAddVIPJobRequest 批量分配高级功能账号请求
type SubmitBatchAddVIPJobRequest struct {
	UserIDList []string `json:"userid_list"` // 要分配高级功能的企业成员userid列表，单次最多100个
}

// SubmitBatchAddVIPJobResponse 批量分配高级功能账号响应
type SubmitBatchAddVIPJobResponse struct {
	ErrCode           int      `json:"errcode"`
	ErrMsg            string   `json:"errmsg"`
	JobID             string   `json:"jobid"`                         // 批量分配高级功能的任务id
	InvalidUserIDList []string `json:"invalid_userid_list,omitempty"` // 非法的userid列表
}

// BatchAddVIPJobResultRequest 查询分配高级功能账号结果请求
type BatchAddVIPJobResultRequest struct {
	JobID string `json:"jobid"` // 批量分配高级功能的任务id
}

// BatchAddVIPJobResultResponse 查询分配高级功能账号结果响应
type BatchAddVIPJobResultResponse struct {
	ErrCode   int       `json:"errcode"`
	ErrMsg    string    `json:"errmsg"`
	JobResult JobResult `json:"job_result"` // 执行任务结果详情
}

// SubmitBatchDelVIPJobRequest 批量取消高级功能账号请求
type SubmitBatchDelVIPJobRequest struct {
	UserIDList []string `json:"userid_list"` // 要撤销分配高级功能的企业成员userid列表，单次最多100个
}

// SubmitBatchDelVIPJobResponse 批量取消高级功能账号响应
type SubmitBatchDelVIPJobResponse struct {
	ErrCode           int      `json:"errcode"`
	ErrMsg            string   `json:"errmsg"`
	JobID             string   `json:"jobid"`                         // 批量取消高级功能的任务id
	InvalidUserIDList []string `json:"invalid_userid_list,omitempty"` // 非法的userid列表
}

// BatchDelVIPJobResultRequest 查询取消高级功能账号结果请求
type BatchDelVIPJobResultRequest struct {
	JobID string `json:"jobid"` // 批量取消高级功能的任务id
}

// BatchDelVIPJobResultResponse 查询取消高级功能账号结果响应
type BatchDelVIPJobResultResponse struct {
	ErrCode   int       `json:"errcode"`
	ErrMsg    string    `json:"errmsg"`
	JobResult JobResult `json:"job_result"` // 执行任务结果详情
}

// JobResult 任务执行结果
type JobResult struct {
	SuccUserIDList []string `json:"succ_userid_list,omitempty"` // 成功的用户列表
	FailUserIDList []string `json:"fail_userid_list,omitempty"` // 失败的用户列表
}

// ListVIPRequest 获取高级功能账号列表请求
type ListVIPRequest struct {
	Cursor string `json:"cursor,omitempty"` // 分页游标
	Limit  int    `json:"limit,omitempty"`  // 每次请求返回的数据上限，默认100，最大200
}

// ListVIPResponse 获取高级功能账号列表响应
type ListVIPResponse struct {
	ErrCode    int      `json:"errcode"`
	ErrMsg     string   `json:"errmsg"`
	HasMore    bool     `json:"has_more"`    // 是否还有更多数据未获取
	NextCursor string   `json:"next_cursor"` // 下一次请求的cursor值
	UserIDList []string `json:"userid_list"` // 符合条件的企业成员userid列表
}
