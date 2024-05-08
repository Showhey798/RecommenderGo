package appcontext

import (
	"context"
	"runtime"

	"recommender.package/pkg/applogger"
)

func SetLoggerContext(ctx context.Context, logger *applogger.Logger) context.Context {
	return context.WithValue(ctx, loggerPtr, logger)
}

func InfoMessage(ctx context.Context, msg string) {
	logger := getLoggerFromContext(ctx)
	if logger == nil {
		return
	}
	_, filename, line, _ := runtime.Caller(1)
	logMsg := applogger.LogMessage{
		SourceFile: filename,
		SourceLine: line,
		Message:    msg,
	}
	(*logger).Info(logMsg)
}

func Info(ctx context.Context, message string, content interface{}) {
	logger := getLoggerFromContext(ctx)
	if logger == nil {
		return
	}
	_, filename, line, _ := runtime.Caller(1)
	msg := applogger.LogMessage{
		SourceFile: filename,
		SourceLine: line,
		Message:    message,
		Content:    content,
	}
	(*logger).Info(msg)
}

func Warn(ctx context.Context, msg string, err error) {
	logger := getLoggerFromContext(ctx)
	if logger == nil {
		return
	}
	(*logger).Warn(msg, err)
}

func getLoggerFromContext(ctx context.Context) *applogger.Logger {
	return (ctx).Value(loggerPtr).(*applogger.Logger)
}

func Error(ctx context.Context, msg string, err error) {
	logger := getLoggerFromContext(ctx)
	if logger == nil {
		return
	}
	(*logger).Error(msg, err)
}
