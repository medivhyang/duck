package log

type Formatter interface {
	Format(e Event) ([]byte, error)
}
