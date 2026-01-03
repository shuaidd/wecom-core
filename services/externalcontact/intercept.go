package externalcontact

import (
	"context"
	"net/url"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/externalcontact"
)

// AddInterceptRule 新建敏感词规则
// 企业和第三方应用可以通过此接口新建敏感词规则
// 文档: https://developer.work.weixin.qq.com/document/path/95130
func (s *Service) AddInterceptRule(ctx context.Context, req *externalcontact.AddInterceptRuleRequest) (*externalcontact.AddInterceptRuleResponse, error) {
	return client.PostAndUnmarshal[externalcontact.AddInterceptRuleResponse](s.client, ctx, "/cgi-bin/externalcontact/add_intercept_rule", req)
}

// GetInterceptRuleList 获取敏感词规则列表
// 企业和第三方应用可以通过此接口获取敏感词规则列表
// 文档: https://developer.work.weixin.qq.com/document/path/95130
func (s *Service) GetInterceptRuleList(ctx context.Context) (*externalcontact.GetInterceptRuleListResponse, error) {
	return client.GetAndUnmarshal[externalcontact.GetInterceptRuleListResponse](s.client, ctx, "/cgi-bin/externalcontact/get_intercept_rule_list", url.Values{})
}

// GetInterceptRule 获取敏感词规则详情
// 企业和第三方应用可以通过此接口获取敏感词规则详情
// 文档: https://developer.work.weixin.qq.com/document/path/95130
func (s *Service) GetInterceptRule(ctx context.Context, req *externalcontact.GetInterceptRuleRequest) (*externalcontact.GetInterceptRuleResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetInterceptRuleResponse](s.client, ctx, "/cgi-bin/externalcontact/get_intercept_rule", req)
}

// UpdateInterceptRule 修改敏感词规则
// 企业和第三方应用可以通过此接口修改敏感词规则
// 文档: https://developer.work.weixin.qq.com/document/path/95130
func (s *Service) UpdateInterceptRule(ctx context.Context, req *externalcontact.UpdateInterceptRuleRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/update_intercept_rule", req)
	return err
}

// DelInterceptRule 删除敏感词规则
// 企业和第三方应用可以通过此接口删除敏感词规则
// 文档: https://developer.work.weixin.qq.com/document/path/95130
func (s *Service) DelInterceptRule(ctx context.Context, req *externalcontact.DelInterceptRuleRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/del_intercept_rule", req)
	return err
}
