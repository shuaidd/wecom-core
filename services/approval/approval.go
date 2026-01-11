package approval

import (
	"context"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/approval"
)

// API endpoints
const (
	setUserVacationQuotaURL = "/cgi-bin/oa/vacation/setoneuserquota"
	createTemplateURL       = "/cgi-bin/oa/approval/create_template"
	updateTemplateURL       = "/cgi-bin/oa/approval/update_template"
	getTemplateDetailURL    = "/cgi-bin/oa/gettemplatedetail"
	applyEventURL           = "/cgi-bin/oa/applyevent"
	getApprovalInfoURL      = "/cgi-bin/oa/getapprovalinfo"
	getApprovalDetailURL    = "/cgi-bin/oa/getapprovaldetail"
	getCorpVacConfigURL     = "/cgi-bin/oa/vacation/getcorpconf"
	getUserVacQuotaURL      = "/cgi-bin/oa/vacation/getuservacationquota"
	oldGetApprovalDataURL   = "/cgi-bin/corp/getapprovaldata"
	getOpenApprovalDataURL  = "/cgi-bin/corp/getopenapprovaldata"
)

// Service 审批服务
type Service struct {
	client *client.Client
}

// New 创建审批服务实例
func New(c *client.Client) *Service {
	return &Service{client: c}
}

// SetUserVacationQuota 修改成员假期余额
func (s *Service) SetUserVacationQuota(ctx context.Context, req *approval.SetUserVacationQuotaRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, setUserVacationQuotaURL, req)
	return err
}

// CreateTemplate 创建审批模板
func (s *Service) CreateTemplate(ctx context.Context, req *approval.CreateTemplateRequest) (*approval.CreateTemplateResponse, error) {
	return client.PostAndUnmarshal[approval.CreateTemplateResponse](s.client, ctx, createTemplateURL, req)
}

// UpdateTemplate 更新审批模板
func (s *Service) UpdateTemplate(ctx context.Context, req *approval.UpdateTemplateRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, updateTemplateURL, req)
	return err
}

// GetTemplateDetail 获取审批模板详情
func (s *Service) GetTemplateDetail(ctx context.Context, req *approval.GetTemplateDetailRequest) (*approval.GetTemplateDetailResponse, error) {
	return client.PostAndUnmarshal[approval.GetTemplateDetailResponse](s.client, ctx, getTemplateDetailURL, req)
}

// ApplyEvent 提交审批申请 (applyevent)
func (s *Service) ApplyEvent(ctx context.Context, req *approval.ApplyEventRequest) (*approval.ApplyEventResponse, error) {
	return client.PostAndUnmarshal[approval.ApplyEventResponse](s.client, ctx, applyEventURL, req)
}

// GetApprovalInfo 批量获取审批单号
func (s *Service) GetApprovalInfo(ctx context.Context, req *approval.GetApprovalInfoRequest) (*approval.GetApprovalInfoResponse, error) {
	return client.PostAndUnmarshal[approval.GetApprovalInfoResponse](s.client, ctx, getApprovalInfoURL, req)
}

// GetApprovalDetail 获取审批申请详情
func (s *Service) GetApprovalDetail(ctx context.Context, req *approval.GetApprovalDetailRequest) (*approval.GetApprovalDetailResponse, error) {
	return client.PostAndUnmarshal[approval.GetApprovalDetailResponse](s.client, ctx, getApprovalDetailURL, req)
}

// GetCorpVacationConfig 获取企业假期管理配置 (GET)
func (s *Service) GetCorpVacationConfig(ctx context.Context) (*approval.GetCorpVacConfigResponse, error) {
	// GET endpoint, no query params besides access_token
	return client.GetAndUnmarshal[approval.GetCorpVacConfigResponse](s.client, ctx, getCorpVacConfigURL, nil)
}

// GetUserVacationQuota 获取成员假期余额
func (s *Service) GetUserVacationQuota(ctx context.Context, req *approval.GetUserVacationQuotaRequest) (*approval.GetUserVacationQuotaResponse, error) {
	return client.PostAndUnmarshal[approval.GetUserVacationQuotaResponse](s.client, ctx, getUserVacQuotaURL, req)
}

// GetApprovalDataOld 旧接口：获取审批数据（保留兼容）
func (s *Service) GetApprovalDataOld(ctx context.Context, req *approval.GetApprovalDataOldRequest) (*approval.GetApprovalDataOldResponse, error) {
	return client.PostAndUnmarshal[approval.GetApprovalDataOldResponse](s.client, ctx, oldGetApprovalDataURL, req)
}

// GetOpenApprovalData 自建应用查询审批单当前状态 (getopenapprovaldata)
func (s *Service) GetOpenApprovalData(ctx context.Context, req *approval.GetOpenApprovalDataRequest) (*approval.GetOpenApprovalDataResponse, error) {
	return client.PostAndUnmarshal[approval.GetOpenApprovalDataResponse](s.client, ctx, getOpenApprovalDataURL, req)
}

// Helper: Build query if needed (kept for future use)
func buildQuery(params map[string]string) url.Values {
	q := url.Values{}
	for k, v := range params {
		if v != "" {
			q.Set(k, v)
		}
	}
	return q
}
