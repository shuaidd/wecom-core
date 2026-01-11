package wedoc

// CreateDocRequest 新建文档请求
type CreateDocRequest struct {
	SpaceID    string   `json:"spaceid,omitempty"`     // 空间spaceid
	FatherID   string   `json:"fatherid,omitempty"`    // 父目录fileid, 在根目录时为空间spaceid
	DocType    uint32   `json:"doc_type"`              // 文档类型, 3:文档 4:表格 10:智能表格
	DocName    string   `json:"doc_name"`              // 文档名字（注意：文件名最多填255个字符, 超过255个字符会被截断）
	AdminUsers []string `json:"admin_users,omitempty"` // 文档管理员userid
}

// CreateDocResponse 新建文档响应
type CreateDocResponse struct {
	URL   string `json:"url"`   // 新建文档的访问链接
	DocID string `json:"docid"` // 新建文档的docid
}

// GetDocBaseInfoRequest 获取文档基础信息请求
type GetDocBaseInfoRequest struct {
	DocID string `json:"docid"` // 文档docid
}

// DocBaseInfo 文档基础信息
type DocBaseInfo struct {
	DocID      string `json:"docid"`       // 文档docid
	DocName    string `json:"doc_name"`    // 文档名字
	CreateTime uint64 `json:"create_time"` // 文档创建时间
	ModifyTime uint64 `json:"modify_time"` // 文档最后修改时间
	DocType    uint32 `json:"doc_type"`    // 3: 文档 4: 表格 10:智能表格
}

// GetDocBaseInfoResponse 获取文档基础信息响应
type GetDocBaseInfoResponse struct {
	DocBaseInfo DocBaseInfo `json:"doc_base_info"` // 文档基础信息
}

// DeleteDocRequest 删除文档请求
type DeleteDocRequest struct {
	DocID  string `json:"docid,omitempty"`  // 文档docid（docid、formid只能填其中一个）
	FormID string `json:"formid,omitempty"` // 收集表id（docid、formid只能填其中一个）
}

// RenameDocRequest 重命名文档请求
type RenameDocRequest struct {
	DocID   string `json:"docid,omitempty"`  // 文档docid（docid、formid只能填其中一个）
	FormID  string `json:"formid,omitempty"` // 收集表id（docid、formid只能填其中一个）
	NewName string `json:"new_name"`         // 重命名后的文档名
}

// ShareDocRequest 分享文档请求
type ShareDocRequest struct {
	DocID  string `json:"docid,omitempty"`  // 文档id（docid、formid只能填其中一个）
	FormID string `json:"formid,omitempty"` // 收集表id（docid、formid只能填其中一个）
}

// ShareDocResponse 分享文档响应
type ShareDocResponse struct {
	ShareURL string `json:"share_url"` // 文档分享链接
}

// GetDocAuthRequest 获取文档权限信息请求
type GetDocAuthRequest struct {
	DocID string `json:"docid"` // 文档id
}

// GetDocAuthResponse 获取文档权限信息响应
type GetDocAuthResponse struct {
	AccessRule    AccessRule    `json:"access_rule"`     // 文档的查看规则
	SecureSetting SecureSetting `json:"secure_setting"`  // 文档安全设置
	DocMemberList []DocMember   `json:"doc_member_list"` // 文档通知范围及权限列表
	CoAuthList    []CoAuth      `json:"co_auth_list"`    // 文档查看权限特定部门列表
}

// AccessRule 文档查看规则
type AccessRule struct {
	EnableCorpInternal             bool `json:"enable_corp_internal"`                // 是否允许企业内成员浏览文档
	CorpInternalAuth               int  `json:"corp_internal_auth"`                  // 企业内成员主动查看文档后获得的权限类型 1:只读 2:读写
	EnableCorpExternal             bool `json:"enable_corp_external"`                // 是否允许企业外成员浏览文档
	CorpExternalAuth               int  `json:"corp_external_auth"`                  // 企业外成员主动查看文档后获得的权限类型 1:只读 2:读写
	CorpInternalApproveOnlyByAdmin bool `json:"corp_internal_approve_only_by_admin"` // 企业内成员浏览文档是否必须由管理员审批
	CorpExternalApproveOnlyByAdmin bool `json:"corp_external_approve_only_by_admin"` // 企业外成员浏览文档是否必须由管理员审批
	BanShareExternal               bool `json:"ban_share_external"`                  // 是否禁止文档分享到企业外
}

