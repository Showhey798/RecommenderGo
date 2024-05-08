package applogger

import (
	"log/slog"
	"os"
	"runtime"
	"strconv"
)

type Logger interface {
	InfoMessage(msg string)
	Info(LogMessage)
	Warn(string, error)
	Error(string, error)
}

type AppLogger struct {
	logger *slog.Logger
}

type LogMessage struct {
	SourceFile string
	SourceLine int
	Message    string
	Content    interface{}
}

func (g *AppLogger) InfoMessage(msg string) {
	_, filename, line, _ := runtime.Caller(1)
	g.logger.Info(msg, "filename", filename, "line", strconv.Itoa(line))
}

func (g *AppLogger) Info(msg LogMessage) {
	g.logger.Info(msg.Message, "filename", msg.SourceFile, "line", msg.SourceLine, slog.Any("content", msg.Content))
}

func (g *AppLogger) Warn(msg string, err error) {
	_, filename, line, _ := runtime.Caller(1)

	if err != nil {
		g.logger.Warn(msg, "filename", filename, "line", strconv.Itoa(line), "err", err)
	} else {
		g.logger.Warn(msg, "filename", filename, "line", strconv.Itoa(line))
	}
}

func (g *AppLogger) Error(msg string, err error) {
	_, filename, line, _ := runtime.Caller(1)

	if err != nil {
		g.logger.Error(msg, "filename", filename, "line", strconv.Itoa(line), "err", err)
	} else {
		g.logger.Error(msg, "filename", filename, "line", strconv.Itoa(line))
	}
}

func NewLogMessage(msg string, obj interface{}) LogMessage {
	_, filename, line, _ := runtime.Caller(1)
	return LogMessage{
		SourceFile: filename,
		SourceLine: line,
		Message:    msg,
		Content:    obj,
	}
}

// TODO: Singleton化する
func New() Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return &AppLogger{logger: logger}
}
