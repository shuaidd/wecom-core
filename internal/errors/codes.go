package errors

// 企业微信错误码常量
const (
	// ErrCodeOK 成功
	ErrCodeOK = 0

	// ErrCodeInvalidCredential 凭证无效
	ErrCodeInvalidCredential = 40001

	// ErrCodeInvalidAccessToken access_token无效
	ErrCodeInvalidAccessToken = 40014

	// ErrCodeAccessTokenExpired access_token过期
	ErrCodeAccessTokenExpired = 42001

	// ErrCodeAPIFreqLimit API调用频率限制
	ErrCodeAPIFreqLimit = 45009

	// ErrCodeIPNotInWhitelist IP不在白名单
	ErrCodeIPNotInWhitelist = 60020

	// ErrCodeSystemBusy 系统繁忙
	ErrCodeSystemBusy = 10001

	// ErrCodeInvalidParameter 参数错误
	ErrCodeInvalidParameter = 40003

	// ErrCodeInvalidUserID 无效的UserID
	ErrCodeInvalidUserID = 40013

	// ErrCodeInvalidDepartmentID 无效的部门ID
	ErrCodeInvalidDepartmentID = 60011

	// ErrCodeUserNotFound 成员不存在
	ErrCodeUserNotFound = 60111

	// ErrCodeDepartmentNotFound 部门不存在
	ErrCodeDepartmentNotFound = 60123
)

// 预定义错误变量
var (
	// ErrInvalidCredential 凭证无效错误
	ErrInvalidCredential = New(ErrCodeInvalidCredential, "invalid credential")

	// ErrInvalidAccessToken access_token无效错误
	ErrInvalidAccessToken = New(ErrCodeInvalidAccessToken, "invalid access_token")

	// ErrAccessTokenExpired access_token过期错误
	ErrAccessTokenExpired = New(ErrCodeAccessTokenExpired, "access_token expired")

	// ErrAPIFreqLimit API调用频率限制错误
	ErrAPIFreqLimit = New(ErrCodeAPIFreqLimit, "api freq out of limit")

	// ErrIPNotInWhitelist IP不在白名单错误
	ErrIPNotInWhitelist = New(ErrCodeIPNotInWhitelist, "ip not in whitelist")

	// ErrSystemBusy 系统繁忙错误
	ErrSystemBusy = New(ErrCodeSystemBusy, "system busy")
)

// IsTokenExpired 判断是否为token过期错误
func IsTokenExpired(err error) bool {
	code := GetErrorCode(err)
	return code == ErrCodeAccessTokenExpired || code == ErrCodeInvalidAccessToken
}

// IsRateLimited 判断是否为频率限制错误
func IsRateLimited(err error) bool {
	code := GetErrorCode(err)
	return code == ErrCodeAPIFreqLimit
}

// IsSystemBusy 判断是否为系统繁忙错误
func IsSystemBusy(err error) bool {
	code := GetErrorCode(err)
	return code == ErrCodeSystemBusy
}

// IsRetriable 判断错误是否可重试
func IsRetriable(err error) bool {
	return IsTokenExpired(err) || IsRateLimited(err) || IsSystemBusy(err)
}
