package kf

import "github.com/shuaidd/wecom-core/types/common"

// GetServiceStateRequest 获取会话状态请求
type GetServiceStateRequest struct {
	OpenKfID       string `json:"open_kfid"`       // 客服账号ID
	ExternalUserID string `json:"external_userid"` // 微信客户的external_userid
}

// GetServiceStateResponse 获取会话状态响应
type GetServiceStateResponse struct {
	common.Response
	ServiceState   int    `json:"service_state"`             // 当前的会话状态,0-未处理,1-由智能助手接待,2-待接入池排队中,3-由人工接待,4-已结束/未开始
	ServicerUserID string `json:"servicer_userid,omitempty"` // 接待人员的userid,仅当state=3时有效
}

// TransServiceStateRequest 变更会话状态请求
type TransServiceStateRequest struct {
	OpenKfID       string `json:"open_kfid"`                 // 客服账号ID
	ExternalUserID string `json:"external_userid"`           // 微信客户的external_userid
	ServiceState   int    `json:"service_state"`             // 变更的目标状态,状态定义参考GetServiceStateResponse
	ServicerUserID string `json:"servicer_userid,omitempty"` // 接待人员的userid,当state=3时要求必填
}

// TransServiceStateResponse 变更会话状态响应
type TransServiceStateResponse struct {
	common.Response
	MsgCode string `json:"msg_code,omitempty"` // 用于发送响应事件消息的code
}
