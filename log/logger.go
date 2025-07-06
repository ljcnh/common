package log

import (
	"context"
	"log/slog"
	"os"

	"go.uber.org/zap"
)

// Logger 定义日志接口，支持 slog 和 zap 的核心方法
type Logger interface {
	Debug(ctx context.Context, msg string, args ...interface{})
	Info(ctx context.Context, msg string, args ...interface{})
	Warn(ctx context.Context, msg string, args ...interface{})
	Error(ctx context.Context, msg string, args ...interface{})
	Sync() error
}

// SlogLogger slog 实现的日志记录器
type SlogLogger struct {
	logger *slog.Logger
}

// NewSlogLogger 创建一个新的 slog 日志记录器
func NewSlogLogger() *SlogLogger {
	return &SlogLogger{
		logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})),
	}
}

func (l *SlogLogger) Debug(ctx context.Context, msg string, args ...interface{}) {
	l.logger.DebugContext(ctx, msg, args...)
}

func (l *SlogLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	l.logger.InfoContext(ctx, msg, args...)
}

func (l *SlogLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	l.logger.WarnContext(ctx, msg, args...)
}

func (l *SlogLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	l.logger.ErrorContext(ctx, msg, args...)
}

func (l *SlogLogger) Sync() error {
	return nil
}

// ZapLogger zap 实现的日志记录器
type ZapLogger struct {
	logger *zap.Logger
}

// NewZapLogger 创建一个新的 zap 日志记录器
func NewZapLogger() (*ZapLogger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &ZapLogger{
		logger: logger,
	}, nil
}

func (l *ZapLogger) Debug(ctx context.Context, msg string, args ...interface{}) {
	fields := convertArgsToZapFields(args)
	l.logger.Debug(msg, fields...)
}

func (l *ZapLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	fields := convertArgsToZapFields(args)
	l.logger.Info(msg, fields...)
}

func (l *ZapLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	fields := convertArgsToZapFields(args)
	l.logger.Warn(msg, fields...)
}

func (l *ZapLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	fields := convertArgsToZapFields(args)
	l.logger.Error(msg, fields...)
}

func (l *ZapLogger) Sync() error {
	return l.logger.Sync()
}

// convertArgsToZapFields 将通用参数转换为 zap 的字段
func convertArgsToZapFields(args []interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(args)/2)
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			if key, ok := args[i].(string); ok {
				fields = append(fields, zap.Any(key, args[i+1]))
			}
		}
	}
	return fields
}
