package updown

import "github.com/shuaidd/wecom-core/types/common"

// OwnerCorpRange 上游企业的对接人规则
type OwnerCorpRange struct {
	DepartmentIDs []string `json:"departmentids,omitempty"` // 部门id列表
	UserIDs       []string `json:"userids,omitempty"`       // 用户id列表
}

// MemberCorpRange 下游企业规则范围
type MemberCorpRange struct {
	GroupIDs []string `json:"groupids,omitempty"` // 分组id列表
	CorpIDs  []string `json:"corpids,omitempty"`  // 企业id列表
}

// RuleInfo 对接规则详情
type RuleInfo struct {
	OwnerCorpRange  *OwnerCorpRange  `json:"owner_corp_range"`  // 上游企业的对接人规则
	MemberCorpRange *MemberCorpRange `json:"member_corp_range"` // 下游企业规则范围
}

// AddRuleRequest 新增对接规则请求
type AddRuleRequest struct {
	ChainID  string    `json:"chain_id"`  // 上下游id
	RuleInfo *RuleInfo `json:"rule_info"` // 上下游关系规则的详情
}

// AddRuleResponse 新增对接规则响应
type AddRuleResponse struct {
	common.Response
	RuleID int64 `json:"rule_id"` // 上下游规则id
}

// ModifyRuleRequest 更新对接规则请求
type ModifyRuleRequest struct {
	ChainID  string    `json:"chain_id"`  // 上下游id
	RuleID   int64     `json:"rule_id"`   // 上下游规则id
	RuleInfo *RuleInfo `json:"rule_info"` // 上下游关系规则的详情
}

// ModifyRuleResponse 更新对接规则响应
type ModifyRuleResponse struct {
	common.Response
}

// DeleteRuleRequest 删除对接规则请求
type DeleteRuleRequest struct {
	ChainID string `json:"chain_id"` // 上下游id
	RuleID  int64  `json:"rule_id"`  // 上下游规则id
}

// DeleteRuleResponse 删除对接规则响应
type DeleteRuleResponse struct {
	common.Response
}

// ListRuleIDsRequest 获取对接规则id列表请求
type ListRuleIDsRequest struct {
	ChainID string `json:"chain_id"` // 上下游id
}

// ListRuleIDsResponse 获取对接规则id列表响应
type ListRuleIDsResponse struct {
	common.Response
	RuleIDs []int64 `json:"rule_ids"` // 上下游关系规则的id列表
}

// GetRuleInfoRequest 获取对接规则详情请求
type GetRuleInfoRequest struct {
	ChainID string `json:"chain_id"` // 上下游id
	RuleID  int64  `json:"rule_id"`  // 上下游规则id
}

// GetRuleInfoResponse 获取对接规则详情响应
type GetRuleInfoResponse struct {
	common.Response
	RuleInfo *RuleInfo `json:"rule_info"` // 上下游关系规则的详情
}
