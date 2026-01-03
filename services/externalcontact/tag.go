package externalcontact

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/externalcontact"
)

// GetCorpTagList 获取企业标签库
// 企业可通过此接口获取企业客户标签详情
// 文档: https://developer.work.weixin.qq.com/document/path/92117
func (s *Service) GetCorpTagList(ctx context.Context, req *externalcontact.GetCorpTagListRequest) (*externalcontact.GetCorpTagListResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetCorpTagListResponse](s.client, ctx, "/cgi-bin/externalcontact/get_corp_tag_list", req)
}

// AddCorpTag 添加企业客户标签
// 企业可通过此接口向客户标签库中添加新的标签组和标签
// 文档: https://developer.work.weixin.qq.com/document/path/92117
func (s *Service) AddCorpTag(ctx context.Context, req *externalcontact.AddCorpTagRequest) (*externalcontact.AddCorpTagResponse, error) {
	return client.PostAndUnmarshal[externalcontact.AddCorpTagResponse](s.client, ctx, "/cgi-bin/externalcontact/add_corp_tag", req)
}

// EditCorpTag 编辑企业客户标签
// 企业可通过此接口编辑客户标签/标签组的名称或次序值
// 文档: https://developer.work.weixin.qq.com/document/path/92117
func (s *Service) EditCorpTag(ctx context.Context, req *externalcontact.EditCorpTagRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/edit_corp_tag", req)
	return err
}

// DeleteCorpTag 删除企业客户标签
// 企业可通过此接口删除客户标签库中的标签，或删除整个标签组
// 文档: https://developer.work.weixin.qq.com/document/path/92117
func (s *Service) DeleteCorpTag(ctx context.Context, req *externalcontact.DeleteCorpTagRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/del_corp_tag", req)
	return err
}

// MarkTag 编辑客户企业标签
// 企业可通过此接口为指定成员的客户添加上由企业统一配置的标签
// 文档: https://developer.work.weixin.qq.com/document/path/92118
func (s *Service) MarkTag(ctx context.Context, req *externalcontact.MarkTagRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/mark_tag", req)
	return err
}

// GetStrategyTagList 获取指定规则组下的企业客户标签
// 企业可通过此接口获取某个规则组内的企业客户标签详情
// 文档: https://developer.work.weixin.qq.com/document/path/94882
func (s *Service) GetStrategyTagList(ctx context.Context, req *externalcontact.GetStrategyTagListRequest) (*externalcontact.GetStrategyTagListResponse, error) {
	return client.PostAndUnmarshal[externalcontact.GetStrategyTagListResponse](s.client, ctx, "/cgi-bin/externalcontact/get_strategy_tag_list", req)
}

// AddStrategyTag 为指定规则组创建企业客户标签
// 企业可通过此接口向规则组中添加新的标签组和标签
// 文档: https://developer.work.weixin.qq.com/document/path/94882
func (s *Service) AddStrategyTag(ctx context.Context, req *externalcontact.AddStrategyTagRequest) (*externalcontact.AddStrategyTagResponse, error) {
	return client.PostAndUnmarshal[externalcontact.AddStrategyTagResponse](s.client, ctx, "/cgi-bin/externalcontact/add_strategy_tag", req)
}

// EditStrategyTag 编辑指定规则组下的企业客户标签
// 企业可通过此接口编辑指定规则组下的客户标签/标签组的名称或次序值
// 文档: https://developer.work.weixin.qq.com/document/path/94882
func (s *Service) EditStrategyTag(ctx context.Context, req *externalcontact.EditStrategyTagRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/edit_strategy_tag", req)
	return err
}

// DeleteStrategyTag 删除指定规则组下的企业客户标签
// 企业可通过此接口删除某个规则组下的标签，或删除整个标签组
// 文档: https://developer.work.weixin.qq.com/document/path/94882
func (s *Service) DeleteStrategyTag(ctx context.Context, req *externalcontact.DeleteStrategyTagRequest) error {
	type response struct{}
	_, err := client.PostAndUnmarshal[response](s.client, ctx, "/cgi-bin/externalcontact/del_strategy_tag", req)
	return err
}
