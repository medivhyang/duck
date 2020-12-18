package log

import (
	"bytes"
	"encoding/json"
)

type Formatter interface {
	Format(e Event) ([]byte, error)
}

type TextFormatter struct{}

func (*TextFormatter) Format(e Event) ([]byte, error) {
	buffer := bytes.Buffer{}
	buffer.WriteString(e.Message)
	return buffer.Bytes(), nil
}

type JSONFormatter struct{}

func (*JSONFormatter) Format(e Event) ([]byte, error) {
	return json.Marshal(e)
}
