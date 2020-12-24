package log

import (
	"fmt"
	"time"
)

func New(module string, level Level, appenders ...Appender) *Logger {
	return &Logger{module: module, level: level, appenders: appenders}
}

type Logger struct {
	module    string
	level     Level
	appenders []Appender
}

func (l *Logger) New(module string) *Logger {
	return New(module, l.level, l.appenders...)
}

func (l *Logger) SetModule(module string) *Logger {
	l.module = module
	return l
}

func (l *Logger) SetLevel(level Level) *Logger {
	l.level = level
	return l
}

func (l *Logger) SetAppenders(appenders ...Appender) *Logger {
	l.appenders = appenders
	return l
}

func (l *Logger) Append(level Level, message string, data ...map[string]interface{}) {
	if level < l.level {
		return
	}
	e := Event{
		Module:  l.module,
		Level:   level,
		Message: message,
		Time:    time.Now(),
	}
	if len(data) > 0 {
		e.Data = data[0]
	}
	for _, appender := range l.appenders {
		_ = appender.Append(e)
	}
	return
}

func (l *Logger) Appendf(level Level, format string, args ...interface{}) {
	l.Append(level, fmt.Sprintf(format, args...))
}

func (l *Logger) Debug(message string, data ...map[string]interface{}) {
	l.Append(LevelDebug, message, data...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Appendf(LevelDebug, format, args...)
}

func (l *Logger) Info(message string, data ...map[string]interface{}) {
	l.Append(LevelInfo, message, data...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Appendf(LevelInfo, format, args...)
}

func (l *Logger) Warn(message string, data ...map[string]interface{}) {
	l.Append(LevelWarn, message, data...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Appendf(LevelWarn, format, args...)
}

func (l *Logger) Error(message string, data ...map[string]interface{}) {
	l.Append(LevelError, message, data...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Appendf(LevelError, format, args...)
}

func (l *Logger) Fatal(message string, data ...map[string]interface{}) {
	l.Append(LevelFatal, message, data...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Appendf(LevelFatal, format, args...)
}
