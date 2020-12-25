package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
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
	buf := bytes.Buffer{}

	buf.WriteString(fmt.Sprintf("%s", e.Time.Format(f.timeLayout)))
	if e.Module != "" {
		buf.WriteString(fmt.Sprintf(": [%s]", e.Module))
	}
	if LevelText(e.Level) != "" {
		buf.WriteString(fmt.Sprintf(": [%s]", LevelText(e.Level)))
	}
	if len(e.File) > 0 {
		buf.WriteString(": ")
		buf.WriteString(e.File)
		buf.WriteString(":")
		buf.WriteString(strconv.Itoa(e.Line))
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
