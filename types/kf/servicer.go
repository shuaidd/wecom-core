package kf

import "github.com/shuaidd/wecom-core/types/common"

// AddServicerRequest 添加接待人员请求
type AddServicerRequest struct {
	OpenKfID         string   `json:"open_kfid"`                    // 客服账号ID
	UserIDList       []string `json:"userid_list,omitempty"`        // 接待人员userid列表,可填充个数:0~100
	DepartmentIDList []uint64 `json:"department_id_list,omitempty"` // 接待人员部门id列表,可填充个数:0~20
}

// ServicerOperationResult 接待人员操作结果
type ServicerOperationResult struct {
	UserID       string `json:"userid,omitempty"`        // 接待人员的userid
	DepartmentID uint64 `json:"department_id,omitempty"` // 接待人员部门的id
	ErrCode      int    `json:"errcode"`                 // 该userid的操作结果
	ErrMsg       string `json:"errmsg"`                  // 结果信息
}

// AddServicerResponse 添加接待人员响应
type AddServicerResponse struct {
	common.Response
	ResultList []ServicerOperationResult `json:"result_list"` // 操作结果
}

// DeleteServicerRequest 删除接待人员请求
type DeleteServicerRequest struct {
	OpenKfID         string   `json:"open_kfid"`                    // 客服账号ID
	UserIDList       []string `json:"userid_list,omitempty"`        // 接待人员userid列表,可填充个数:0~100
	DepartmentIDList []uint64 `json:"department_id_list,omitempty"` // 接待人员部门id列表,可填充个数:0~100
}

// DeleteServicerResponse 删除接待人员响应
type DeleteServicerResponse struct {
	common.Response
	ResultList []ServicerOperationResult `json:"result_list"` // 操作结果
}

// ListServicerRequest 获取接待人员列表请求
type ListServicerRequest struct {
	OpenKfID string `json:"open_kfid"` // 客服账号ID
}

// Servicer 接待人员信息
type Servicer struct {
	UserID       string `json:"userid,omitempty"`        // 接待人员的userid
	DepartmentID uint64 `json:"department_id,omitempty"` // 接待人员部门的id
	Status       uint   `json:"status,omitempty"`        // 接待人员的接待状态,0:接待中,1:停止接待
	StopType     uint   `json:"stop_type,omitempty"`     // 接待状态为停止接待的子类型,0:停止接待,1:暂时挂起
}

// ListServicerResponse 获取接待人员列表响应
type ListServicerResponse struct {
	common.Response
	ServicerList []Servicer `json:"servicer_list"` // 客服账号的接待人员列表
}
