package oauth

import (
	"context"
	"fmt"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/oauth"
)

const (
	// OAuth2AuthorizeURL OAuth2授权URL
	oauth2AuthorizeURL = "https://open.weixin.qq.com/connect/oauth2/authorize"
)

// Service OAuth服务
type Service struct {
	client *client.Client
}

// NewService 创建OAuth服务
func NewService(c *client.Client) *Service {
	return &Service{
		client: c,
	}
}

// BuildAuthorizeURL 构造网页授权链接
// 用于在打开的网页里面携带用户的身份信息
func (s *Service) BuildAuthorizeURL(params oauth.BuildAuthorizeURLParams) (string, error) {
	// 验证必填参数
	if params.CorpID == "" {
		return "", fmt.Errorf("CorpID is required")
	}
	if params.RedirectURI == "" {
		return "", fmt.Errorf("RedirectURI is required")
	}
	if params.Scope == "" {
		return "", fmt.Errorf("Scope is required")
	}

	// snsapi_privateinfo时必须填写agentid
	if params.Scope == oauth.ScopePrivateInfo && params.AgentID == "" {
		return "", fmt.Errorf("AgentID is required when Scope is snsapi_privateinfo")
	}

	// 构造URL参数
	query := url.Values{}
	query.Set("appid", params.CorpID)
	query.Set("redirect_uri", params.RedirectURI)
	query.Set("response_type", "code")
	query.Set("scope", string(params.Scope))
	if params.State != "" {
		query.Set("state", params.State)
	}
	if params.AgentID != "" {
		query.Set("agentid", params.AgentID)
	}

	// 拼接URL
	authURL := fmt.Sprintf("%s?%s#wechat_redirect", oauth2AuthorizeURL, query.Encode())
	return authURL, nil
}

// GetUserInfo 获取访问用户身份
// 根据code获取成员信息
func (s *Service) GetUserInfo(ctx context.Context, code string) (*oauth.GetUserInfoResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("code is required")
	}

	query := url.Values{}
	query.Set("code", code)

	return client.GetAndUnmarshal[oauth.GetUserInfoResponse](
		s.client,
		ctx,
		"/cgi-bin/auth/getuserinfo",
		query,
	)
}

// GetUserDetail 获取访问用户敏感信息
// 通过user_ticket获取成员授权的敏感字段
func (s *Service) GetUserDetail(ctx context.Context, userTicket string) (*oauth.GetUserDetailResponse, error) {
	if userTicket == "" {
		return nil, fmt.Errorf("userTicket is required")
	}

	req := &oauth.GetUserDetailRequest{
		UserTicket: userTicket,
	}

	return client.PostAndUnmarshal[oauth.GetUserDetailResponse](
		s.client,
		ctx,
		"/cgi-bin/auth/getuserdetail",
		req,
	)
}

// GetTFAInfo 获取用户二次验证信息
// 用于获取触发二次验证的成员userid和tfa_code
func (s *Service) GetTFAInfo(ctx context.Context, code string) (*oauth.GetTFAInfoResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("code is required")
	}

	req := &oauth.GetTFAInfoRequest{
		Code: code,
	}

	return client.PostAndUnmarshal[oauth.GetTFAInfoResponse](
		s.client,
		ctx,
		"/cgi-bin/auth/get_tfa_info",
		req,
	)
}
