package log

import "time"

type Event struct {
	Module  string
	Level   Level
	Message string
	Time    time.Time
	Data    map[string]interface{}
	File    string
	Line    int
}
