package reserve_meeting

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/common"
	"github.com/shuaidd/wecom-core/types/reserve_meeting"
)

// GetInvitees 获取会议受邀成员列表
// 文档: https://developer.work.weixin.qq.com/document/path/94051
func (s *Service) GetInvitees(ctx context.Context, meetingID string) (*reserve_meeting.GetInviteesResponse, error) {
	req := &reserve_meeting.GetInviteesRequest{
		MeetingID: meetingID,
	}
	return client.PostAndUnmarshal[reserve_meeting.GetInviteesResponse](s.client, ctx, "/cgi-bin/meeting/get_invitees", req)
}

// GetInviteesWithCursor 分页获取会议受邀成员列表
// 文档: https://developer.work.weixin.qq.com/document/path/94051
func (s *Service) GetInviteesWithCursor(ctx context.Context, meetingID, cursor string) (*reserve_meeting.GetInviteesResponse, error) {
	req := &reserve_meeting.GetInviteesRequest{
		MeetingID: meetingID,
		Cursor:    cursor,
	}
	return client.PostAndUnmarshal[reserve_meeting.GetInviteesResponse](s.client, ctx, "/cgi-bin/meeting/get_invitees", req)
}

// SetInvitees 更新会议受邀成员列表
// 文档: https://developer.work.weixin.qq.com/document/path/94050
func (s *Service) SetInvitees(ctx context.Context, req *reserve_meeting.SetInviteesRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/meeting/set_invitees", req)
	return err
}

// GetGuests 获取会议嘉宾列表
// 文档: https://developer.work.weixin.qq.com/document/path/94138
func (s *Service) GetGuests(ctx context.Context, meetingID string) (*reserve_meeting.GetGuestsResponse, error) {
	req := &reserve_meeting.GetGuestsRequest{
		MeetingID: meetingID,
	}
	return client.PostAndUnmarshal[reserve_meeting.GetGuestsResponse](s.client, ctx, "/cgi-bin/meeting/get_guests", req)
}

// SetGuests 更新会议嘉宾列表
// 文档: https://developer.work.weixin.qq.com/document/path/94137
func (s *Service) SetGuests(ctx context.Context, req *reserve_meeting.SetGuestsRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, "/cgi-bin/meeting/set_guests", req)
	return err
}

// GetRealtimeAttendeeList 获取实时会中成员列表
// 文档: https://developer.work.weixin.qq.com/document/path/93627
func (s *Service) GetRealtimeAttendeeList(ctx context.Context, meetingID string) (*reserve_meeting.GetRealtimeAttendeeListResponse, error) {
	req := &reserve_meeting.GetRealtimeAttendeeListRequest{
		MeetingID: meetingID,
	}
	return client.PostAndUnmarshal[reserve_meeting.GetRealtimeAttendeeListResponse](s.client, ctx, "/cgi-bin/meeting/get_realtime_attendee_list", req)
}

// GetRealtimeAttendeeListWithCursor 分页获取实时会中成员列表
// 文档: https://developer.work.weixin.qq.com/document/path/93627
func (s *Service) GetRealtimeAttendeeListWithCursor(ctx context.Context, req *reserve_meeting.GetRealtimeAttendeeListRequest) (*reserve_meeting.GetRealtimeAttendeeListResponse, error) {
	return client.PostAndUnmarshal[reserve_meeting.GetRealtimeAttendeeListResponse](s.client, ctx, "/cgi-bin/meeting/get_realtime_attendee_list", req)
}

// GetAttendeeList 获取已参会成员列表
// 文档: https://developer.work.weixin.qq.com/document/path/93736
func (s *Service) GetAttendeeList(ctx context.Context, meetingID string) (*reserve_meeting.GetAttendeeListResponse, error) {
	req := &reserve_meeting.GetAttendeeListRequest{
		MeetingID: meetingID,
	}
	return client.PostAndUnmarshal[reserve_meeting.GetAttendeeListResponse](s.client, ctx, "/cgi-bin/meeting/get_attendee_list", req)
}

