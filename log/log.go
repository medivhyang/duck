package log

import "os"

type Fields = map[string]interface{}

var DefaultInstance = New(os.Stdout, LevelDebug)

func Write(e Entry) {
	DefaultInstance.Write(e)
}

func Debug(message string, data ...map[string]interface{}) {
	DefaultInstance.Debug(message, data...)
}

func Info(message string, data ...map[string]interface{}) {
	DefaultInstance.Info(message, data...)
}

func Warn(message string, data ...map[string]interface{}) {
	DefaultInstance.Warn(message, data...)
}

func Error(message string, data ...map[string]interface{}) {
	DefaultInstance.Error(message, data...)
}

func Fatal(message string, data ...map[string]interface{}) {
	DefaultInstance.Fatal(message, data...)
}
