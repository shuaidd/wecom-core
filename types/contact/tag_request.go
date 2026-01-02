package contact

// CreateTagRequest 创建标签请求
type CreateTagRequest struct {
	// TagName 标签名称，长度限制为32个字以内（汉字或英文字母），标签名不可与其他标签重名
	TagName string `json:"tagname"`
	// TagID 标签id，非负整型，指定此参数时新增的标签会生成对应的标签id，不指定时则以目前最大的id自增
	TagID int `json:"tagid,omitempty"`
}

// UpdateTagRequest 更新标签请求
type UpdateTagRequest struct {
	// TagID 标签ID
	TagID int `json:"tagid"`
	// TagName 标签名称，长度限制为32个字（汉字或英文字母），标签不可与其他标签重名
	TagName string `json:"tagname"`
}

// AddTagUsersRequest 增加标签成员请求
type AddTagUsersRequest struct {
	// TagID 标签ID
	TagID int `json:"tagid"`
	// UserList 企业成员ID列表，注意：userlist、partylist不能同时为空
	UserList []string `json:"userlist,omitempty"`
	// PartyList 企业部门ID列表，注意：userlist、partylist不能同时为空
	PartyList []int `json:"partylist,omitempty"`
}

// DeleteTagUsersRequest 删除标签成员请求
type DeleteTagUsersRequest struct {
	// TagID 标签ID
	TagID int `json:"tagid"`
	// UserList 企业成员ID列表，注意：userlist、partylist不能同时为空
	UserList []string `json:"userlist,omitempty"`
	// PartyList 企业部门ID列表，注意：userlist、partylist不能同时为空
	PartyList []int `json:"partylist,omitempty"`
}
