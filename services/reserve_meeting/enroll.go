package reserve_meeting

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/reserve_meeting"
)

// EnrollQueryByTmpOpenID 获取会议成员报名ID
// 文档: https://developer.work.weixin.qq.com/document/path/94415
func (s *Service) EnrollQueryByTmpOpenID(ctx context.Context, req *reserve_meeting.EnrollQueryByTmpOpenIDRequest) (*reserve_meeting.EnrollQueryByTmpOpenIDResponse, error) {
	return client.PostAndUnmarshal[reserve_meeting.EnrollQueryByTmpOpenIDResponse](s.client, ctx, "/cgi-bin/meeting/enroll/query_by_tmp_openid", req)
}

// EnrollList 获取会议报名信息
// 文档: https://developer.work.weixin.qq.com/document/path/93853
func (s *Service) EnrollList(ctx context.Context, meetingID string) (*reserve_meeting.EnrollListResponse, error) {
	req := &reserve_meeting.EnrollListRequest{
		MeetingID: meetingID,
	}
	return client.PostAndUnmarshal[reserve_meeting.EnrollListResponse](s.client, ctx, "/cgi-bin/meeting/enroll/list", req)
}

// EnrollListWithCursor 分页获取会议报名信息
// 文档: https://developer.work.weixin.qq.com/document/path/93853
func (s *Service) EnrollListWithCursor(ctx context.Context, req *reserve_meeting.EnrollListRequest) (*reserve_meeting.EnrollListResponse, error) {
	return client.PostAndUnmarshal[reserve_meeting.EnrollListResponse](s.client, ctx, "/cgi-bin/meeting/enroll/list", req)
}

// EnrollListByStatus 按状态获取会议报名信息
// 文档: https://developer.work.weixin.qq.com/document/path/93853
func (s *Service) EnrollListByStatus(ctx context.Context, meetingID string, status int32) (*reserve_meeting.EnrollListResponse, error) {
	req := &reserve_meeting.EnrollListRequest{
		MeetingID: meetingID,
		Status:    status,
	}
	return client.PostAndUnmarshal[reserve_meeting.EnrollListResponse](s.client, ctx, "/cgi-bin/meeting/enroll/list", req)
}

// EnrollGetConfig 获取会议报名配置
// 文档: https://developer.work.weixin.qq.com/document/path/93697
func (s *Service) EnrollGetConfig(ctx context.Context, meetingID string) (*reserve_meeting.EnrollGetConfigResponse, error) {
	req := &reserve_meeting.EnrollGetConfigRequest{
		MeetingID: meetingID,
	}
	return client.PostAndUnmarshal[reserve_meeting.EnrollGetConfigResponse](s.client, ctx, "/cgi-bin/meeting/enroll/get_config", req)
}

// EnrollSetConfig 修改会议报名配置
// 文档: https://developer.work.weixin.qq.com/document/path/93696
func (s *Service) EnrollSetConfig(ctx context.Context, req *reserve_meeting.EnrollSetConfigRequest) (*reserve_meeting.EnrollSetConfigResponse, error) {
	return client.PostAndUnmarshal[reserve_meeting.EnrollSetConfigResponse](s.client, ctx, "/cgi-bin/meeting/enroll/set_config", req)
}

// EnrollDelete 删除会议报名信息
// 文档: https://developer.work.weixin.qq.com/document/path/93700
func (s *Service) EnrollDelete(ctx context.Context, req *reserve_meeting.EnrollDeleteRequest) (*reserve_meeting.EnrollDeleteResponse, error) {
	return client.PostAndUnmarshal[reserve_meeting.EnrollDeleteResponse](s.client, ctx, "/cgi-bin/meeting/enroll/delete", req)
}

// EnrollImport 导入会议报名信息
// 文档: https://developer.work.weixin.qq.com/document/path/93856
func (s *Service) EnrollImport(ctx context.Context, req *reserve_meeting.EnrollImportRequest) (*reserve_meeting.EnrollImportResponse, error) {
	return client.PostAndUnmarshal[reserve_meeting.EnrollImportResponse](s.client, ctx, "/cgi-bin/meeting/enroll/import", req)
}

// EnrollApprove 审批会议报名信息
// 文档: https://developer.work.weixin.qq.com/document/path/93701
func (s *Service) EnrollApprove(ctx context.Context, req *reserve_meeting.EnrollApproveRequest) (*reserve_meeting.EnrollApproveResponse, error) {
	return client.PostAndUnmarshal[reserve_meeting.EnrollApproveResponse](s.client, ctx, "/cgi-bin/meeting/enroll/approve", req)
}
