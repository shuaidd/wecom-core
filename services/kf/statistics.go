package kf

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/kf"
)

// GetCorpStatistic 获取「客户数据统计」企业汇总数据
// 通过此接口，可以获取咨询会话数、咨询客户数等企业汇总统计数据
// 文档: https://developer.work.weixin.qq.com/document/path/95489
func (s *Service) GetCorpStatistic(ctx context.Context, req *kf.GetCorpStatisticRequest) (*kf.GetCorpStatisticResponse, error) {
	return client.PostAndUnmarshal[kf.GetCorpStatisticResponse](s.client, ctx, "/cgi-bin/kf/get_corp_statistic", req)
}

// GetServicerStatistic 获取「客户数据统计」接待人员明细数据
// 通过此接口，可获取接入人工会话数、咨询会话数等与接待人员相关的统计信息
// 文档: https://developer.work.weixin.qq.com/document/path/95490
func (s *Service) GetServicerStatistic(ctx context.Context, req *kf.GetServicerStatisticRequest) (*kf.GetServicerStatisticResponse, error) {
	return client.PostAndUnmarshal[kf.GetServicerStatisticResponse](s.client, ctx, "/cgi-bin/kf/get_servicer_statistic", req)
}
