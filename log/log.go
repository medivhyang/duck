package log

import "os"

type Fields = map[string]interface{}

var Default = New(WithOutput(os.Stdout), WithLevel(LevelDebug))

func Write(e Entry) {
	Default.Write(e)
}

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
