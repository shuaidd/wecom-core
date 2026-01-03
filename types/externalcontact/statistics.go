package externalcontact

// GroupChatStatisticRequest 获取群聊数据统计请求（按群主聚合）
type GroupChatStatisticRequest struct {
	DayBeginTime int64        `json:"day_begin_time"`
	DayEndTime   int64        `json:"day_end_time,omitempty"`
	OwnerFilter  *OwnerFilter `json:"owner_filter"`
	OrderBy      int          `json:"order_by,omitempty"`
	OrderAsc     int          `json:"order_asc,omitempty"`
	Offset       int          `json:"offset,omitempty"`
	Limit        int          `json:"limit,omitempty"`
}

// GroupChatStatisticData 群聊统计数据
type GroupChatStatisticData struct {
	NewChatCnt            int `json:"new_chat_cnt"`
	ChatTotal             int `json:"chat_total"`
	ChatHasMsg            int `json:"chat_has_msg"`
	NewMemberCnt          int `json:"new_member_cnt"`
	MemberTotal           int `json:"member_total"`
	MemberHasMsg          int `json:"member_has_msg"`
	MsgTotal              int `json:"msg_total"`
	MigrateTraineeChatCnt int `json:"migrate_trainee_chat_cnt,omitempty"`
}

// GroupChatStatisticItem 群聊统计项（按群主聚合）
type GroupChatStatisticItem struct {
	Owner string                 `json:"owner"`
	Data  GroupChatStatisticData `json:"data"`
}

// GroupChatStatisticResponse 获取群聊数据统计响应（按群主聚合）
type GroupChatStatisticResponse struct {
	Total      int                      `json:"total"`
	NextOffset int                      `json:"next_offset"`
	Items      []GroupChatStatisticItem `json:"items"`
}

// GroupChatStatisticGroupByDayRequest 获取群聊数据统计请求（按自然日聚合）
type GroupChatStatisticGroupByDayRequest struct {
	DayBeginTime int64        `json:"day_begin_time"`
	DayEndTime   int64        `json:"day_end_time,omitempty"`
	OwnerFilter  *OwnerFilter `json:"owner_filter"`
}

// GroupChatStatisticGroupByDayItem 群聊统计项（按自然日聚合）
type GroupChatStatisticGroupByDayItem struct {
	StatTime int64                  `json:"stat_time"`
	Data     GroupChatStatisticData `json:"data"`
}

// GroupChatStatisticGroupByDayResponse 获取群聊数据统计响应（按自然日聚合）
type GroupChatStatisticGroupByDayResponse struct {
	Items []GroupChatStatisticGroupByDayItem `json:"items"`
}

// GetUserBehaviorDataRequest 获取联系客户统计数据请求
type GetUserBehaviorDataRequest struct {
	UserID    []string `json:"userid,omitempty"`
	PartyID   []int    `json:"partyid,omitempty"`
	StartTime int64    `json:"start_time"`
	EndTime   int64    `json:"end_time"`
}

// BehaviorData 联系客户行为数据
type BehaviorData struct {
	StatTime            int64   `json:"stat_time"`
	ChatCnt             int     `json:"chat_cnt"`
	MessageCnt          int     `json:"message_cnt"`
	ReplyPercentage     float64 `json:"reply_percentage,omitempty"`
	AvgReplyTime        int     `json:"avg_reply_time,omitempty"`
	NegativeFeedbackCnt int     `json:"negative_feedback_cnt"`
	NewApplyCnt         int     `json:"new_apply_cnt"`
	NewContactCnt       int     `json:"new_contact_cnt"`
}

// GetUserBehaviorDataResponse 获取联系客户统计数据响应
type GetUserBehaviorDataResponse struct {
	BehaviorData []BehaviorData `json:"behavior_data"`
}
