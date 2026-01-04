package wecom

import (
	"context"

	"github.com/shuaidd/wecom-core/config"
	"github.com/shuaidd/wecom-core/internal/auth"
	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/internal/retry"
	"github.com/shuaidd/wecom-core/services/agent"
	"github.com/shuaidd/wecom-core/services/contact"
	"github.com/shuaidd/wecom-core/services/corpgroup"
	"github.com/shuaidd/wecom-core/services/externalcontact"
	"github.com/shuaidd/wecom-core/services/invoice"
	"github.com/shuaidd/wecom-core/services/ip"
	"github.com/shuaidd/wecom-core/services/kf"
	"github.com/shuaidd/wecom-core/services/media"
	"github.com/shuaidd/wecom-core/services/message"
	"github.com/shuaidd/wecom-core/services/oauth"
	"github.com/shuaidd/wecom-core/services/qrcode"
	"github.com/shuaidd/wecom-core/services/security"
	"github.com/shuaidd/wecom-core/services/updown"
)

// Client 企业微信SDK客户端
type Client struct {
	// Agent 应用管理服务
	Agent *agent.Service
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
	// Invoice 电子发票服务
	Invoice *invoice.Service
	// KF 微信客服服务
	KF *kf.Service

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

	// 4. 创建重试策略
	retryPolicy := retry.NewPolicy(
		cfg.MaxRetries,
		cfg.InitialBackoff,
		cfg.MaxBackoff,
	)
	retryExecutor := retry.NewExecutor(retryPolicy, cfg.Logger)

	// 5. 创建 HTTP 客户端
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

	// 6. 创建服务客户端
	c := &Client{
		config:          cfg,
		tokenManager:    tokenManager,
		httpClient:      httpClient,
		Agent:           agent.NewService(httpClient),
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
		Invoice:         invoice.NewService(httpClient),
		KF:              kf.NewService(httpClient),
	}

	return c, nil
}

// WithTraceID 将 TraceId 添加到 context
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return client.WithTraceID(ctx, traceID)
}
