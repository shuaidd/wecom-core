package externalcontact

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/externalcontact"
)

// GetGroupChatStatistic 获取群聊数据统计（按群主聚合）
// 获取指定日期的统计数据，按群主聚合的方式
// 文档: https://developer.work.weixin.qq.com/document/path/92133
func (s *Service) GetGroupChatStatistic(ctx context.Context, req *externalcontact.GroupChatStatisticRequest) (*externalcontact.GroupChatStatisticResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GroupChatStatisticResponse](s.client, ctx, "/cgi-bin/externalcontact/groupchat/statistic", req)
}

// GetGroupChatStatisticGroupByDay 获取群聊数据统计（按自然日聚合）
// 获取指定日期的统计数据，按自然日聚合的方式
// 文档: https://developer.work.weixin.qq.com/document/path/92133
func (s *Service) GetGroupChatStatisticGroupByDay(ctx context.Context, req *externalcontact.GroupChatStatisticGroupByDayRequest) (*externalcontact.GroupChatStatisticGroupByDayResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GroupChatStatisticGroupByDayResponse](s.client, ctx, "/cgi-bin/externalcontact/groupchat/statistic_group_by_day", req)
}

// GetUserBehaviorData 获取联系客户统计数据
// 企业可通过此接口获取成员联系客户的数据，包括发起申请数、新增客户数、聊天数、发送消息数和删除/拉黑成员的客户数等指标
// 文档: https://developer.work.weixin.qq.com/document/path/92132
func (s *Service) GetUserBehaviorData(ctx context.Context, req *externalcontact.GetUserBehaviorDataRequest) (*externalcontact.GetUserBehaviorDataResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetUserBehaviorDataResponse](s.client, ctx, "/cgi-bin/externalcontact/get_user_behavior_data", req)
}
