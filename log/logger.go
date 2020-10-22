package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"
)

type Logger interface {
	Debug(message string, data ...interface{})
	Info(message string, data ...interface{})
	Warn(message string, data ...interface{})
	Error(message string, data ...interface{})
	Fatal(message string, data ...interface{})
}

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

type Format string

const (
	FormatJSON Format = "json"
	FormatText Format = "text"
)

type Middleware = func(e Entry) Entry

type logger struct {
	level       Level
	output      io.Writer
	module      string
	format      Format
	middlewares []Middleware

	beauty     bool
	indent     string
	prefix     string
	timeFormat string
}

type Entry struct {
	Level   Level
	Message string
	Data    map[string]interface{}
	Time    time.Time
}

type View struct {
	Time    string                 `json:"time"`
	Level   string                 `json:"level"`
	Module  string                 `json:"module,omitempty"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

func (e *Entry) View(module string) View {
	return View{
		Time:    e.Time.Format(time.RFC3339),
		Level:   LevelText(e.Level),
		Module:  module,
		Message: e.Message,
		Data:    e.Data,
	}
}

func New(output io.Writer, level Level, middlewares ...Middleware) *logger {
	return &logger{
		level:       level,
		output:      output,
		middlewares: middlewares,
	}
}

func (l *logger) Level(level Level) *logger {
	l.level = level
	return l
}

func (l *logger) Module(name string) *logger {
	l.module = name
	return l
}

func (l *logger) Output(writer io.Writer) *logger {
	l.output = writer
	return l
}

func (l *logger) Beauty(ok bool) *logger {
	l.beauty = ok
	return l
}

func (l *logger) Indent(prefix string, indent string) *logger {
	l.prefix = prefix
	l.indent = indent
	return l
}

func (l *logger) Format(format Format) *logger {
	l.format = format
	return l
}

func (l *logger) Use(middlewares ...Middleware) *logger {
	l.middlewares = append(l.middlewares, middlewares...)
	return l
}

func (l *logger) Write(e Entry) {
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
	switch l.format {
	case FormatJSON:
		l.writeJSON(e)
	case FormatText:
		fallthrough
	default:
		l.writeText(e)
	}
}

func (l *logger) writeText(e Entry) {
	if l.output == nil {
		return
	}
	v := e.View(l.module)
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("%s: %s", v.Time, v.Level))
	if l.module != "" {
		b.WriteString(fmt.Sprintf(": %s", l.module))
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

func (l *logger) writeJSON(e Entry) {
	if l.output == nil {
		return
	}
	v := e.View(l.module)
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

func (l *logger) Debug(message string, data ...map[string]interface{}) {
	var finalData map[string]interface{}
	if len(data) > 0 {
		finalData = data[0]
	}
	l.Write(Entry{Level: LevelDebug, Message: message, Data: finalData})
}

func (l *logger) Info(message string, data ...map[string]interface{}) {
	var finalData map[string]interface{}
	if len(data) > 0 {
		finalData = data[0]
	}
	l.Write(Entry{Level: LevelInfo, Message: message, Data: finalData})
}

func (l *logger) Warn(message string, data ...map[string]interface{}) {
	var finalData map[string]interface{}
	if len(data) > 0 {
		finalData = data[0]
	}
	l.Write(Entry{Level: LevelWarn, Message: message, Data: finalData})
}

func (l *logger) Error(message string, data ...map[string]interface{}) {
	var finalData map[string]interface{}
	if len(data) > 0 {
		finalData = data[0]
	}
	l.Write(Entry{Level: LevelError, Message: message, Data: finalData})
}

func (l *logger) Fatal(message string, data ...map[string]interface{}) {
	var finalData map[string]interface{}
	if len(data) > 0 {
		finalData = data[0]
	}
	l.Write(Entry{Level: LevelFatal, Message: message, Data: finalData})
}
