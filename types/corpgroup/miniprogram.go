package corpgroup

import "github.com/shuaidd/wecom-core/types/common"

// TransferSessionRequest 获取下级/下游企业小程序session请求
type TransferSessionRequest struct {
	UserID     string `json:"userid"`      // 通过code2Session接口获取到的加密的userid
	SessionKey string `json:"session_key"` // 通过code2Session接口获取到的属于上级/上游企业的会话密钥
}

// TransferSessionResponse 获取下级/下游企业小程序session响应
type TransferSessionResponse struct {
	common.Response
	UserID     string `json:"userid"`      // 下级/下游企业用户的ID
	SessionKey string `json:"session_key"` // 属于下级/下游企业的会话密钥
}
