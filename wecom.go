package wecom

import (
	"context"

	"github.com/shuaidd/wecom-core/config"
	"github.com/shuaidd/wecom-core/internal/auth"
	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/internal/retry"
	"github.com/shuaidd/wecom-core/pkg/interceptor"
	"github.com/shuaidd/wecom-core/services/agent"
	"github.com/shuaidd/wecom-core/services/calendar"
	"github.com/shuaidd/wecom-core/services/checkin"
	"github.com/shuaidd/wecom-core/services/contact"
	"github.com/shuaidd/wecom-core/services/corpgroup"
	"github.com/shuaidd/wecom-core/services/email"
	"github.com/shuaidd/wecom-core/services/externalcontact"
	"github.com/shuaidd/wecom-core/services/invoice"
	"github.com/shuaidd/wecom-core/services/ip"
	"github.com/shuaidd/wecom-core/services/kf"
	"github.com/shuaidd/wecom-core/services/media"
	"github.com/shuaidd/wecom-core/services/meeting"
	"github.com/shuaidd/wecom-core/services/message"
	"github.com/shuaidd/wecom-core/services/oauth"
	"github.com/shuaidd/wecom-core/services/qrcode"
	"github.com/shuaidd/wecom-core/services/reserve_meeting"
	"github.com/shuaidd/wecom-core/services/security"
	"github.com/shuaidd/wecom-core/services/updown"
	"github.com/shuaidd/wecom-core/services/webinar"
	"github.com/shuaidd/wecom-core/services/wedoc"
	"github.com/shuaidd/wecom-core/services/wedrive"
	"github.com/shuaidd/wecom-core/services/approval"
)

// 暴露给用户的类型别名，用于自定义 API 调用

// AgentConfig 应用配置（用于配置多应用）
type AgentConfig = config.AgentConfig

// Response 响应对象，包含原始响应数据
type Response = client.Response

// CommonResponse 企业微信API通用响应字段
// 用户自定义响应类型时应该嵌入此类型
type CommonResponse = client.CommonResponse

// RequestInterceptor 请求拦截器
type RequestInterceptor = interceptor.RequestInterceptor

// ResponseInterceptor 响应拦截器（解析前）
type ResponseInterceptor = interceptor.ResponseInterceptor

// AfterResponseInterceptor 响应后拦截器（解析后）
type AfterResponseInterceptor = interceptor.AfterResponseInterceptor

// InterceptorResponse 拦截器中使用的响应对象
type InterceptorResponse = interceptor.Response

// Client 企业微信SDK客户端
type Client struct {
	// Agent 应用管理服务
	Agent *agent.Service
	// Calendar 日历服务
	Calendar *calendar.Service
	// Contact 通讯录服务
	Contact *contact.Service
	// IP IP相关服务
	IP *ip.Service
	// QRCode 企业二维码服务
	QRCode *qrcode.Service
	// OAuth 身份验证服务
	OAuth *oauth.Service
	// UpDown 上下游服务
	UpDown *updown.Service
	// CorpGroup 企业互联服务
	CorpGroup *corpgroup.Service
	// Security 安全管理服务
	Security *security.Service
	// Message 消息服务
	Message *message.Service
	// ExternalContact 外部联系人服务
	ExternalContact *externalcontact.Service
	// Media 素材管理服务
	Media *media.Service
	// Checkin 打卡服务
	Checkin *checkin.Service
	// Invoice 电子发票服务
	Invoice *invoice.Service
	// KF 微信客服服务
	KF *kf.Service
	// Email 邮件服务
	Email *email.Service
	// Wedoc 微文档服务
	Wedoc *wedoc.Service
	// Wedrive 微盘服务
	Wedrive *wedrive.Service
	// Meeting 会议服务
	Meeting *meeting.Service
	// ReserveMeeting 预约会议高级管理服务
	ReserveMeeting *reserve_meeting.Service
	// Webinar 网络研讨会服务
	Webinar *webinar.Service
	// Approval 审批服务
	Approval *approval.Service

	// 内部组件(不对外暴露)
	config       *config.Config
	tokenManager *auth.TokenManager
	httpClient   *client.Client
}