// GetAttendeeListWithCursor 分页获取已参会成员列表
// 文档: https://developer.work.weixin.qq.com/document/path/93736
func (s *Service) GetAttendeeListWithCursor(ctx context.Context, req *reserve_meeting.GetAttendeeListRequest) (*reserve_meeting.GetAttendeeListResponse, error) {
	return client.PostAndUnmarshal[reserve_meeting.GetAttendeeListResponse](s.client, ctx, "/cgi-bin/meeting/get_attendee_list", req)
}

// WaitingRoomGetCurrentUserList 获取实时等候室成员列表
// 文档: https://developer.work.weixin.qq.com/document/path/93695
func (s *Service) WaitingRoomGetCurrentUserList(ctx context.Context, meetingID string) (*reserve_meeting.WaitingRoomGetCurrentUserListResponse, error) {
	req := &reserve_meeting.WaitingRoomGetCurrentUserListRequest{
		MeetingID: meetingID,
	}
	return client.PostAndUnmarshal[reserve_meeting.WaitingRoomGetCurrentUserListResponse](s.client, ctx, "/cgi-bin/meeting/waitingroom/get_current_user_list", req)
}

// WaitingRoomGetCurrentUserListWithCursor 分页获取实时等候室成员列表
// 文档: https://developer.work.weixin.qq.com/document/path/93695
func (s *Service) WaitingRoomGetCurrentUserListWithCursor(ctx context.Context, meetingID string, limit int32, cursor string) (*reserve_meeting.WaitingRoomGetCurrentUserListResponse, error) {
	req := &reserve_meeting.WaitingRoomGetCurrentUserListRequest{
		MeetingID: meetingID,
		Limit:     limit,
		Cursor:    cursor,
	}
	return client.PostAndUnmarshal[reserve_meeting.WaitingRoomGetCurrentUserListResponse](s.client, ctx, "/cgi-bin/meeting/waitingroom/get_current_user_list", req)
}

// WaitingRoomGetUserList 获取等候室成员记录
// 文档: https://developer.work.weixin.qq.com/document/path/94103
func (s *Service) WaitingRoomGetUserList(ctx context.Context, meetingID string) (*reserve_meeting.WaitingRoomGetUserListResponse, error) {
	req := &reserve_meeting.WaitingRoomGetUserListRequest{
		MeetingID: meetingID,
	}
	return client.PostAndUnmarshal[reserve_meeting.WaitingRoomGetUserListResponse](s.client, ctx, "/cgi-bin/meeting/waitingroom/get_user_list", req)
}

// WaitingRoomGetUserListWithCursor 分页获取等候室成员记录
// 文档: https://developer.work.weixin.qq.com/document/path/94103
func (s *Service) WaitingRoomGetUserListWithCursor(ctx context.Context, meetingID string, limit int32, cursor string) (*reserve_meeting.WaitingRoomGetUserListResponse, error) {
	req := &reserve_meeting.WaitingRoomGetUserListRequest{
		MeetingID: meetingID,
		Limit:     limit,
		Cursor:    cursor,
	}
	return client.PostAndUnmarshal[reserve_meeting.WaitingRoomGetUserListResponse](s.client, ctx, "/cgi-bin/meeting/waitingroom/get_user_list", req)
}

// CreateCustomerShortURL 创建用户专属参会链接
// 文档: https://developer.work.weixin.qq.com/document/path/93994
func (s *Service) CreateCustomerShortURL(ctx context.Context, req *reserve_meeting.CreateCustomerShortURLRequest) (*reserve_meeting.CreateCustomerShortURLResponse, error) {
	return client.PostAndUnmarshal[reserve_meeting.CreateCustomerShortURLResponse](s.client, ctx, "/cgi-bin/meeting/create_customer_short_url", req)
}

// GetCustomerShortURL 获取用户专属参会链接
// 文档: https://developer.work.weixin.qq.com/document/path/93993
func (s *Service) GetCustomerShortURL(ctx context.Context, meetingID string) (*reserve_meeting.GetCustomerShortURLResponse, error) {
	req := &reserve_meeting.GetCustomerShortURLRequest{
		MeetingID: meetingID,
	}
	return client.PostAndUnmarshal[reserve_meeting.GetCustomerShortURLResponse](s.client, ctx, "/cgi-bin/meeting/get_customer_short_url", req)
}
