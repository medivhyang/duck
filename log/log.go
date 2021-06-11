package log

import "time"

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

var levelText = map[Level]string{
	LevelDebug: "debug",
	LevelInfo:  "info",
	LevelWarn:  "warn",
	LevelError: "error",
	LevelFatal: "fatal",
}

func LevelText(l Level) string {
	return levelText[l]
}

type (
	Event struct {
		Module  string
		Level   Level
		Message string
		Data    map[string]interface{}
		Time    time.Time
	}
	Appender interface {
		Append(e Event) error
	}
	Formatter interface {
		Format(e Event) ([]byte, error)
	}
	Fields = map[string]interface{}
)

var Default = New("", LevelDebug, NewConsoleAppender(NewTextFormatter()))

func Debug(message string, data ...map[string]interface{}) {
	Default.Debug(message, data...)
}

func Debugf(format string, args ...interface{}) {
	Default.Debugf(format, args...)
}

func Info(message string, data ...map[string]interface{}) {
	Default.Info(message, data...)
}

func Infof(format string, args ...interface{}) {
	Default.Infof(format, args...)
}

func Warn(message string, data ...map[string]interface{}) {
	Default.Warn(message, data...)
}

func Warnf(format string, args ...interface{}) {
	Default.Warnf(format, args...)
}

func Error(message string, data ...map[string]interface{}) {
	Default.Error(message, data...)
}

func Errorf(format string, args ...interface{}) {
	Default.Errorf(format, args...)
}

func Fatal(message string, data ...map[string]interface{}) {
	Default.Fatal(message, data...)
}

func Fatalf(format string, args ...interface{}) {
	Default.Fatalf(format, args...)
}