// New 创建企业微信SDK客户端
func New(opts ...config.Option) (*Client, error) {
	// 1. 创建配置
	cfg := config.New(opts...)

	// 2. 验证配置
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	// 3. 创建 Token 管理器
	tokenManager := auth.NewTokenManager(
		cfg.CorpID,
		cfg.CorpSecret,
		cfg.BaseURL,
		cfg.Cache,
		cfg.Logger,
	)

	// 4. 注册多应用配置到 TokenManager
	if len(cfg.Agents) > 0 {
		for key, agentCfg := range cfg.Agents {
			tokenManager.RegisterAgent(key, agentCfg.AgentID, agentCfg.Secret)
		}
	}

	// 5. 创建重试策略
	retryPolicy := retry.NewPolicy(
		cfg.MaxRetries,
		cfg.InitialBackoff,
		cfg.MaxBackoff,
	)
	retryExecutor := retry.NewExecutor(retryPolicy, cfg.Logger)

	// 6. 创建 HTTP 客户端
	httpClient := client.New(
		cfg.BaseURL,
		cfg.Timeout,
		cfg.Logger,
		tokenManager,
		retryExecutor,
	)

	if cfg.Debug {
		httpClient.SetDebug(true)
	}

	// 7. 注册拦截器
	for _, interceptor := range cfg.RequestInterceptors {
		httpClient.AddRequestInterceptor(interceptor)
	}
	for _, interceptor := range cfg.ResponseInterceptors {
		httpClient.AddResponseInterceptor(interceptor)
	}
	for _, interceptor := range cfg.AfterResponseInterceptors {
		httpClient.AddAfterResponseInterceptor(interceptor)
	}

	// 8. 创建服务客户端
	c := &Client{
		config:          cfg,
		tokenManager:    tokenManager,
		httpClient:      httpClient,
		Agent:           agent.NewService(httpClient),
		Calendar:        calendar.NewService(httpClient),
		Contact:         contact.NewService(httpClient),
		IP:              ip.NewService(httpClient),
		QRCode:          qrcode.NewService(httpClient),
		OAuth:           oauth.NewService(httpClient),
		UpDown:          updown.NewService(httpClient),
		CorpGroup:       corpgroup.NewService(httpClient),
		Security:        security.NewService(httpClient),
		Message:         message.NewService(httpClient),
		ExternalContact: externalcontact.NewService(httpClient),
		Media:           media.NewService(httpClient),
		Checkin:         checkin.NewService(httpClient),
		Invoice:         invoice.NewService(httpClient),
		KF:              kf.NewService(httpClient),
		Email:           email.NewService(httpClient),
		Wedoc:           wedoc.New(httpClient),
		Wedrive:         wedrive.New(httpClient),
		Meeting:         meeting.NewService(httpClient),
		ReserveMeeting:  reserve_meeting.NewService(httpClient),
		Webinar:         webinar.NewService(httpClient),
		Approval:        approval.New(httpClient),
	}

	return c, nil
}

// WithTraceID 将 TraceId 添加到 context
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return client.WithTraceID(ctx, traceID)
}

// WithAgentName 将应用名称添加到 context
// 使用此函数可以在调用API时指定使用哪个应用的凭证
func WithAgentName(ctx context.Context, agentName string) context.Context {
	return client.WithAgentName(ctx, agentName)
}

// WithAgentID 将应用ID添加到 context
// 使用此函数可以在调用API时指定使用哪个应用的凭证
func WithAgentID(ctx context.Context, agentID int64) context.Context {
	return client.WithAgentID(ctx, agentID)
}

// CustomGet 发送自定义 GET 请求
// 自动处理 access_token 注入和重试逻辑
//
// 示例：
//
//	resp, err := client.CustomGet(ctx, "/cgi-bin/custom/api", url.Values{"param": []string{"value"}})
//	if err != nil {
//	    return err
//	}
//	// 手动解析响应
//	var result YourCustomType
//	if err := resp.Unmarshal(&result); err != nil {
//	    return err
//	}
func (c *Client) CustomGet(ctx context.Context, path string, query map[string]string) (*client.Response, error) {
	q := make(map[string][]string)
	for k, v := range query {
		q[k] = []string{v}
	}
	return c.httpClient.Get(ctx, path, q)
}

// CustomPost 发送自定义 POST 请求
// 自动处理 access_token 注入和重试逻辑
//
// 示例：
//
//	req := YourCustomRequest{Field: "value"}
//	resp, err := client.CustomPost(ctx, "/cgi-bin/custom/api", req)
//	if err != nil {
//	    return err
//	}
//	// 手动解析响应
//	var result YourCustomType
//	if err := resp.Unmarshal(&result); err != nil {
//	    return err
//	}
func (c *Client) CustomPost(ctx context.Context, path string, body any) (*client.Response, error) {
	return c.httpClient.Post(ctx, path, body)
}

// CustomGetAndUnmarshal 发送自定义 GET 请求并自动解析响应
// 自动处理 access_token 注入、重试逻辑和响应解析
//
// 示例：
//
//	type CustomResponse struct {
//	    client.CommonResponse
//	    Data string `json:"data"`
//	}
//	result, err := client.CustomGetAndUnmarshal[CustomResponse](ctx, "/cgi-bin/custom/api", map[string]string{"param": "value"})
//	if err != nil {
//	    return err
//	}
//	fmt.Println(result.Data)
func CustomGetAndUnmarshal[T any](c *Client, ctx context.Context, path string, query map[string]string) (*T, error) {
	q := make(map[string][]string)
	for k, v := range query {
		q[k] = []string{v}
	}
	return client.GetAndUnmarshal[T](c.httpClient, ctx, path, q)
}

// CustomPostAndUnmarshal 发送自定义 POST 请求并自动解析响应
// 自动处理 access_token 注入、重试逻辑和响应解析
//
// 示例：
//
//	type CustomRequest struct {
//	    Field string `json:"field"`
//	}
//	type CustomResponse struct {
//	    client.CommonResponse
//	    Result string `json:"result"`
//	}
//	req := CustomRequest{Field: "value"}
//	result, err := client.CustomPostAndUnmarshal[CustomResponse](c, ctx, "/cgi-bin/custom/api", req)
//	if err != nil {
//	    return err
//	}
//	fmt.Println(result.Result)
func CustomPostAndUnmarshal[T any](c *Client, ctx context.Context, path string, body any) (*T, error) {
	return client.PostAndUnmarshal[T](c.httpClient, ctx, path, body)
}
