package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

func NewTextFormatter() *TextFormatter {
	return &TextFormatter{layout: time.RFC3339}
}

type TextFormatter struct {
	layout string
}

func (f *TextFormatter) Layout(l string) *TextFormatter {
	f.layout = l
	return f
}

func (f *TextFormatter) Format(e Event) ([]byte, error) {
	buf := bytes.Buffer{}

	buf.WriteString(fmt.Sprintf("%s", e.Time.Format(f.layout)))
	if e.Module != "" {
		buf.WriteString(fmt.Sprintf(": [%s]", e.Module))
	}
	if LevelText(e.Level) != "" {
		buf.WriteString(fmt.Sprintf(": [%s]", LevelText(e.Level)))
	}
	buf.WriteString(fmt.Sprintf(": %s", e.Message))
	if e.Data != nil {
		bs, err := json.Marshal(e.Data)
		if err != nil {
			buf.WriteString(fmt.Sprintf(": <json marshal fail: %s>", err.Error()))
		} else {
			buf.WriteString(fmt.Sprintf(": %s", string(bs)))
		}
	}
	buf.WriteString("\n")

	return buf.Bytes(), nil
}
