package log

import (
	"os"
)

type ConsoleAppender struct {
	level     Level
	formatter Formatter
}

func NewConsoleAppender(formatter Formatter) *ConsoleAppender {
	return &ConsoleAppender{formatter: formatter}
}

func (a *ConsoleAppender) SetLevel(level Level) Appender {
	a.level = level
	return a
}

func (a *ConsoleAppender) SetFormatter(formatter Formatter) Appender {
	a.formatter = formatter
	return a
}

func (a *ConsoleAppender) Append(e Event) error {
	if e.Level < a.level {
		return nil
	}
	if a.formatter == nil {
		return nil
	}
	bs, err := a.formatter.Format(e)
	if err != nil {
		return err
	}
	if _, err := os.Stdout.Write(bs); err != nil {
		return err
	}
	return nil
}
