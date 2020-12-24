package log

type Appender interface {
	Append(Event) error
}
