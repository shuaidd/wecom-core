package updown

import "github.com/shuaidd/wecom-core/types/common"

// FailContactInfo 导入失败的联系人信息
type FailContactInfo struct {
	Mobile  string `json:"mobile"`  // 导入失败的联系人手机号
	ErrCode int    `json:"errcode"` // 导入失败的联系人错误码
	ErrMsg  string `json:"errmsg"`  // 导入失败的联系人错误码描述
}

// FailCorpInfo 导入失败的企业信息
type FailCorpInfo struct {
	CorpName        string            `json:"corp_name"`         // 自定义企业名称
	CustomID        string            `json:"custom_id"`         // 自定义企业id
	ErrCode         int               `json:"errcode"`           // 该企业导入操作的结果错误码
	ErrMsg          string            `json:"errmsg"`            // 该企业导入操作的结果错误码描述
	ContactInfoList []FailContactInfo `json:"contact_info_list"` // 导入失败的联系人结果
}

// TaskResult 任务结果
type TaskResult struct {
	ChainID      string         `json:"chain_id"`      // 上下游id
	ImportStatus int            `json:"import_status"` // 导入状态。1:全部企业导入成功，2:部分企业导入成功，3:全部企业导入失败
	FailList     []FailCorpInfo `json:"fail_list"`     // 导入失败结果列表
}

// GetTaskResultResponse 获取异步任务结果响应
type GetTaskResultResponse struct {
	common.Response
	Status int         `json:"status"` // 任务状态，1表示任务开始，2表示任务进行中，3表示任务已完成
	Result *TaskResult `json:"result"` // 详细的处理结果
}
