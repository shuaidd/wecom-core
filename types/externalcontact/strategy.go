package externalcontact

// StrategyPrivilege 规则组权限配置
type StrategyPrivilege struct {
	ViewCustomerList        bool `json:"view_customer_list"`
	ViewCustomerData        bool `json:"view_customer_data"`
	ViewRoomList            bool `json:"view_room_list"`
	ContactMe               bool `json:"contact_me"`
	JoinRoom                bool `json:"join_room"`
	ShareCustomer           bool `json:"share_customer"`
	OperResignCustomer      bool `json:"oper_resign_customer"`
	OperResignGroup         bool `json:"oper_resign_group"`
	SendCustomerMsg         bool `json:"send_customer_msg"`
	EditWelcomeMsg          bool `json:"edit_welcome_msg"`
	ViewBehaviorData        bool `json:"view_behavior_data"`
	ViewRoomData            bool `json:"view_room_data"`
	SendGroupMsg            bool `json:"send_group_msg"`
	RoomDeduplication       bool `json:"room_deduplication"`
	RapidReply              bool `json:"rapid_reply"`
	OnjobCustomerTransfer   bool `json:"onjob_customer_transfer"`
	EditAntiSpamRule        bool `json:"edit_anti_spam_rule"`
	ExportCustomerList      bool `json:"export_customer_list"`
	ExportCustomerData      bool `json:"export_customer_data"`
	ExportCustomerGroupList bool `json:"export_customer_group_list"`
	ManageCustomerTag       bool `json:"manage_customer_tag"`
}

// StrategyRange 规则组管理范围
type StrategyRange struct {
	Type    int    `json:"type"`
	UserID  string `json:"userid,omitempty"`
	PartyID int    `json:"partyid,omitempty"`
}

// Strategy 客户联系规则组
type Strategy struct {
	StrategyID   int                `json:"strategy_id"`
	ParentID     int                `json:"parent_id,omitempty"`
	StrategyName string             `json:"strategy_name"`
	CreateTime   int64              `json:"create_time,omitempty"`
	AdminList    []string           `json:"admin_list,omitempty"`
	Privilege    *StrategyPrivilege `json:"privilege,omitempty"`
}

// ListStrategyRequest 获取规则组列表请求
type ListStrategyRequest struct {
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

// StrategyItem 规则组项
type StrategyItem struct {
	StrategyID int `json:"strategy_id"`
}

// ListStrategyResponse 获取规则组列表响应
type ListStrategyResponse struct {
	Strategy   []StrategyItem `json:"strategy"`
	NextCursor string         `json:"next_cursor,omitempty"`
}

// GetStrategyRequest 获取规则组详情请求
type GetStrategyRequest struct {
	StrategyID int `json:"strategy_id"`
}

// GetStrategyResponse 获取规则组详情响应
type GetStrategyResponse struct {
	Strategy Strategy `json:"strategy"`
}

// GetStrategyRangeRequest 获取规则组管理范围请求
type GetStrategyRangeRequest struct {
	StrategyID int    `json:"strategy_id"`
	Cursor     string `json:"cursor,omitempty"`
	Limit      int    `json:"limit,omitempty"`
}

// GetStrategyRangeResponse 获取规则组管理范围响应
type GetStrategyRangeResponse struct {
	Range      []StrategyRange `json:"range"`
	NextCursor string          `json:"next_cursor,omitempty"`
}

// CreateStrategyRequest 创建规则组请求
type CreateStrategyRequest struct {
	ParentID     int                `json:"parent_id,omitempty"`
	StrategyName string             `json:"strategy_name"`
	AdminList    []string           `json:"admin_list"`
	Privilege    *StrategyPrivilege `json:"privilege,omitempty"`
	Range        []StrategyRange    `json:"range,omitempty"`
}

// CreateStrategyResponse 创建规则组响应
type CreateStrategyResponse struct {
	StrategyID int `json:"strategy_id"`
}

// EditStrategyRequest 编辑规则组请求
type EditStrategyRequest struct {
	StrategyID   int                `json:"strategy_id"`
	StrategyName string             `json:"strategy_name,omitempty"`
	AdminList    []string           `json:"admin_list,omitempty"`
	Privilege    *StrategyPrivilege `json:"privilege,omitempty"`
	RangeAdd     []StrategyRange    `json:"range_add,omitempty"`
	RangeDel     []StrategyRange    `json:"range_del,omitempty"`
}

// DeleteStrategyRequest 删除规则组请求
type DeleteStrategyRequest struct {
	StrategyID int `json:"strategy_id"`
}
