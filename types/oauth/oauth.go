package oauth

import "github.com/shuaidd/wecom-core/types/common"

// AuthScope OAuth2授权作用域
type AuthScope string

const (
	// ScopeBase 静默授权，可获取成员的基础信息（UserId）
	ScopeBase AuthScope = "snsapi_base"
	// ScopePrivateInfo 手动授权，可获取成员的详细信息，包含头像、二维码等敏感信息
	ScopePrivateInfo AuthScope = "snsapi_privateinfo"
)

// BuildAuthorizeURLParams 构造网页授权链接参数
type BuildAuthorizeURLParams struct {
	// CorpID 企业的CorpID
	CorpID string
	// RedirectURI 授权后重定向的回调链接地址（需要urlencode处理）
	RedirectURI string
	// Scope 应用授权作用域
	Scope AuthScope
	// State 重定向后会带上state参数，企业可以填写a-zA-Z0-9的参数值，长度不可超过128个字节
	State string
	// AgentID 应用agentid（snsapi_privateinfo时必填）
	AgentID string
}

// GetUserInfoResponse 获取访问用户身份响应
type GetUserInfoResponse struct {
	common.Response
	// UserID 成员UserID（企业成员时返回）
	UserID string `json:"userid,omitempty"`
	// UserTicket 成员票据，最大为512字节，有效期为1800s（scope为snsapi_privateinfo时返回）
	UserTicket string `json:"user_ticket,omitempty"`
	// OpenID 非企业成员的标识，对当前企业唯一
	OpenID string `json:"openid,omitempty"`
	// ExternalUserID 外部联系人id
	ExternalUserID string `json:"external_userid,omitempty"`
}

// GetUserDetailRequest 获取访问用户敏感信息请求
type GetUserDetailRequest struct {
	// UserTicket 成员票据
	UserTicket string `json:"user_ticket"`
}

// GetUserDetailResponse 获取访问用户敏感信息响应
type GetUserDetailResponse struct {
	common.Response
	// UserID 成员UserID
	UserID string `json:"userid"`
	// Gender 性别。0表示未定义，1表示男性，2表示女性
	Gender string `json:"gender,omitempty"`
	// Avatar 头像url
	Avatar string `json:"avatar,omitempty"`
	// QRCode 员工个人二维码
	QRCode string `json:"qr_code,omitempty"`
	// Mobile 手机（第三方应用不可获取）
	Mobile string `json:"mobile,omitempty"`
	// Email 邮箱（第三方应用不可获取）
	Email string `json:"email,omitempty"`
	// BizMail 企业邮箱（第三方应用不可获取）
	BizMail string `json:"biz_mail,omitempty"`
	// Address 地址（第三方应用不可获取）
	Address string `json:"address,omitempty"`
}

// GetTFAInfoRequest 获取用户二次验证信息请求
type GetTFAInfoRequest struct {
	// Code 用户进入二次验证页面时，企业微信颁发的code
	Code string `json:"code"`
}

// GetTFAInfoResponse 获取用户二次验证信息响应
type GetTFAInfoResponse struct {
	common.Response
	// UserID 成员UserID
	UserID string `json:"userid"`
	// TFACode 二次验证授权码，有效期五分钟，且只能使用一次
	TFACode string `json:"tfa_code"`
}
