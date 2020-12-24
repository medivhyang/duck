package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{timeLayout: time.RFC3339}
}

type JSONFormatter struct {
	pretty     bool
	prefix     string
	indent     string
	timeLayout string
}

func (f *JSONFormatter) SetPretty(pretty bool) *JSONFormatter {
	f.pretty = pretty
	return f
}

func (f *JSONFormatter) SetIndent(prefix string, indent string) *JSONFormatter {
	f.prefix = prefix
	f.indent = indent
	return f
}

func (f *JSONFormatter) SetTimeLayout(layout string) *JSONFormatter {
	f.timeLayout = layout
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
		Time:    e.Time.Format(f.timeLayout),
		Data:    e.Data,
	}

	buffer := bytes.Buffer{}
	if f.pretty {
		bs, err := json.MarshalIndent(view, f.prefix, f.indent)
		if err != nil {
			buffer.WriteString(fmt.Sprintf("{\"error\":\"<json marshal indent fail: %s>\"}", err.Error()))
		} else {
			buffer.Write(bs)
		}
	} else {
		bs, err := json.Marshal(view)
		if err != nil {
			buffer.WriteString(fmt.Sprintf("{\"error\":\"<json marshal fail: %s>\"}", err.Error()))
		} else {
			buffer.Write(bs)
		}
	}

	buffer.WriteString("\r\n")

	return buffer.Bytes(), nil
}
