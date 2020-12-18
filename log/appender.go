package log

import "io"

type Appender interface {
	Append(e Event) error
	SetFormatter(Formatter)
}

type WriterAppender struct {
	Formatter Formatter
	Writer    io.Writer
}

func (a *WriterAppender) Append(e Event) error {
	b, err := a.Formatter.Format(e)
	if err != nil {
		return err
	}
	_, err = a.Writer.Write(b)
	return err
}

func (a *WriterAppender) SetFormatter(f Formatter) {
	a.Formatter = f
}
