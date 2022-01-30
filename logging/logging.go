package logging

import (
	"os"

	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

// func IsDebugOrHigher() bool {
// 	lvl := log.GetLevel()
// 	return lvl == log.DebugLevel || lvl == log.TraceLevel
// }

func IsDebugOrHigher(logger *log.Logger) bool {
	lvl := logger.GetLevel()
	return lvl == log.DebugLevel || lvl == log.TraceLevel
}

func GetDefaultLogger(debug bool) *log.Logger {
	level := log.InfoLevel
	if debug {
		level = log.DebugLevel
	}

	logger := &log.Logger{
		Out:   os.Stderr,
		Level: level,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "%time% [%lvl%] harness-go-sdk %msg%\n",
		},
	}

	return logger
}

type LeveledLogger struct {
	Logger *log.Logger
}

func (l *LeveledLogger) Error(msg string, keysAndValues ...interface{}) {
	l.Logger.Errorf(msg, keysAndValues...)
}

func (l *LeveledLogger) Info(msg string, keysAndValues ...interface{}) {
	l.Logger.Infof(msg, keysAndValues...)
}

func (l *LeveledLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.Logger.Debugf(msg, keysAndValues...)
}

func (l *LeveledLogger) Warn(msg string, keysAndValues ...interface{}) {
	l.Logger.Warnf(msg, keysAndValues...)
}
