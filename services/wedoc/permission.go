package wedoc

import (
	"context"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/common"
	"github.com/shuaidd/wecom-core/types/wedoc"
)

const (
	getDocAuthURL         = "/cgi-bin/wedoc/doc_get_auth"
	modDocSaftySettingURL = "/cgi-bin/wedoc/mod_doc_safty_setting"
	modDocJoinRuleURL     = "/cgi-bin/wedoc/mod_doc_join_rule"
	modDocMemberURL       = "/cgi-bin/wedoc/mod_doc_member"
	getSheetPrivURL       = "/cgi-bin/wedoc/smartsheet/content_priv/get_sheet_priv"
	updateSheetPrivURL    = "/cgi-bin/wedoc/smartsheet/content_priv/update_sheet_priv"
	createRuleURL         = "/cgi-bin/wedoc/smartsheet/content_priv/create_rule"
	modRuleMemberURL      = "/cgi-bin/wedoc/smartsheet/content_priv/mod_rule_member"
	deleteRuleURL         = "/cgi-bin/wedoc/smartsheet/content_priv/delete_rule"
)

// GetDocAuth 获取文档权限信息
// 该接口用于获取文档、表格、智能表格的查看规则、文档通知范围及权限、安全设置信息
func (s *Service) GetDocAuth(ctx context.Context, req *wedoc.GetDocAuthRequest) (*wedoc.GetDocAuthResponse, error) {
	return client.PostAndUnmarshal[wedoc.GetDocAuthResponse](s.client, ctx, getDocAuthURL, req)
}

// ModDocSaftySetting 修改文档安全设置
// 该接口用于修改文档、表格、智能表格的安全设置
func (s *Service) ModDocSaftySetting(ctx context.Context, req *wedoc.ModDocSaftySettingRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, modDocSaftySettingURL, req)
	return err
}

// ModDocJoinRule 修改文档查看规则
// 该接口用于修改文档、表格、智能表格查看规则
func (s *Service) ModDocJoinRule(ctx context.Context, req *wedoc.ModDocJoinRuleRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, modDocJoinRuleURL, req)
	return err
}

// ModDocMember 修改文档通知范围及权限
// 该接口用于修改文档、表格、智能表格通知范围列表，可以新增文档、表格、智能表格通知范围并设置权限、修改已有范围的权限以及删除文档、表格、智能表格通知范围内的人员
func (s *Service) ModDocMember(ctx context.Context, req *wedoc.ModDocMemberRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, modDocMemberURL, req)
	return err
}

// GetSheetPriv 查询智能表格子表权限
// 该接口用于查询智能表格子表权限详情
func (s *Service) GetSheetPriv(ctx context.Context, req *wedoc.GetSheetPrivRequest) (*wedoc.GetSheetPrivResponse, error) {
	return client.PostAndUnmarshal[wedoc.GetSheetPrivResponse](s.client, ctx, getSheetPrivURL, req)
}

// UpdateSheetPriv 更新智能表格子表权限
// 该接口用于设置全员权限或者成员额外权限的权限详情
func (s *Service) UpdateSheetPriv(ctx context.Context, req *wedoc.UpdateSheetPrivRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, updateSheetPrivURL, req)
	return err
}

// CreateRule 新增智能表格指定成员额外权限
// 该接口用于新增智能表格指定成员额外权限
func (s *Service) CreateRule(ctx context.Context, req *wedoc.CreateRuleRequest) (*wedoc.CreateRuleResponse, error) {
	return client.PostAndUnmarshal[wedoc.CreateRuleResponse](s.client, ctx, createRuleURL, req)
}

// ModRuleMember 更新智能表格指定成员额外权限
// 该接口用于更新智能表格指定成员额外权限，成员最多可设置50个
func (s *Service) ModRuleMember(ctx context.Context, req *wedoc.ModRuleMemberRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, modRuleMemberURL, req)
	return err
}

// DeleteRule 删除智能表格指定成员额外权限
// 该接口用于删除智能表格指定成员额外权限
func (s *Service) DeleteRule(ctx context.Context, req *wedoc.DeleteRuleRequest) error {
	_, err := client.PostAndUnmarshal[common.Response](s.client, ctx, deleteRuleURL, req)
	return err
}
