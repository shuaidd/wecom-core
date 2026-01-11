package calendar

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/calendar"
	"github.com/shuaidd/wecom-core/types/common"
)

const (
	addScheduleURL           = "/cgi-bin/oa/schedule/add"
	getScheduleURL           = "/cgi-bin/oa/schedule/get"
	updateScheduleURL        = "/cgi-bin/oa/schedule/update"
	deleteScheduleURL        = "/cgi-bin/oa/schedule/del"
	addAttendeesURL          = "/cgi-bin/oa/schedule/add_attendees"
	deleteAttendeesURL       = "/cgi-bin/oa/schedule/del_attendees"
	getScheduleByCalendarURL = "/cgi-bin/oa/schedule/get_by_calendar"
)

// CreateSchedule 创建日程
// 该接口用于在日历中创建一个日程
func (s *Service) CreateSchedule(ctx context.Context, req *calendar.CreateScheduleRequest) (*calendar.CreateScheduleResponse, error) {
	return client.PostAndUnmarshal[calendar.CreateScheduleResponse](s.client, ctx, addScheduleURL, req)
}

// GetSchedule 获取日程详情
// 该接口用于获取指定的日程详情
func (s *Service) GetSchedule(ctx context.Context, req *calendar.GetScheduleRequest) (*calendar.GetScheduleResponse, error) {
	return client.PostAndUnmarshal[calendar.GetScheduleResponse](s.client, ctx, getScheduleURL, req)
}

// UpdateSchedule 更新日程
// 该接口用于在日历中更新指定的日程
// 注意，更新操作是覆盖式，而不是增量式
func (s *Service) UpdateSchedule(ctx context.Context, req *calendar.UpdateScheduleRequest) (*calendar.UpdateScheduleResponse, error) {
	return client.PostAndUnmarshal[calendar.UpdateScheduleResponse](s.client, ctx, updateScheduleURL, req)
}

// DeleteSchedule 取消日程
// 该接口用于取消指定的日程
func (s *Service) DeleteSchedule(ctx context.Context, req *calendar.DeleteScheduleRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, deleteScheduleURL, req)
	return err
}

// AddAttendees 新增日程参与者
// 该接口用于在日历中更新指定的日程参与者列表
// 注意，该接口是增量式
func (s *Service) AddAttendees(ctx context.Context, req *calendar.AddAttendeesRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, addAttendeesURL, req)
	return err
}

// DeleteAttendees 删除日程参与者
// 该接口用于在日历中更新指定的日程参与者列表
// 注意，该接口是增量式
func (s *Service) DeleteAttendees(ctx context.Context, req *calendar.DeleteAttendeesRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, deleteAttendeesURL, req)
	return err
}

// GetScheduleByCalendar 获取日历下的日程列表
// 该接口用于获取指定的日历下的日程列表
// 仅可获取应用自己创建的日历下的日程
func (s *Service) GetScheduleByCalendar(ctx context.Context, req *calendar.GetScheduleByCalendarRequest) (*calendar.GetScheduleByCalendarResponse, error) {
	return client.PostAndUnmarshal[calendar.GetScheduleByCalendarResponse](s.client, ctx, getScheduleByCalendarURL, req)
}
