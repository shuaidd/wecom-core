package externalcontact

// ApplicableRange 敏感词适用范围
type ApplicableRange struct {
	UserList       []string `json:"user_list,omitempty"`
	DepartmentList []int    `json:"department_list,omitempty"`
}

// InterceptRule 敏感词规则
type InterceptRule struct {
	RuleID          string           `json:"rule_id,omitempty"`
	RuleName        string           `json:"rule_name"`
	WordList        []string         `json:"word_list,omitempty"`
	SemanticsList   []int            `json:"semantics_list,omitempty"`
	InterceptType   int              `json:"intercept_type,omitempty"`
	ApplicableRange *ApplicableRange `json:"applicable_range,omitempty"`
	CreateTime      int64            `json:"create_time,omitempty"`
}

// AddInterceptRuleRequest 新建敏感词规则请求
type AddInterceptRuleRequest struct {
	RuleName        string           `json:"rule_name"`
	WordList        []string         `json:"word_list"`
	SemanticsList   []int            `json:"semantics_list,omitempty"`
	InterceptType   int              `json:"intercept_type"`
	ApplicableRange *ApplicableRange `json:"applicable_range"`
}

// AddInterceptRuleResponse 新建敏感词规则响应
type AddInterceptRuleResponse struct {
	RuleID string `json:"rule_id"`
}

// GetInterceptRuleListResponse 获取敏感词规则列表响应
type GetInterceptRuleListResponse struct {
	RuleList []struct {
		RuleID     string `json:"rule_id"`
		RuleName   string `json:"rule_name"`
		CreateTime int64  `json:"create_time"`
	} `json:"rule_list"`
}

// GetInterceptRuleRequest 获取敏感词规则详情请求
type GetInterceptRuleRequest struct {
	RuleID string `json:"rule_id"`
}

// GetInterceptRuleResponse 获取敏感词规则详情响应
type GetInterceptRuleResponse struct {
	Rule InterceptRule `json:"rule"`
}

// UpdateInterceptRuleRequest 修改敏感词规则请求
type UpdateInterceptRuleRequest struct {
	RuleID                 string           `json:"rule_id"`
	RuleName               string           `json:"rule_name,omitempty"`
	WordList               []string         `json:"word_list,omitempty"`
	ExtraRule              *struct {
		SemanticsList []int `json:"semantics_list,omitempty"`
	} `json:"extra_rule,omitempty"`
	InterceptType          int              `json:"intercept_type,omitempty"`
	AddApplicableRange     *ApplicableRange `json:"add_applicable_range,omitempty"`
	RemoveApplicableRange  *ApplicableRange `json:"remove_applicable_range,omitempty"`
}

// DelInterceptRuleRequest 删除敏感词规则请求
type DelInterceptRuleRequest struct {
	RuleID string `json:"rule_id"`
}
