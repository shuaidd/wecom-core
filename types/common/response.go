package common

// Response 通用响应
type Response struct {
	// ErrCode 错误码，0表示成功
	ErrCode int `json:"errcode"`
	// ErrMsg 错误消息
	ErrMsg string `json:"errmsg"`
}
