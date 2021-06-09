package log

type Appender interface {
	Append(e Event) error
}
