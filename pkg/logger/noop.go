package logger

// NoopLogger 空日志记录器（不输出任何日志）
type NoopLogger struct{}

// NewNoopLogger 创建空日志记录器
func NewNoopLogger() Logger {
	return &NoopLogger{}
}

// Debug 调试日志（无操作）
func (l *NoopLogger) Debug(msg string, fields ...Field) {}

// Info 信息日志（无操作）
func (l *NoopLogger) Info(msg string, fields ...Field) {}

// Warn 警告日志（无操作）
func (l *NoopLogger) Warn(msg string, fields ...Field) {}

// Error 错误日志（无操作）
func (l *NoopLogger) Error(msg string, fields ...Field) {}
