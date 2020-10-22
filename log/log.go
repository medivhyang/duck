package log

import "os"

var DefaultInstance = New(os.Stdout, LevelDebug)

func Write(e Entry) {
	DefaultInstance.Write(e)
}

func Debug(message string, data ...interface{}) {
	DefaultInstance.Debug(message, data...)
}

func Info(message string, data ...interface{}) {
	DefaultInstance.Info(message, data...)
}

func Warn(message string, data ...interface{}) {
	DefaultInstance.Warn(message, data...)
}

func Error(message string, data ...interface{}) {
	DefaultInstance.Error(message, data...)
}

func Fatal(message string, data ...interface{}) {
	DefaultInstance.Fatal(message, data...)
}
