package logger

// Logger 日志记录器接口
type Logger interface {
	// Debug 调试日志
	Debug(msg string, fields ...Field)

	// Info 信息日志
	Info(msg string, fields ...Field)

	// Warn 警告日志
	Warn(msg string, fields ...Field)

	// Error 错误日志
	Error(msg string, fields ...Field)
}

// Field 日志字段
type Field struct {
	Key   string
	Value interface{}
}

// F 创建日志字段的快捷函数
func F(key string, value interface{}) Field {
	return Field{
		Key:   key,
		Value: value,
	}
}
