package log

import (
	"fmt"
	"runtime"
	"time"
)

func New(module string, level Level, appenders ...Appender) *Logger {
	return &Logger{
		module:       module,
		level:        level,
		fileSkip:     2,
		timeLocation: time.Local,
		appenders:    appenders,
	}
}

type Logger struct {
	module       string
	level        Level
	enableFile   bool
	fileLevel    Level
	fileSkip     int
	timeLocation *time.Location
	appenders    []Appender
}

func (l *Logger) New(module string) *Logger {
	return &Logger{
		module:       module,
		level:        l.level,
		enableFile:   l.enableFile,
		fileLevel:    l.fileLevel,
		fileSkip:     l.fileSkip,
		timeLocation: l.timeLocation,
		appenders:    append([]Appender{}, l.appenders...),
	}
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

func (l *Logger) EnableFile(level Level) *Logger {
	l.enableFile = true
	l.fileLevel = level
	return l
}

func (l *Logger) DisableFile() *Logger {
	l.enableFile = false
	return l
}

func (l *Logger) SetFileSkip(skip int) *Logger {
	l.fileSkip = skip
	return l
}

func (l *Logger) SetTimeLocation(location *time.Location) *Logger {
	l.timeLocation = location
	return l
}

func (l *Logger) append(level Level, message string, data ...map[string]interface{}) {
	if level < l.level {
		return
	}
	e := Event{
		Module:  l.module,
		Level:   level,
		Message: message,
		Time:    time.Now(),
	}
	if l.timeLocation != nil {
		e.Time = e.Time.In(l.timeLocation)
	}
	if len(data) > 0 {
		e.Data = data[0]
	}
	if l.enableFile && level >= l.fileLevel {
		_, file, line, ok := runtime.Caller(l.fileSkip)
		if !ok {
			file = "???"
			line = 0
		}
		e.File = file
		e.Line = line
	}
	for _, appender := range l.appenders {
		_ = appender.Append(e)
	}
	return
}

func (l *Logger) Debug(message string, data ...map[string]interface{}) {
	l.append(LevelDebug, message, data...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.append(LevelDebug, fmt.Sprintf(format, args...))
}

func (l *Logger) Info(message string, data ...map[string]interface{}) {
	l.append(LevelInfo, message, data...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.append(LevelInfo, fmt.Sprintf(format, args...))
}

func (l *Logger) Warn(message string, data ...map[string]interface{}) {
	l.append(LevelWarn, message, data...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.append(LevelWarn, fmt.Sprintf(format, args...))
}

func (l *Logger) Error(message string, data ...map[string]interface{}) {
	l.append(LevelError, message, data...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.append(LevelError, fmt.Sprintf(format, args...))
}

func (l *Logger) Fatal(message string, data ...map[string]interface{}) {
	l.append(LevelFatal, message, data...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.append(LevelFatal, fmt.Sprintf(format, args...))
}
