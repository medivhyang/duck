package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

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

type Encoding string

const (
	EncodingJSON Encoding = "json"
	EncodingText Encoding = "text"
)

type Middleware = func(e Entry) Entry

type Logger struct {
	output      io.Writer
	level       Level
	module      string
	encoding    Encoding
	middlewares []Middleware

	beauty bool
	indent string
	prefix string

	timeFormat string
}

type Entry struct {
	Level   Level
	Message string
	Data    map[string]interface{}
	Time    time.Time
}

type view struct {
	Time    string                 `json:"time"`
	Level   string                 `json:"level"`
	Module  string                 `json:"module,omitempty"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

func (e *Entry) view(module string, timeFormat string) view {
	if timeFormat == "" {
		timeFormat = time.RFC3339
	}
	return view{
		Time:    e.Time.Format(timeFormat),
		Level:   LevelText(e.Level),
		Module:  module,
		Message: e.Message,
		Data:    e.Data,
	}
}

func New(options ...Option) *Logger {
	return (&Logger{level: LevelDebug, output: os.Stdout}).Apply(options...)
}

func (l *Logger) Apply(options ...Option) *Logger {
	for _, option := range options {
		option(l)
	}
	return l
}

func (l *Logger) Write(e Entry) {
	if l.output == nil {
		return
	}
	for _, mid := range l.middlewares {
		e = mid(e)
	}
	if e.Level < l.level {
		return
	}
	if e.Time == (time.Time{}) {
		e.Time = time.Now()
	}
	switch l.encoding {
	case EncodingJSON:
		l.writeJSON(e)
	case EncodingText:
		fallthrough
	default:
		l.writeText(e)
	}
}

func (l *Logger) writeText(e Entry) {
	if l.output == nil {
		return
	}
	v := e.view(l.module, l.timeFormat)
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("%s: [%s]", v.Time, v.Level))
	if l.module != "" {
		b.WriteString(fmt.Sprintf(": [%s]", l.module))
	}
	b.WriteString(fmt.Sprintf(": %s", v.Message))
	if v.Data != nil {
		bytes, err := json.Marshal(v.Data)
		if err != nil {
			b.WriteString(fmt.Sprintf(": <json marshal fail: %s>", err.Error()))
		} else {
			b.WriteString(fmt.Sprintf(": %s", string(bytes)))
		}
	}
	b.WriteString("\r\n")
	_, _ = io.WriteString(l.output, b.String())
	return
}

func (l *Logger) writeJSON(e Entry) {
	if l.output == nil {
		return
	}
	v := e.view(l.module, l.timeFormat)
	b := bytes.Buffer{}
	if l.beauty {
		bs, err := json.MarshalIndent(v, l.prefix, l.indent)
		if err != nil {
			b.WriteString(fmt.Sprintf("{\"error\":\"<json marshal indent fail: %s>\"}", err.Error()))
		} else {
			b.Write(bs)
		}
	} else {
		bs, err := json.Marshal(v)
		if err != nil {
			b.WriteString(fmt.Sprintf("{\"error\":\"<json marshal fail: %s>\"}", err.Error()))
		} else {
			b.Write(bs)
		}
	}
	b.WriteString("\r\n")
	_, _ = l.output.Write(b.Bytes())
}

func (l *Logger) Debug(message string, data ...map[string]interface{}) {
	var finalData map[string]interface{}
	if len(data) > 0 {
		finalData = data[0]
	}
	l.Write(Entry{Level: LevelDebug, Message: message, Data: finalData})
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Write(Entry{Level: LevelDebug, Message: fmt.Sprintf(format, args...)})
}

func (l *Logger) Info(message string, data ...map[string]interface{}) {
	var finalData map[string]interface{}
	if len(data) > 0 {
		finalData = data[0]
	}
	l.Write(Entry{Level: LevelInfo, Message: message, Data: finalData})
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Write(Entry{Level: LevelInfo, Message: fmt.Sprintf(format, args...)})
}

func (l *Logger) Warn(message string, data ...map[string]interface{}) {
	var finalData map[string]interface{}
	if len(data) > 0 {
		finalData = data[0]
	}
	l.Write(Entry{Level: LevelWarn, Message: message, Data: finalData})
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Write(Entry{Level: LevelWarn, Message: fmt.Sprintf(format, args...)})
}

func (l *Logger) Error(message string, data ...map[string]interface{}) {
	var finalData map[string]interface{}
	if len(data) > 0 {
		finalData = data[0]
	}
	l.Write(Entry{Level: LevelError, Message: message, Data: finalData})
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Write(Entry{Level: LevelError, Message: fmt.Sprintf(format, args...)})
}

func (l *Logger) Fatal(message string, data ...map[string]interface{}) {
	var finalData map[string]interface{}
	if len(data) > 0 {
		finalData = data[0]
	}
	l.Write(Entry{Level: LevelFatal, Message: message, Data: finalData})
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Write(Entry{Level: LevelFatal, Message: fmt.Sprintf(format, args...)})
}

type Option func(l *Logger)

func WithLevel(level Level) Option {
	return func(l *Logger) {
		l.level = level
	}
}

func WithModule(name string) Option {
	return func(l *Logger) {
		l.module = name
	}
}

func WithOutput(writer io.Writer) Option {
	return func(l *Logger) {
		l.output = writer
	}
}

func WithBeauty(ok bool) Option {
	return func(l *Logger) {
		l.beauty = ok
	}
}

func WithIndent(prefix string, indent string) Option {
	return func(l *Logger) {
		l.prefix = prefix
		l.indent = indent
	}
}

func WithEncoding(encoding Encoding) Option {
	return func(l *Logger) {
		l.encoding = encoding
	}
}

func WithTimeFormat(format string) Option {
	return func(l *Logger) {
		l.timeFormat = format
	}
}

func WithMiddlewares(middlewares ...Middleware) Option {
	return func(l *Logger) {
		l.middlewares = append(l.middlewares, middlewares...)
	}
}
