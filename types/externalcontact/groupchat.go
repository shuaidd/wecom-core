package externalcontact

// OwnerFilter 群主过滤
type OwnerFilter struct {
	UserIDList []string `json:"userid_list,omitempty"`
}

// ListGroupChatRequest 获取客户群列表请求
type ListGroupChatRequest struct {
	StatusFilter int          `json:"status_filter,omitempty"`
	OwnerFilter  *OwnerFilter `json:"owner_filter,omitempty"`
	Cursor       string       `json:"cursor,omitempty"`
	Limit        int          `json:"limit"`
}

// GroupChatItem 客户群列表项
type GroupChatItem struct {
	ChatID string `json:"chat_id"`
	Status int    `json:"status"`
}

// ListGroupChatResponse 获取客户群列表响应
type ListGroupChatResponse struct {
	GroupChatList []GroupChatItem `json:"group_chat_list"`
	NextCursor    string          `json:"next_cursor,omitempty"`
}

// GetGroupChatRequest 获取客户群详情请求
type GetGroupChatRequest struct {
	ChatID   string `json:"chat_id"`
	NeedName int    `json:"need_name,omitempty"`
}

// GroupChatInvitor 邀请者
type GroupChatInvitor struct {
	UserID string `json:"userid"`
}

// GroupChatMember 客户群成员
type GroupChatMember struct {
	UserID        string            `json:"userid"`
	Type          int               `json:"type"`
	UnionID       string            `json:"unionid,omitempty"`
	JoinTime      int64             `json:"join_time"`
	JoinScene     int               `json:"join_scene"`
	Invitor       *GroupChatInvitor `json:"invitor,omitempty"`
	GroupNickname string            `json:"group_nickname,omitempty"`
	Name          string            `json:"name,omitempty"`
}

// GroupChatAdmin 客户群管理员
type GroupChatAdmin struct {
	UserID string `json:"userid"`
}

// GroupChat 客户群详情
type GroupChat struct {
	ChatID        string            `json:"chat_id"`
	Name          string            `json:"name"`
	Owner         string            `json:"owner"`
	CreateTime    int64             `json:"create_time"`
	Notice        string            `json:"notice,omitempty"`
	MemberList    []GroupChatMember `json:"member_list"`
	AdminList     []GroupChatAdmin  `json:"admin_list,omitempty"`
	MemberVersion string            `json:"member_version,omitempty"`
}

// GetGroupChatResponse 获取客户群详情响应
type GetGroupChatResponse struct {
	GroupChat GroupChat `json:"group_chat"`
}

// OpenGIDToChatIDRequest 客户群opengid转换请求
type OpenGIDToChatIDRequest struct {
	OpenGID string `json:"opengid"`
}

// OpenGIDToChatIDResponse 客户群opengid转换响应
type OpenGIDToChatIDResponse struct {
	ChatID string `json:"chat_id"`
}
