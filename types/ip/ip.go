package ip

import "github.com/shuaidd/wecom-core/types/common"

// GetCallbackIPResponse 获取企业微信回调IP段响应
type GetCallbackIPResponse struct {
	common.Response
	// IPList 企业微信回调服务器IP段
	IPList []string `json:"ip_list"`
}

// GetAPIDomainIPResponse 获取企业微信接口IP段响应
type GetAPIDomainIPResponse struct {
	common.Response
	// IPList 企业微信API域名IP段
	IPList []string `json:"ip_list"`
}
