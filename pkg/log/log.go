package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Sugar wraps the Logger to provide a more ergonomic, but slightly slower,
// API. Sugaring a Logger is quite inexpensive, so it's reasonable for a
// single application to use both Loggers and SugaredLoggers, converting
// between them on the boundaries of performance-sensitive code.
func Sugar() *zap.SugaredLogger {
	return Logger.Sugar()
}

// Named adds a new path segment to the Logger's name. Segments are joined by
// periods. By default, Loggers are unnamed.
func Named(s string) *zap.Logger {
	return Logger.Named(s)
}

// WithOptions clones the current Logger, applies the supplied Options, and
// returns the resulting Logger. It's safe to use concurrently.
func WithOptions(opts ...zap.Option) *zap.Logger {
	return Logger.WithOptions(opts...)
}

// With creates a child Logger and adds structured context to it. Fields added
// to the child don't affect the parent, and vice versa.
func With(fields ...zap.Field) *zap.Logger {
	return Logger.With(fields...)
}

// Check returns a CheckedEntry if logging a message at the specified level
// is enabled. It's a completely optional optimization; in high-performance
// applications, Check can help avoid allocating a slice to hold fields.
func Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	return Logger.Check(lvl, msg)
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the Logger.
func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the Logger.
func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the Logger.
func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the Logger.
func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

// DPanic logs a message at DPanicLevel. The message includes any fields
// passed at the log site, as well as any fields accumulated on the Logger.
//
// If the Logger is in development mode, it then panics (DPanic means
// "development panic"). This is useful for catching errors that are
// recoverable, but shouldn't ever happen.
func DPanic(msg string, fields ...zap.Field) {
	Logger.DPanic(msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the Logger.
//
// The Logger then panics, even if logging at PanicLevel is disabled.
func Panic(msg string, fields ...zap.Field) {
	Logger.Panic(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the Logger.
//
// The Logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}

// Sync calls the underlying Core's Sync method, flushing any buffered log
// entries. Applications should take care to call Sync before exiting.
func Sync() error {
	return Logger.Sync()
}

// Core returns the Logger's underlying zapcore.Core.
func Core() zapcore.Core {
	return Logger.Core()
}
