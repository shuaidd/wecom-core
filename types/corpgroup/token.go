package corpgroup

import "github.com/shuaidd/wecom-core/types/common"

// GetTokenRequest 获取下级/下游企业的access_token请求
type GetTokenRequest struct {
	CorpID       string `json:"corpid"`                  // 已授权的下级/下游企业corpid
	AgentID      int64  `json:"agentid"`                 // 已授权的下级/下游企业应用ID
	BusinessType *int   `json:"business_type,omitempty"` // 填0则为企业互联/局校互联，填1则表示上下游企业，默认0
}

// GetTokenResponse 获取下级/下游企业的access_token响应
type GetTokenResponse struct {
	common.Response
	AccessToken string `json:"access_token"` // 获取到的下级/下游企业调用凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证的有效时间（秒）
}
