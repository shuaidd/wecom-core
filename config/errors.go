package config

import "errors"

var (
	// ErrMissingCorpID 缺少企业ID
	ErrMissingCorpID = errors.New("corpID is required")

	// ErrMissingCorpSecret 缺少应用凭证密钥
	ErrMissingCorpSecret = errors.New("corpSecret is required")

	// ErrInvalidTimeout 无效的超时时间
	ErrInvalidTimeout = errors.New("timeout must be greater than 0")

	// ErrInvalidMaxRetries 无效的重试次数
	ErrInvalidMaxRetries = errors.New("maxRetries must be greater than or equal to 0")
)
