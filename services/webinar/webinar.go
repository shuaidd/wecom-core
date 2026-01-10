package webinar

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/webinar"
)

type Service struct {
	client *client.Client
}

func NewService(c *client.Client) *Service {
	return &Service{client: c}
}

func (s *Service) Update(ctx context.Context, req *webinar.UpdateWebinarRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/meeting/webinar/update", req)
	return err
}

func (s *Service) SetConfig(ctx context.Context, req *webinar.SetWebinarConfigRequest) (*webinar.SetWebinarConfigResponse, error) {
	return client.PostAndUnmarshal[webinar.SetWebinarConfigResponse](s.client, ctx, "/cgi-bin/meeting/webinar/enroll/set_config", req)
}

func (s *Service) DeleteEnroll(ctx context.Context, req *webinar.DeleteWebinarEnrollRequest) (*webinar.DeleteWebinarEnrollResponse, error) {
	return client.PostAndUnmarshal[webinar.DeleteWebinarEnrollResponse](s.client, ctx, "/cgi-bin/meeting/webinar/enroll/delete", req)
}

func (s *Service) Cancel(ctx context.Context, meetingID string) error {
	req := &webinar.CancelWebinarRequest{
		MeetingID: meetingID,
	}
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/meeting/webinar/cancel", req)
	return err
}

func (s *Service) ApproveEnroll(ctx context.Context, req *webinar.ApproveWebinarEnrollRequest) (*webinar.ApproveWebinarEnrollResponse, error) {
	return client.PostAndUnmarshal[webinar.ApproveWebinarEnrollResponse](s.client, ctx, "/cgi-bin/meeting/webinar/enroll/approve", req)
}

func (s *Service) UpdateWarmUp(ctx context.Context, req *webinar.UpdateWebinarWarmUpRequest) error {
	_, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/cgi-bin/meeting/webinar/update_warm_up", req)
	return err
}

func (s *Service) ListGuest(ctx context.Context, req *webinar.ListWebinarGuestRequest) (*webinar.ListWebinarGuestResponse, error) {
	return client.PostAndUnmarshal[webinar.ListWebinarGuestResponse](s.client, ctx, "/cgi-bin/meeting/webinar/list_guest", req)
}

func (s *Service) QueryEnrollByTmpOpenID(ctx context.Context, req *webinar.QueryWebinarEnrollByTmpOpenIDRequest) (*webinar.QueryWebinarEnrollByTmpOpenIDResponse, error) {
	return client.PostAndUnmarshal[webinar.QueryWebinarEnrollByTmpOpenIDResponse](s.client, ctx, "/cgi-bin/meeting/webinar/enroll/query_by_tmp_openid", req)
}

func (s *Service) ListEnroll(ctx context.Context, req *webinar.ListWebinarEnrollRequest) (*webinar.ListWebinarEnrollResponse, error) {
	return client.PostAndUnmarshal[webinar.ListWebinarEnrollResponse](s.client, ctx, "/cgi-bin/meeting/webinar/enroll/list", req)
}

func (s *Service) GetConfig(ctx context.Context, meetingID string) (*webinar.GetWebinarConfigResponse, error) {
	req := &webinar.GetWebinarConfigRequest{
		MeetingID: meetingID,
	}
	return client.PostAndUnmarshal[webinar.GetWebinarConfigResponse](s.client, ctx, "/cgi-bin/meeting/webinar/enroll/get_config", req)
}

func (s *Service) GetInfo(ctx context.Context, req *webinar.GetWebinarRequest) (*webinar.GetWebinarResponse, error) {
	return client.PostAndUnmarshal[webinar.GetWebinarResponse](s.client, ctx, "/cgi-bin/meeting/webinar/get", req)
}
