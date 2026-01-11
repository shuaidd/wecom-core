package contact

import "github.com/shuaidd/wecom-core/types/common"

// ListUserIDsRequest 获取成员ID列表请求
type ListUserIDsRequest struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用不填
	Cursor string `json:"cursor,omitempty"`
	// Limit 分页，预期请求的数据量，取值范围 1 ~ 10000
	Limit int `json:"limit,omitempty"`
}

// DeptUser 用户-部门关系
type DeptUser struct {
	// UserID 用户userid，当用户在多个部门下时会有多条记录
	UserID string `json:"userid"`
	// Department 用户所属部门
	Department int `json:"department"`
}

// ListUserIDsResponse 获取成员ID列表响应
type ListUserIDsResponse struct {
	common.Response
	// NextCursor 分页游标，下次请求时填写以获取之后分页的记录。如果该字段返回空则表示已没有更多数据
	NextCursor string `json:"next_cursor,omitempty"`
	// DeptUser 用户-部门关系列表
	DeptUser []DeptUser `json:"dept_user"`
}

// ExportRequest 导出请求
type ExportRequest struct {
	// EncodingAESKey Base64编码后的加密密钥
	EncodingAESKey string `json:"encoding_aeskey"`
	// BlockSize 每块数据的数量，支持范围[10^4,10^6]，默认值为10^6
	BlockSize int `json:"block_size,omitempty"`
}

// ExportTagUserRequest 导出标签成员请求
type ExportTagUserRequest struct {
	// TagID 需要导出的标签
	TagID int `json:"tagid"`
	// EncodingAESKey Base64编码后的加密密钥
	EncodingAESKey string `json:"encoding_aeskey"`
	// BlockSize 每块数据的人员数和部门数之和，支持范围[10^4,10^6]，默认值为10^6
	BlockSize int `json:"block_size,omitempty"`
}

// ExportResponse 导出响应
type ExportResponse struct {
	common.Response
	// JobID 任务ID，可通过获取导出结果接口查询任务结果
	JobID string `json:"jobid"`
}

// GetExportResultResponse 获取导出结果响应
type GetExportResultResponse struct {
	common.Response
	// Status 任务状态：0-未处理，1-处理中，2-完成，3-异常失败
	Status int `json:"status"`
	// DataList 数据文件列表
	DataList []ExportDataFile `json:"data_list,omitempty"`
}

// ExportDataFile 导出数据文件
type ExportDataFile struct {
	// URL 数据文件下载地址，支持断点续传
	URL string `json:"url"`
	// MD5 数据文件的md5
	MD5 string `json:"md5"`
	// Size 数据文件大小
	Size int `json:"size"`
}
