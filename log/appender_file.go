package log

import "os"

type FileAppender struct {
	level     Level
	formatter Formatter
	file      *os.File
}

func NewFileAppender(filename string, formatter Formatter) (*FileAppender, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0744)
	if err != nil {
		return nil, err
	}
	return &FileAppender{file: file, formatter: formatter}, nil
}

func (a *FileAppender) SetLevel(level Level) *FileAppender {
	a.level = level
	return a
}

func (a *FileAppender) SetFormatter(f Formatter) *FileAppender {
	a.formatter = f
	return a
}

func (a *FileAppender) Append(e Event) error {
	if e.Level < a.level {
		return nil
	}
	if a.file == nil {
		return nil
	}
	if a.formatter == nil {
		return nil
	}
	bs, err := a.formatter.Format(e)
	if err != nil {
		return err
	}
	if _, err := a.file.Write(bs); err != nil {
		return err
	}
	return nil
}
