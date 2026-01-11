package meeting

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/common"
	"github.com/shuaidd/wecom-core/types/meeting"
)

// Service 会议服务
type Service struct {
	client *client.Client
}

// NewService 创建会议服务实例
func NewService(c *client.Client) *Service {
	return &Service{client: c}
}

// Create 创建预约会议
// 文档: https://developer.work.weixin.qq.com/document/path/93624
func (s *Service) Create(ctx context.Context, req *meeting.CreateMeetingRequest) (*meeting.CreateMeetingResponse, error) {
	return client.PostAndUnmarshal[meeting.CreateMeetingResponse](s.client, ctx, "/cgi-bin/meeting/create", req)
}

// Update 修改预约会议
// 文档: https://developer.work.weixin.qq.com/document/path/93625
func (s *Service) Update(ctx context.Context, req *meeting.UpdateMeetingRequest) error {
	_, err := client.PostAndUnmarshal[meeting.UpdateMeetingResponse](s.client, ctx, "/cgi-bin/meeting/update", req)
	return err
}

// Cancel 取消预约会议
// 文档: https://developer.work.weixin.qq.com/document/path/93626
func (s *Service) Cancel(ctx context.Context, meetingID string) error {
	req := &meeting.CancelMeetingRequest{
		MeetingID: meetingID,
	}
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/meeting/cancel", req)
	return err
}

// GetInfo 获取会议详情
// 文档: https://developer.work.weixin.qq.com/document/path/93628
func (s *Service) GetInfo(ctx context.Context, meetingID string) (*meeting.GetMeetingInfoResponse, error) {
	req := &meeting.GetMeetingInfoRequest{
		MeetingID: meetingID,
	}
	return client.PostAndUnmarshal[meeting.GetMeetingInfoResponse](s.client, ctx, "/cgi-bin/meeting/get_info", req)
}

// GetUserMeetingIDs 获取成员会议ID列表
// 文档: https://developer.work.weixin.qq.com/document/path/93854
func (s *Service) GetUserMeetingIDs(ctx context.Context, req *meeting.GetUserMeetingIDsRequest) (*meeting.GetUserMeetingIDsResponse, error) {
	return client.PostAndUnmarshal[meeting.GetUserMeetingIDsResponse](s.client, ctx, "/cgi-bin/meeting/get_user_meetingid", req)
}

// GetStartList 获取会议发起记录
// 文档: https://developer.work.weixin.qq.com/document/path/96191
func (s *Service) GetStartList(ctx context.Context, req *meeting.GetMeetingStartListRequest) (*meeting.GetMeetingStartListResponse, error) {
	return client.PostAndUnmarshal[meeting.GetMeetingStartListResponse](s.client, ctx, "/cgi-bin/meeting/statistics/get_start_list", req)
}
