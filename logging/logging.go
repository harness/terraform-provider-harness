package logging

import (
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

func IsDebugOrHigher(logger *log.Logger) bool {
	lvl := logger.GetLevel()
	return lvl == log.DebugLevel || lvl == log.TraceLevel
}

// NewLogger creates a new logger with default settings.
func NewLogger() *log.Logger {
	logger := log.New()
	logger.Formatter = &easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "%time% [%lvl%] harness-go-sdk %msg%\n",
	}
	return logger
}

type LeveledLogger interface {
	Error(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Debug(msg string, keysAndValues ...interface{})
	IsDebugOrHigher() bool
}

type Logger struct {
	Logger *log.Logger
}

func (l *Logger) Error(msg string, keysAndValues ...interface{}) {
	l.Logger.Errorf(msg, keysAndValues...)
}

func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	l.Logger.Infof(msg, keysAndValues...)
}

func (l *Logger) Debug(msg string, keysAndValues ...interface{}) {
	l.Logger.Debugf(msg, keysAndValues...)
}

func (l *Logger) Warn(msg string, keysAndValues ...interface{}) {
	l.Logger.Warnf(msg, keysAndValues...)
}
