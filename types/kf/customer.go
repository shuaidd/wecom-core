package kf

import "github.com/shuaidd/wecom-core/types/common"

// BatchGetCustomerRequest 批量获取客户基础信息请求
type BatchGetCustomerRequest struct {
	ExternalUserIDList        []string `json:"external_userid_list"`        // external_userid列表，可填充个数：1~100
	NeedEnterSessionContext   int      `json:"need_enter_session_context"`  // 是否需要返回客户48小时内最后一次进入会话的上下文信息。0-不返回 1-返回。默认不返回
}

// BatchGetCustomerResponse 批量获取客户基础信息响应
type BatchGetCustomerResponse struct {
	common.Response
	CustomerList           []CustomerInfo `json:"customer_list"`            // 客户列表
	InvalidExternalUserID  []string       `json:"invalid_external_userid"`  // 无效的external_userid列表
}

// CustomerInfo 客户信息
type CustomerInfo struct {
	ExternalUserID       string                `json:"external_userid"`         // 微信客户的external_userid
	Nickname             string                `json:"nickname"`                // 微信昵称
	Avatar               string                `json:"avatar"`                  // 微信头像。第三方应用和代开发应用均不可获取
	Gender               int                   `json:"gender"`                  // 性别。第三方应用和代开发应用均不可获取，统一返回0
	UnionID              string                `json:"unionid"`                 // unionid，需要绑定微信开发者账号才能获取到
	EnterSessionContext  *EnterSessionContext  `json:"enter_session_context"`   // 48小时内最后一次进入会话的上下文信息
}

// EnterSessionContext 进入会话上下文信息
type EnterSessionContext struct {
	Scene          string          `json:"scene"`            // 进入会话的场景值
	SceneParam     string          `json:"scene_param"`      // 进入会话的自定义参数
	WechatChannels *WechatChannels `json:"wechat_channels"`  // 进入会话的视频号信息，从视频号进入会话才有值
}

// WechatChannels 视频号信息
type WechatChannels struct {
	Nickname     string `json:"nickname"`       // 视频号名称，视频号场景值为1、2、3时返回此项
	ShopNickname string `json:"shop_nickname"`  // 视频号小店名称，视频号场景值为4、5时返回此项
	Scene        int    `json:"scene"`          // 视频号场景值。1：视频号主页，2：视频号直播间商品列表页，3：视频号商品橱窗页，4：视频号小店商品详情页，5：视频号小店订单页
}

// GetUpgradeServiceConfigResponse 获取配置的专员与客户群响应
type GetUpgradeServiceConfigResponse struct {
	common.Response
	MemberRange     *MemberRange     `json:"member_range"`     // 专员服务配置范围
	GroupchatRange  *GroupchatRange  `json:"groupchat_range"`  // 客户群配置范围
}

// MemberRange 专员服务配置范围
type MemberRange struct {
	UserIDList       []string `json:"userid_list"`         // 专员userid列表
	DepartmentIDList []uint64 `json:"department_id_list"`  // 专员部门列表
}

// GroupchatRange 客户群配置范围
type GroupchatRange struct {
	ChatIDList []string `json:"chat_id_list"`  // 客户群列表
}

// UpgradeServiceRequest 为客户升级为专员或客户群服务请求
type UpgradeServiceRequest struct {
	OpenKfID       string              `json:"open_kfid"`         // 客服账号ID
	ExternalUserID string              `json:"external_userid"`   // 微信客户的external_userid
	Type           int                 `json:"type"`              // 表示是升级到专员服务还是客户群服务。1:专员服务。2:客户群服务
	Member         *UpgradeMember      `json:"member,omitempty"`  // 推荐的服务专员，type等于1时有效
	Groupchat      *UpgradeGroupchat   `json:"groupchat,omitempty"`  // 推荐的客户群，type等于2时有效
}

// UpgradeMember 推荐的服务专员
type UpgradeMember struct {
	UserID  string `json:"userid"`           // 服务专员的userid
	Wording string `json:"wording,omitempty"`  // 推荐语
}

// UpgradeGroupchat 推荐的客户群
type UpgradeGroupchat struct {
	ChatID  string `json:"chat_id"`          // 客户群id
	Wording string `json:"wording,omitempty"`  // 推荐语
}

// CancelUpgradeServiceRequest 为客户取消推荐请求
type CancelUpgradeServiceRequest struct {
	OpenKfID       string `json:"open_kfid"`        // 客服账号ID
	ExternalUserID string `json:"external_userid"`  // 微信客户的external_userid
}