// SecureSetting 文档安全设置
type SecureSetting struct {
	EnableReadonlyCopy    bool      `json:"enable_readonly_copy"`    // 仅浏览权限的成员是否允许导出、复制、打印
	Watermark             Watermark `json:"watermark"`               // 文档水印设置
	EnableReadonlyComment bool      `json:"enable_readonly_comment"` // 是否允许只读成员评论
}

// Watermark 文档水印
type Watermark struct {
	MarginType      int    `json:"margin_type"`       // 水印密度 1:稀疏 2:紧密
	ShowVisitorName bool   `json:"show_visitor_name"` // 是否展示访问者名字
	ShowText        bool   `json:"show_text"`         // 是否展示水印文字
	Text            string `json:"text"`              // 水印文字
}

// DocMember 文档通知范围成员
type DocMember struct {
	Type              int    `json:"type"`                          // 文档通知范围成员种类 1:user
	UserID            string `json:"userid,omitempty"`              // 企业成员的userid
	TmpExternalUserID string `json:"tmp_external_userid,omitempty"` // 外部用户临时id
	Auth              int    `json:"auth"`                          // 该文档通知范围成员的权限 1:只读 2:读写 7:管理员
}

// CoAuth 文档查看权限特定部门
type CoAuth struct {
	Type         int    `json:"type"`         // 特定部门列表 2:部门
	DepartmentID uint64 `json:"departmentid"` // 特定部门id
	Auth         int    `json:"auth"`         // 权限类型 1:只读 2:读写
}

// ModDocSaftySettingRequest 修改文档安全设置请求
type ModDocSaftySettingRequest struct {
	DocID              string     `json:"docid"`                          // 操作的文档id
	EnableReadonlyCopy *bool      `json:"enable_readonly_copy,omitempty"` // 是否允许只读成员复制、下载文档
	Watermark          *Watermark `json:"watermark,omitempty"`            // 水印设置
}

// ModDocJoinRuleRequest 修改文档查看规则请求
type ModDocJoinRuleRequest struct {
	DocID                          string   `json:"docid"`                                         // 操作的docid
	EnableCorpInternal             *bool    `json:"enable_corp_internal,omitempty"`                // 是否允许企业内成员浏览文档
	CorpInternalAuth               *int     `json:"corp_internal_auth,omitempty"`                  // 企业内成员主动查看文档后获得的权限类型 1:只读 2:读写
	EnableCorpExternal             *bool    `json:"enable_corp_external,omitempty"`                // 是否允许企业外成员浏览文档
	CorpExternalAuth               *int     `json:"corp_external_auth,omitempty"`                  // 企业外成员主浏览文档后获得的权限类型 1:只读 2:读写
	CorpInternalApproveOnlyByAdmin *bool    `json:"corp_internal_approve_only_by_admin,omitempty"` // 企业内成员加入文档是否必须由管理员审批
	CorpExternalApproveOnlyByAdmin *bool    `json:"corp_external_approve_only_by_admin,omitempty"` // 企业外成员加入文档是否必须由管理员审批
	BanShareExternal               *bool    `json:"ban_share_external,omitempty"`                  // 是否禁止文档分享到企业外
	UpdateCoAuthList               *bool    `json:"update_co_auth_list,omitempty"`                 // 是否更新文档查看权限的特定部门
	CoAuthList                     []CoAuth `json:"co_auth_list,omitempty"`                        // 需要更新文档查看权限特定部门时, 覆盖之前部门
}

// ModDocMemberRequest 修改文档通知范围及权限请求
type ModDocMemberRequest struct {
	DocID                string      `json:"docid"`                             // 操作的文档id
	UpdateFileMemberList []DocMember `json:"update_file_member_list,omitempty"` // 更新文档通知范围的列表, 批次大小最大100
	DelFileMemberList    []DocMember `json:"del_file_member_list,omitempty"`    // 删除的文档通知范围列表，批次大小最大100
}
