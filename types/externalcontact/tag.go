package externalcontact

// CorpTag 企业客户标签
type CorpTag struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time,omitempty"`
	Order      uint32 `json:"order,omitempty"`
	Deleted    bool   `json:"deleted,omitempty"`
}

// CorpTagGroup 企业客户标签组
type CorpTagGroup struct {
	GroupID    string    `json:"group_id"`
	GroupName  string    `json:"group_name"`
	CreateTime int64     `json:"create_time,omitempty"`
	Order      uint32    `json:"order,omitempty"`
	Deleted    bool      `json:"deleted,omitempty"`
	StrategyID int       `json:"strategy_id,omitempty"`
	Tag        []CorpTag `json:"tag,omitempty"`
}

// GetCorpTagListRequest 获取企业标签库请求
type GetCorpTagListRequest struct {
	TagID   []string `json:"tag_id,omitempty"`
	GroupID []string `json:"group_id,omitempty"`
}

// GetCorpTagListResponse 获取企业标签库响应
type GetCorpTagListResponse struct {
	TagGroup []CorpTagGroup `json:"tag_group"`
}

// AddCorpTagRequest 添加企业客户标签请求
type AddCorpTagRequest struct {
	GroupID   string           `json:"group_id,omitempty"`
	GroupName string           `json:"group_name,omitempty"`
	Order     uint32           `json:"order,omitempty"`
	Tag       []AddCorpTagItem `json:"tag"`
	AgentID   int              `json:"agentid,omitempty"`
}

// AddCorpTagItem 添加标签项
type AddCorpTagItem struct {
	Name  string `json:"name"`
	Order uint32 `json:"order,omitempty"`
}

// AddCorpTagResponse 添加企业客户标签响应
type AddCorpTagResponse struct {
	TagGroup CorpTagGroup `json:"tag_group"`
}

// EditCorpTagRequest 编辑企业客户标签请求
type EditCorpTagRequest struct {
	ID      string `json:"id"`
	Name    string `json:"name,omitempty"`
	Order   uint32 `json:"order,omitempty"`
	AgentID int    `json:"agentid,omitempty"`
}

// DeleteCorpTagRequest 删除企业客户标签请求
type DeleteCorpTagRequest struct {
	TagID   []string `json:"tag_id,omitempty"`
	GroupID []string `json:"group_id,omitempty"`
	AgentID int      `json:"agentid,omitempty"`
}

// MarkTagRequest 编辑客户企业标签请求
type MarkTagRequest struct {
	UserID         string   `json:"userid"`
	ExternalUserID string   `json:"external_userid"`
	AddTag         []string `json:"add_tag,omitempty"`
	RemoveTag      []string `json:"remove_tag,omitempty"`
}

// GetStrategyTagListRequest 获取指定规则组下的企业客户标签请求
type GetStrategyTagListRequest struct {
	StrategyID int      `json:"strategy_id,omitempty"`
	TagID      []string `json:"tag_id,omitempty"`
	GroupID    []string `json:"group_id,omitempty"`
}

// GetStrategyTagListResponse 获取指定规则组下的企业客户标签响应
type GetStrategyTagListResponse struct {
	TagGroup []CorpTagGroup `json:"tag_group"`
}

// AddStrategyTagRequest 为指定规则组创建企业客户标签请求
type AddStrategyTagRequest struct {
	StrategyID int              `json:"strategy_id"`
	GroupID    string           `json:"group_id,omitempty"`
	GroupName  string           `json:"group_name,omitempty"`
	Order      uint32           `json:"order,omitempty"`
	Tag        []AddCorpTagItem `json:"tag"`
}

// AddStrategyTagResponse 为指定规则组创建企业客户标签响应
type AddStrategyTagResponse struct {
	TagGroup CorpTagGroup `json:"tag_group"`
}

// EditStrategyTagRequest 编辑指定规则组下的企业客户标签请求
type EditStrategyTagRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Order uint32 `json:"order,omitempty"`
}

// DeleteStrategyTagRequest 删除指定规则组下的企业客户标签请求
type DeleteStrategyTagRequest struct {
	TagID   []string `json:"tag_id,omitempty"`
	GroupID []string `json:"group_id,omitempty"`
}
