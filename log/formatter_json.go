package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{layout: time.RFC3339}
}

type JSONFormatter struct {
	pretty bool
	prefix string
	indent string
	layout string
}

func (f *JSONFormatter) Pretty(prefix string, indent string) *JSONFormatter {
	f.pretty = true
	f.prefix = prefix
	f.indent = indent
	return f
}

func (f *JSONFormatter) Layout(l string) *JSONFormatter {
	f.layout = l
	return f
}

func (f *JSONFormatter) Format(e Event) ([]byte, error) {
	view := struct {
		Module  string                 `json:"module,omitempty"`
		Level   string                 `json:"level,omitempty"`
		Message string                 `json:"message,omitempty"`
		Time    string                 `json:"time,omitempty"`
		Data    map[string]interface{} `json:"data,omitempty"`
	}{
		Module:  e.Module,
		Level:   LevelText(e.Level),
		Message: e.Message,
		Time:    e.Time.Format(f.layout),
		Data:    e.Data,
	}

	buf := bytes.Buffer{}
	if f.pretty {
		bs, err := json.MarshalIndent(view, f.prefix, f.indent)
		if err != nil {
			buf.WriteString(fmt.Sprintf("{\"error\":\"<json marshal indent fail: %s>\"}", err.Error()))
		} else {
			buf.Write(bs)
		}
	} else {
		bs, err := json.Marshal(view)
		if err != nil {
			buf.WriteString(fmt.Sprintf("{\"error\":\"<json marshal fail: %s>\"}", err.Error()))
		} else {
			buf.Write(bs)
		}
	}

	buf.WriteString("\r\n")

	return buf.Bytes(), nil
}
