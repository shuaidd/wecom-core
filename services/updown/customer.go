package updown

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/updown"
)

// UnionIDToExternalUserID 通过unionid和openid查询external_userid
// 文档: https://developer.work.weixin.qq.com/document/path/95818
func (s *Service) UnionIDToExternalUserID(ctx context.Context, req *updown.UnionIDToExternalUserIDRequest) ([]updown.ExternalUserIDInfo, error) {
	result, err := client.PostAndUnmarshal[updown.UnionIDToExternalUserIDResponse](s.client, ctx, "/cgi-bin/corpgroup/unionid_to_external_userid", req)
	if err != nil {
		return nil, err
	}
	return result.ExternalUserIDInfo, nil
}

// UnionIDToPendingID unionid查询pending_id
// 文档: https://developer.work.weixin.qq.com/document/path/97357
func (s *Service) UnionIDToPendingID(ctx context.Context, req *updown.UnionIDToPendingIDRequest) (string, error) {
	result, err := client.PostAndUnmarshal[updown.UnionIDToPendingIDResponse](s.client, ctx, "/cgi-bin/corpgroup/unionid_to_pending_id", req)
	if err != nil {
		return "", err
	}
	return result.PendingID, nil
}

// ExternalUserIDToPendingID external_userid查询pending_id
// 文档: https://developer.work.weixin.qq.com/document/path/97357
func (s *Service) ExternalUserIDToPendingID(ctx context.Context, req *updown.ExternalUserIDToPendingIDRequest) ([]updown.PendingIDResult, error) {
	result, err := client.PostAndUnmarshal[updown.ExternalUserIDToPendingIDResponse](s.client, ctx, "/cgi-bin/corpgroup/batch/external_userid_to_pending_id", req)
	if err != nil {
		return nil, err
	}
	return result.Result, nil
}
