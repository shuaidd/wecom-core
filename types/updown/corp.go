package updown

import "github.com/shuaidd/wecom-core/types/common"

// Chain 上下游信息
type Chain struct {
	ChainID   string `json:"chain_id"`   // 上下游id
	ChainName string `json:"chain_name"` // 上下游名称
}

// GetChainListResponse 获取上下游列表响应
type GetChainListResponse struct {
	common.Response
	Chains []Chain `json:"chains"` // 上下游列表
}

// Group 分组信息
type Group struct {
	GroupID   int64  `json:"groupid"`    // 分组id
	GroupName string `json:"group_name"` // 分组名称
	ParentID  int64  `json:"parentid"`   // 父分组id，根分组id为1
	Order     int    `json:"order"`      // 父部门中的次序值
}

// GetChainGroupRequest 获取上下游通讯录分组请求
type GetChainGroupRequest struct {
	ChainID string `json:"chain_id"`          // 上下游id
	GroupID *int64 `json:"groupid,omitempty"` // 分组id，可选
}

// GetChainGroupResponse 获取上下游通讯录分组响应
type GetChainGroupResponse struct {
	common.Response
	Groups []Group `json:"groups"` // 分组列表
}

// GroupCorp 分组下的企业信息
type GroupCorp struct {
	GroupID       int64  `json:"groupid"`                  // 企业所属上下游的分组id
	CorpID        string `json:"corpid,omitempty"`         // 企业id，已加入的企业返回
	CorpName      string `json:"corp_name"`                // 企业名称
	CustomID      string `json:"custom_id"`                // 上下游企业自定义id
	InviteUserID  string `json:"invite_userid,omitempty"`  // 邀请人的userid
	PendingCorpID string `json:"pending_corpid,omitempty"` // 未加入企业id
	IsJoined      int    `json:"is_joined"`                // 企业是否已加入
}

// GetChainCorpInfoListRequest 获取企业上下游通讯录分组下的企业详情列表请求
type GetChainCorpInfoListRequest struct {
	ChainID     string `json:"chain_id"`               // 上下游id
	GroupID     *int64 `json:"groupid,omitempty"`      // 分组id，可选
	NeedPending bool   `json:"need_pending,omitempty"` // 是否需要返回未加入的企业
	Cursor      string `json:"cursor,omitempty"`       // 分页游标
	Limit       int    `json:"limit,omitempty"`        // 分页大小
}

// GetChainCorpInfoListResponse 获取企业上下游通讯录分组下的企业详情列表响应
type GetChainCorpInfoListResponse struct {
	common.Response
	HasMore    bool        `json:"has_more"`    // 是否还有更多记录
	NextCursor string      `json:"next_cursor"` // 下次请求时应传入的cursor
	GroupCorps []GroupCorp `json:"group_corps"` // 分组列表数据
}

// GetChainCorpInfoRequest 获取企业上下游通讯录下的企业信息请求
type GetChainCorpInfoRequest struct {
	ChainID       string `json:"chain_id"`                 // 上下游id
	CorpID        string `json:"corpid,omitempty"`         // 已加入企业id
	PendingCorpID string `json:"pending_corpid,omitempty"` // 待加入企业id
}

// GetChainCorpInfoResponse 获取企业上下游通讯录下的企业信息响应
type GetChainCorpInfoResponse struct {
	common.Response
	CorpName            string `json:"corp_name"`            // 企业名称
	QualificationStatus int    `json:"qualification_status"` // 企业验证状态，1未验证，2已验证，3已认证
	CustomID            string `json:"custom_id"`            // 上下游企业自定义id
	GroupID             int64  `json:"groupid"`              // 企业所属上下游的分组id
	IsJoined            bool   `json:"is_joined"`            // 企业是否已加入
}

// RemoveCorpRequest 移除企业请求
type RemoveCorpRequest struct {
	ChainID       string `json:"chain_id"`                 // 上下游id
	CorpID        string `json:"corpid,omitempty"`         // 需要移除的下游企业corpid
	PendingCorpID string `json:"pending_corpid,omitempty"` // 需要移除的未加入下游企业corpid
}

// RemoveCorpResponse 移除企业响应
type RemoveCorpResponse struct {
	common.Response
}

// GetChainUserCustomIDRequest 查询成员自定义id请求
type GetChainUserCustomIDRequest struct {
	ChainID string `json:"chain_id"` // 上下游id
	CorpID  string `json:"corpid"`   // 已加入企业id
	UserID  string `json:"userid"`   // 企业内的成员
}

// GetChainUserCustomIDResponse 查询成员自定义id响应
type GetChainUserCustomIDResponse struct {
	common.Response
	UserCustomID string `json:"user_custom_id"` // 成员自定义id
}

// GetCorpSharedChainListRequest 获取下级企业加入的上下游请求
type GetCorpSharedChainListRequest struct {
	CorpID string `json:"corpid,omitempty"` // 已加入企业id
}

// GetCorpSharedChainListResponse 获取下级企业加入的上下游响应
type GetCorpSharedChainListResponse struct {
	common.Response
	Chains []Chain `json:"chains"` // 上下游列表
}
