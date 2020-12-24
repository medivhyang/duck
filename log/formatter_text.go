package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

func NewTextFormatter() *TextFormatter {
	return &TextFormatter{timeLayout: time.RFC3339}
}

type TextFormatter struct {
	timeLayout string
}

func (f *TextFormatter) SetTimeLayout(layout string) *TextFormatter {
	f.timeLayout = layout
	return f
}

func (f *TextFormatter) Format(e Event) ([]byte, error) {
	buffer := bytes.Buffer{}

	buffer.WriteString(fmt.Sprintf("%s", e.Time.Format(f.timeLayout)))
	if e.Module != "" {
		buffer.WriteString(fmt.Sprintf(": [%s]", e.Module))
	}
	if LevelText(e.Level) != "" {
		buffer.WriteString(fmt.Sprintf(": [%s]", LevelText(e.Level)))
	}
	buffer.WriteString(fmt.Sprintf(": %s", e.Message))
	if e.Data != nil {
		bs, err := json.Marshal(e.Data)
		if err != nil {
			buffer.WriteString(fmt.Sprintf(": <json marshal fail: %s>", err.Error()))
		} else {
			buffer.WriteString(fmt.Sprintf(": %s", string(bs)))
		}
	}
	buffer.WriteString("\r\n")
	return buffer.Bytes(), nil
}
