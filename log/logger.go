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

func (l *Logger) Module(module string) *Logger {
	l.module = module
	return l
}

func (l *Logger) Level(level Level) *Logger {
	l.level = level
	return l
}

func (l *Logger) Append(appenders ...Appender) *Logger {
	l.appenders = append(l.appenders, appenders...)
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

func (l *Logger) FileSkip(skip int) *Logger {
	l.fileSkip = skip
	return l
}

func (l *Logger) TimeLocation(location *time.Location) *Logger {
	l.timeLocation = location
	return l
}

func (l *Logger) appendEvent(input EventInput) {
	if input.Level < l.level {
		return
	}
	var e = Event{
		Module:  input.Module,
		Level:   input.Level,
		Message: input.Message,
		Data:    input.Data,
		Time:    time.Now(),
	}
	if l.timeLocation != nil {
		e.Time = e.Time.In(l.timeLocation)
	}
	if l.enableFile && input.Level >= l.fileLevel {
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
	e := EventInput{
		Level:   LevelDebug,
		Message: message,
		Data:    nil,
	}
	if len(data) > 0 {
		e.Data = data[0]
	}
	l.appendEvent(e)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	e := EventInput{
		Level:   LevelDebug,
		Message: fmt.Sprintf(format, args...),
	}
	l.appendEvent(e)
}

func (l *Logger) Info(message string, data ...map[string]interface{}) {
	e := EventInput{
		Level:   LevelInfo,
		Message: message,
		Data:    nil,
	}
	if len(data) > 0 {
		e.Data = data[0]
	}
	l.appendEvent(e)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	e := EventInput{
		Level:   LevelInfo,
		Message: fmt.Sprintf(format, args...),
	}
	l.appendEvent(e)
}

func (l *Logger) Warn(message string, data ...map[string]interface{}) {
	e := EventInput{
		Level:   LevelWarn,
		Message: message,
		Data:    nil,
	}
	if len(data) > 0 {
		e.Data = data[0]
	}
	l.appendEvent(e)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	e := EventInput{
		Level:   LevelWarn,
		Message: fmt.Sprintf(format, args...),
	}
	l.appendEvent(e)
}

func (l *Logger) Error(message string, data ...map[string]interface{}) {
	e := EventInput{
		Level:   LevelError,
		Message: message,
		Data:    nil,
	}
	if len(data) > 0 {
		e.Data = data[0]
	}
	l.appendEvent(e)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	e := EventInput{
		Level:   LevelError,
		Message: fmt.Sprintf(format, args...),
	}
	l.appendEvent(e)
}

func (l *Logger) Fatal(message string, data ...map[string]interface{}) {
	e := EventInput{
		Level:   LevelFatal,
		Message: message,
		Data:    nil,
	}
	if len(data) > 0 {
		e.Data = data[0]
	}
	l.appendEvent(e)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	e := EventInput{
		Level:   LevelFatal,
		Message: fmt.Sprintf(format, args...),
	}
	l.appendEvent(e)
}
