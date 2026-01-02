package updown

import "github.com/shuaidd/wecom-core/types/common"

// ContactInfo 联系人信息
type ContactInfo struct {
	Name         string `json:"name"`          // 上下游联系人姓名
	IdentityType int    `json:"identity_type"` // 联系人身份类型。1:成员，2:负责人
	Mobile       string `json:"mobile"`        // 手机号
	UserCustomID string `json:"user_custom_id,omitempty"` // 上下游用户自定义id
}

// ChainContact 上下游联系人
type ChainContact struct {
	CorpName        string        `json:"corp_name"`                 // 上下游企业名称
	GroupPath       string        `json:"group_path,omitempty"`      // 导入后企业所在分组
	CustomID        string        `json:"custom_id,omitempty"`       // 上下游企业自定义id
	ContactInfoList []ContactInfo `json:"contact_info_list"`         // 上下游联系人信息列表
}

// ImportChainContactRequest 批量导入上下游联系人请求
type ImportChainContactRequest struct {
	ChainID     string         `json:"chain_id"`     // 上下游id
	ContactList []ChainContact `json:"contact_list"` // 上下游联系人列表
}

// ImportChainContactResponse 批量导入上下游联系人响应
type ImportChainContactResponse struct {
	common.Response
	JobID string `json:"jobid"` // 异步任务id
}
