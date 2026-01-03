package security

// ListMemberOperLogRequest 获取成员操作记录请求
type ListMemberOperLogRequest struct {
	StartTime int64  `json:"start_time"`          // 开始时间
	EndTime   int64  `json:"end_time"`            // 结束时间，范围不能超过7天
	OperType  int    `json:"oper_type,omitempty"` // 操作类型，不填表示全部
	UserID    string `json:"userid,omitempty"`    // 操作者userid过滤
	Cursor    string `json:"cursor,omitempty"`    // 分页游标
	Limit     int    `json:"limit,omitempty"`     // 最大记录数，默认400，最多400
}

// ListMemberOperLogResponse 获取成员操作记录响应
type ListMemberOperLogResponse struct {
	ErrCode    int                `json:"errcode"`
	ErrMsg     string             `json:"errmsg"`
	HasMore    bool               `json:"has_more"`    // 是否还有下一页
	NextCursor string             `json:"next_cursor"` // 下一页的分页游标
	RecordList []MemberOperRecord `json:"record_list"` // 记录列表
}

// MemberOperRecord 成员操作记录
type MemberOperRecord struct {
	Time       int64  `json:"time"`        // 操作时间
	UserID     string `json:"userid"`      // 操作者userid
	OperType   int    `json:"oper_type"`   // 操作类型
	DetailInfo string `json:"detail_info"` // 相关数据
	IP         string `json:"ip"`          // 操作者ip
}
