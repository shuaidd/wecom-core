package errors

import (
	"errors"
	"fmt"
)

// Error 企业微信API错误
type Error struct {
	// Code 错误码
	Code int
	// Message 错误消息
	Message string
	// Cause 原始错误
	Cause error
}

// New 创建企业微信错误
func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// Wrap 包装原始错误
func Wrap(err error, code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Cause:   err,
	}
}

// Error 实现 error 接口
func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("wecom error [%d]: %s (cause: %v)", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("wecom error [%d]: %s", e.Code, e.Message)
}

// Unwrap 支持 errors.Unwrap
func (e *Error) Unwrap() error {
	return e.Cause
}

// Is 判断错误类型
func (e *Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

// IsWecomError 判断是否为企业微信错误
func IsWecomError(err error) bool {
	var e *Error
	return errors.As(err, &e)
}

// GetErrorCode 获取错误码，如果不是企业微信错误则返回 0
func GetErrorCode(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Code
	}
	return 0
}
