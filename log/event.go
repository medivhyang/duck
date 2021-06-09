package log

import "time"

type Event struct {
	Module  string
	Level   Level
	Message string
	Data    map[string]interface{}
	Time    time.Time
	File    string
	Line    int
}

type EventInput struct {
	Module  string
	Level   Level
	Message string
	Data    map[string]interface{}
}
