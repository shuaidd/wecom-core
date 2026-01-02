package updown

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/updown"
)

// AddRule 新增对接规则
// 文档: https://developer.work.weixin.qq.com/document/path/95792
func (s *Service) AddRule(ctx context.Context, req *updown.AddRuleRequest) (int64, error) {
	result, err := client.PostAndUnmarshal[updown.AddRuleResponse](s.client, ctx, "/cgi-bin/corpgroup/rule/add_rule", req)
	if err != nil {
		return 0, err
	}
	return result.RuleID, nil
}

// ModifyRule 更新对接规则
// 文档: https://developer.work.weixin.qq.com/document/path/95793
func (s *Service) ModifyRule(ctx context.Context, req *updown.ModifyRuleRequest) error {
	_, err := client.PostAndUnmarshal[updown.ModifyRuleResponse](s.client, ctx, "/cgi-bin/corpgroup/rule/modify_rule", req)
	return err
}

// DeleteRule 删除对接规则
// 文档: https://developer.work.weixin.qq.com/document/path/95794
func (s *Service) DeleteRule(ctx context.Context, req *updown.DeleteRuleRequest) error {
	_, err := client.PostAndUnmarshal[updown.DeleteRuleResponse](s.client, ctx, "/cgi-bin/corpgroup/rule/delete_rule", req)
	return err
}

// ListRuleIDs 获取对接规则id列表
// 文档: https://developer.work.weixin.qq.com/document/path/95795
func (s *Service) ListRuleIDs(ctx context.Context, req *updown.ListRuleIDsRequest) ([]int64, error) {
	result, err := client.PostAndUnmarshal[updown.ListRuleIDsResponse](s.client, ctx, "/cgi-bin/corpgroup/rule/list_ids", req)
	if err != nil {
		return nil, err
	}
	return result.RuleIDs, nil
}

// GetRuleInfo 获取对接规则详情
// 文档: https://developer.work.weixin.qq.com/document/path/95796
func (s *Service) GetRuleInfo(ctx context.Context, req *updown.GetRuleInfoRequest) (*updown.RuleInfo, error) {
	result, err := client.PostAndUnmarshal[updown.GetRuleInfoResponse](s.client, ctx, "/cgi-bin/corpgroup/rule/get_rule_info", req)
	if err != nil {
		return nil, err
	}
	return result.RuleInfo, nil
}
