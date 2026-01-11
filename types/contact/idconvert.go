package contact

import "github.com/shuaidd/wecom-core/types/common"

// ConvertTmpExternalUserIDRequest tmp_external_userid转换请求
type ConvertTmpExternalUserIDRequest struct {
	BusinessType          int      `json:"business_type"`
	UserType              int      `json:"user_type"`
	TmpExternalUserIDList []string `json:"tmp_external_userid_list"`
}

// ConvertResult 转换结果
type ConvertResult struct {
	TmpExternalUserID string `json:"tmp_external_userid"`
	ExternalUserID    string `json:"external_userid,omitempty"`
	CorpID            string `json:"corpid,omitempty"`
	UserID            string `json:"userid,omitempty"`
}

// ConvertTmpExternalUserIDResponse tmp_external_userid转换响应
type ConvertTmpExternalUserIDResponse struct {
	common.Response
	Results                      []ConvertResult `json:"results,omitempty"`
	InvalidTmpExternalUserIDList []string        `json:"invalid_tmp_external_userid_list,omitempty"`
}
