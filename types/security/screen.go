package security

// GetScreenOperRecordRequest 获取截屏/录屏操作记录请求
type GetScreenOperRecordRequest struct {
	StartTime        int64    `json:"start_time"`                   // 开始时间
	EndTime          int64    `json:"end_time"`                     // 结束时间，开始时间到结束时间的范围不能超过14天
	UserIDList       []string `json:"userid_list,omitempty"`        // 需要查询的截屏操作者的userid，单次最多可以传100个用户
	DepartmentIDList []int    `json:"department_id_list,omitempty"` // 需要查询的截屏操作者部门的department_id，单次最多可以传100个部门id
	ScreenShotType   int      `json:"screen_shot_type,omitempty"`   // 截屏内容的类型，不设置默认为全部
	Cursor           string   `json:"cursor,omitempty"`             // 由企业微信后台返回，第一次调用可不填
	Limit            int      `json:"limit,omitempty"`              // 限制返回的条数，最多设置为1000
}

// GetScreenOperRecordResponse 获取截屏/录屏操作记录响应
type GetScreenOperRecordResponse struct {
	ErrCode    int                `json:"errcode"`
	ErrMsg     string             `json:"errmsg"`
	HasMore    bool               `json:"has_more"`    // 是否还有更多数据
	NextCursor string             `json:"next_cursor"` // 下一次调用将该值填到cursor字段
	RecordList []ScreenOperRecord `json:"record_list"` // 记录列表
}

// ScreenOperRecord 截屏/录屏操作记录
type ScreenOperRecord struct {
	Time              int64  `json:"time"`                // 操作时间
	UserID            string `json:"userid"`              // 企业用户账号id
	DepartmentID      int    `json:"department_id"`       // 企业用户部门id
	ScreenShotType    int    `json:"screen_shot_type"`    // 截屏内容的类型
	ScreenShotContent string `json:"screen_shot_content"` // 截屏内容
	System            string `json:"system"`              // 企业用户的操作系统
}
