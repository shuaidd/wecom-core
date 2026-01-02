package qrcode

import "github.com/shuaidd/wecom-core/types/common"

// GetJoinQRCodeResponse 获取加入企业二维码响应
type GetJoinQRCodeResponse struct {
	common.Response
	// JoinQRCode 二维码链接，有效期7天
	JoinQRCode string `json:"join_qrcode"`
}

// BatchInviteRequest 邀请成员请求
type BatchInviteRequest struct {
	// User 成员ID列表，最多支持1000个
	User []string `json:"user,omitempty"`
	// Party 部门ID列表，最多支持100个
	Party []int `json:"party,omitempty"`
	// Tag 标签ID列表，最多支持100个
	Tag []int `json:"tag,omitempty"`
}

// BatchInviteResponse 邀请成员响应
type BatchInviteResponse struct {
	common.Response
	// InvalidUser 非法成员列表
	InvalidUser []string `json:"invaliduser,omitempty"`
	// InvalidParty 非法部门列表
	InvalidParty []int `json:"invalidparty,omitempty"`
	// InvalidTag 非法标签列表
	InvalidTag []int `json:"invalidtag,omitempty"`
}
