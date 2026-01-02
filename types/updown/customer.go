package updown

import "github.com/shuaidd/wecom-core/types/common"

// ExternalUserIDInfo 外部联系人信息
type ExternalUserIDInfo struct {
	CorpID         string `json:"corpid"`          // 所属企业id
	ExternalUserID string `json:"external_userid"` // 外部联系人id
}

// UnionIDToExternalUserIDRequest 通过unionid和openid查询external_userid请求
type UnionIDToExternalUserIDRequest struct {
	UnionID        string `json:"unionid"`                  // 微信客户的unionid
	OpenID         string `json:"openid"`                   // 微信客户的openid
	CorpID         string `json:"corpid,omitempty"`         // 需要换取的企业corpid
	MassCallTicket string `json:"mass_call_ticket,omitempty"` // 大批量调用凭据
}

// UnionIDToExternalUserIDResponse 通过unionid和openid查询external_userid响应
type UnionIDToExternalUserIDResponse struct {
	common.Response
	ExternalUserIDInfo []ExternalUserIDInfo `json:"external_userid_info"` // 外部联系人信息列表
}

// UnionIDToPendingIDRequest unionid查询pending_id请求
type UnionIDToPendingIDRequest struct {
	UnionID string `json:"unionid"` // 微信客户的unionid
	OpenID  string `json:"openid"`  // 微信客户的openid
}

// UnionIDToPendingIDResponse unionid查询pending_id响应
type UnionIDToPendingIDResponse struct {
	common.Response
	PendingID string `json:"pending_id"` // unionid和openid对应的pending_id
}

// PendingIDResult pending_id转换结果
type PendingIDResult struct {
	ExternalUserID string `json:"external_userid"` // 转换的external_userid
	PendingID      string `json:"pending_id"`      // 临时外部联系人ID
}

// ExternalUserIDToPendingIDRequest external_userid查询pending_id请求
type ExternalUserIDToPendingIDRequest struct {
	ChatID         string   `json:"chat_id,omitempty"`   // 群id
	ExternalUserID []string `json:"external_userid"`     // 上游或下游企业外部联系人id，最多同时查询100个
}

// ExternalUserIDToPendingIDResponse external_userid查询pending_id响应
type ExternalUserIDToPendingIDResponse struct {
	common.Response
	Result []PendingIDResult `json:"result"` // 转换结果
}
