package externalcontact

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/externalcontact"
)

// ListStrategy 获取规则组列表
// 企业可通过此接口获取企业配置的所有客户规则组id列表
// 文档: https://developer.work.weixin.qq.com/document/path/94883
func (s *Service) ListStrategy(ctx context.Context, req *externalcontact.ListStrategyRequest) (*externalcontact.ListStrategyResponse, error) {
	return client.PostAndUnmarshal[externalcontact.ListStrategyResponse](s.client, ctx, "/cgi-bin/externalcontact/customer_strategy/list", req)
}

// GetStrategy 获取规则组详情
// 企业可通过此接口获取某个客户规则组的详细信息
// 文档: https://developer.work.weixin.qq.com/document/path/94883
func (s *Service) GetStrategy(ctx context.Context, strategyID int) (*externalcontact.GetStrategyResponse, error) {
	req := &externalcontact.GetStrategyRequest{
		StrategyID: strategyID,
	}
	return client.PostAndUnmarshal[externalcontact.GetStrategyResponse](s.client, ctx, "/cgi-bin/externalcontact/customer_strategy/get", req)
}

// GetStrategyRange 获取规则组管理范围
// 企业可通过此接口获取某个客户规则组管理的成员和部门列表
// 文档: https://developer.work.weixin.qq.com/document/path/94883
func (s *Service) GetStrategyRange(ctx context.Context, req *externalcontact.GetStrategyRangeRequest) (*externalcontact.GetStrategyRangeResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetStrategyRangeResponse](s.client, ctx, "/cgi-bin/externalcontact/customer_strategy/get_range", req)
}

// CreateStrategy 创建新的规则组
// 企业可通过此接口创建一个新的客户规则组
// 文档: https://developer.work.weixin.qq.com/document/path/94883
func (s *Service) CreateStrategy(ctx context.Context, req *externalcontact.CreateStrategyRequest) (*externalcontact.CreateStrategyResponse, error) {
	return client.PostAndUnmarshal[externalcontact.CreateStrategyResponse](s.client, ctx, "/cgi-bin/externalcontact/customer_strategy/create", req)
}

// EditStrategy 编辑规则组及其管理范围
// 企业可通过此接口编辑规则组的基本信息和修改客户规则组管理范围
// 文档: https://developer.work.weixin.qq.com/document/path/94883
func (s *Service) EditStrategy(ctx context.Context, req *externalcontact.EditStrategyRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/customer_strategy/edit", req)
	return err
}

// DeleteStrategy 删除规则组
// 企业可通过此接口删除某个规则组
// 文档: https://developer.work.weixin.qq.com/document/path/94883
func (s *Service) DeleteStrategy(ctx context.Context, strategyID int) error {
	req := &externalcontact.DeleteStrategyRequest{
		StrategyID: strategyID,
	}
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/customer_strategy/del", req)
	return err
}
