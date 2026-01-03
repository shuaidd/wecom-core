package security

// GetServerDomainIPResponse 获取企业微信域名IP信息响应
type GetServerDomainIPResponse struct {
	ErrCode    int          `json:"errcode"`
	ErrMsg     string       `json:"errmsg"`
	DomainList []DomainInfo `json:"domain_list"` // 域名列表
	IPList     []IPInfo     `json:"ip_list"`     // IP列表
}

// DomainInfo 域名信息
type DomainInfo struct {
	Domain          string `json:"domain"`           // 域名
	UniversalDomain string `json:"universal_domian"` // 泛域名
	Protocol        string `json:"protocol"`         // 协议 如TCP UDP
	Port            []int  `json:"port"`             // 端口号列表
	IsNecessary     int    `json:"is_necessary"`     // 是否必要，0-否 1-是
	Description     string `json:"description"`      // 域名涉及到的功能的描述信息
}

// IPInfo IP信息
type IPInfo struct {
	IP          string `json:"ip"`           // ip地址
	Protocol    string `json:"protocol"`     // 协议 如TCP UDP
	Port        []int  `json:"port"`         // 端口号列表
	IsNecessary int    `json:"is_necessary"` // 是否必要，0-否 1-是
	Description string `json:"description"`  // IP涉及到的功能的描述信息
}
