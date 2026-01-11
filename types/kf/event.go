package kf

import "github.com/shuaidd/wecom-core/types/common"

// SendMsgOnEventRequest 发送欢迎语等事件响应消息请求
// 当特定的事件回调消息包含code字段,或通过接口变更到特定的会话状态,会返回code字段
// 开发者可以此code为凭证,调用该接口给用户发送相应事件场景下的消息
type SendMsgOnEventRequest struct {
	Code    string `json:"code"`            // 事件响应消息对应的code,通过事件回调下发,仅可使用一次
	MsgID   string `json:"msgid,omitempty"` // 消息ID,不多于32字节,字符串取值范围:[0-9a-zA-Z_-]*
	MsgType string `json:"msgtype"`         // 消息类型:text-文本消息,msgmenu-菜单消息

	// 消息内容,根据msgtype选择填充对应的字段
	Text    *TextContent    `json:"text,omitempty"`    // 文本消息
	MsgMenu *MsgMenuContent `json:"msgmenu,omitempty"` // 菜单消息
}

// SendMsgOnEventResponse 发送欢迎语等事件响应消息响应
type SendMsgOnEventResponse struct {
	common.Response
	MsgID string `json:"msgid"` // 消息ID
}
