package updown

import "github.com/shuaidd/wecom-core/types/common"

// CorpShareInfo 应用共享信息
type CorpShareInfo struct {
	CorpID   string `json:"corpid"`    // 下级/下游企业corpid
	CorpName string `json:"corp_name"` // 下级/下游企业名称
	AgentID  int64  `json:"agentid"`   // 下级/下游企业应用id
}

// ListAppShareInfoRequest 获取应用共享信息请求
type ListAppShareInfoRequest struct {
	AgentID      int64  `json:"agentid"`                // 上级/上游企业应用agentid
	BusinessType *int   `json:"business_type,omitempty"` // 填0则为企业互联/局校互联，填1则表示上下游企业
	CorpID       string `json:"corpid,omitempty"`       // 下级/下游企业corpid
	Limit        int    `json:"limit,omitempty"`        // 返回的最大记录数，最大值100
	Cursor       string `json:"cursor,omitempty"`       // 用于分页查询的游标
}

// ListAppShareInfoResponse 获取应用共享信息响应
type ListAppShareInfoResponse struct {
	common.Response
	Ending     int             `json:"ending"`      // 1表示拉取完毕，0表示数据没有拉取完
	NextCursor string          `json:"next_cursor"` // 分页游标
	CorpList   []CorpShareInfo `json:"corp_list"`   // 应用共享信息
}
