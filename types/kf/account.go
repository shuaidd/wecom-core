package kf

import "github.com/shuaidd/wecom-core/types/common"

// AddAccountRequest 添加客服账号请求
type AddAccountRequest struct {
	Name    string `json:"name"`     // 客服名称，不多于16个字符
	MediaID string `json:"media_id"` // 客服头像临时素材，不多于128个字节
}

// AddAccountResponse 添加客服账号响应
type AddAccountResponse struct {
	common.Response
	OpenKfID string `json:"open_kfid"` // 新创建的客服账号ID
}

// DeleteAccountRequest 删除客服账号请求
type DeleteAccountRequest struct {
	OpenKfID string `json:"open_kfid"` // 客服账号ID，不多于64字节
}

// DeleteAccountResponse 删除客服账号响应
type DeleteAccountResponse struct {
	common.Response
}

// UpdateAccountRequest 修改客服账号请求
type UpdateAccountRequest struct {
	OpenKfID string `json:"open_kfid"`         // 要修改的客服账号ID，不多于64字节
	Name     string `json:"name,omitempty"`    // 新的客服名称，不多于16个字符
	MediaID  string `json:"media_id,omitempty"` // 新的客服头像临时素材，不多于128个字节
}

// UpdateAccountResponse 修改客服账号响应
type UpdateAccountResponse struct {
	common.Response
}

// ListAccountRequest 获取客服账号列表请求
type ListAccountRequest struct {
	Offset uint32 `json:"offset,omitempty"` // 分页，偏移量，默认为0
	Limit  uint32 `json:"limit,omitempty"`  // 分页，预期请求的数据量，默认为100，取值范围 1 ~ 100
}

// Account 客服账号信息
type Account struct {
	OpenKfID        string `json:"open_kfid"`         // 客服账号ID
	Name            string `json:"name"`              // 客服名称
	Avatar          string `json:"avatar"`            // 客服头像URL
	ManagePrivilege bool   `json:"manage_privilege"`  // 当前调用接口的应用身份，是否有该客服账号的管理权限
}

// ListAccountResponse 获取客服账号列表响应
type ListAccountResponse struct {
	common.Response
	AccountList []Account `json:"account_list"` // 账号信息列表
}

// AddContactWayRequest 获取客服账号链接请求
type AddContactWayRequest struct {
	OpenKfID string `json:"open_kfid"`       // 客服账号ID
	Scene    string `json:"scene,omitempty"` // 场景值，字符串类型，不多于32字节，字符串取值范围(正则表达式)：[0-9a-zA-Z_-]*
}

// AddContactWayResponse 获取客服账号链接响应
type AddContactWayResponse struct {
	common.Response
	URL string `json:"url"` // 客服链接，开发者可将该链接嵌入到H5页面中
}
