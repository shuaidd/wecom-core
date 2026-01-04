package media

// MediaType 媒体文件类型
type MediaType string

const (
	// MediaTypeImage 图片
	MediaTypeImage MediaType = "image"
	// MediaTypeVoice 语音
	MediaTypeVoice MediaType = "voice"
	// MediaTypeVideo 视频
	MediaTypeVideo MediaType = "video"
	// MediaTypeFile 普通文件
	MediaTypeFile MediaType = "file"
)

// UploadImageResponse 上传图片响应
type UploadImageResponse struct {
	// URL 图片URL,永久有效
	URL string `json:"url"`
}

// UploadMediaResponse 上传临时素材响应
type UploadMediaResponse struct {
	// Type 媒体文件类型
	Type string `json:"type"`
	// MediaID 媒体文件上传后获取的唯一标识,3天内有效
	MediaID string `json:"media_id"`
	// CreatedAt 媒体文件上传时间戳
	CreatedAt string `json:"created_at"`
}

// UploadByURLRequest 异步上传临时素材请求
type UploadByURLRequest struct {
	// Scene 场景值。1-客户联系入群欢迎语素材
	Scene int `json:"scene"`
	// Type 媒体文件类型。目前仅支持video-视频，file-普通文件
	Type string `json:"type"`
	// Filename 文件名，标识文件展示的名称
	Filename string `json:"filename"`
	// URL 文件cdn url
	URL string `json:"url"`
	// MD5 文件md5
	MD5 string `json:"md5"`
}

// UploadByURLResponse 异步上传临时素材响应
type UploadByURLResponse struct {
	// JobID 任务id。可通过此jobid查询结果
	JobID string `json:"jobid"`
}

// GetUploadByURLResultRequest 查询异步任务结果请求
type GetUploadByURLResultRequest struct {
	// JobID 任务id。最长为128字节，60分钟内有效
	JobID string `json:"jobid"`
}

// UploadTaskStatus 上传任务状态
type UploadTaskStatus int

const (
	// UploadTaskStatusProcessing 处理中
	UploadTaskStatusProcessing UploadTaskStatus = 1
	// UploadTaskStatusCompleted 完成
	UploadTaskStatusCompleted UploadTaskStatus = 2
	// UploadTaskStatusFailed 异常失败
	UploadTaskStatusFailed UploadTaskStatus = 3
)

// UploadTaskDetail 上传任务详情
type UploadTaskDetail struct {
	// ErrCode 任务失败返回码。当status为3时返回非0，其他返回0
	ErrCode int `json:"errcode"`
	// ErrMsg 任务失败错误码描述
	ErrMsg string `json:"errmsg"`
	// MediaID 媒体文件上传后获取的唯一标识，3天内有效。当status为2时返回
	MediaID string `json:"media_id,omitempty"`
	// CreatedAt 媒体文件创建的时间戳。当status为2时返回
	CreatedAt string `json:"created_at,omitempty"`
}

// GetUploadByURLResultResponse 查询异步任务结果响应
type GetUploadByURLResultResponse struct {
	// Status 任务状态。1-处理中，2-完成，3-异常失败
	Status UploadTaskStatus `json:"status"`
	// Detail 结果明细
	Detail UploadTaskDetail `json:"detail"`
}
