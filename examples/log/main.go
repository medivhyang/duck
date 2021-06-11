package main

import (
	"github.com/medivhyang/duck/log"
)

func main() {
	log.Info("hello world", log.Fields{"name": "Medivh"})
	log.Infof("hello %s", "Medivh")

	l := log.New("demo", log.LevelDebug, log.NewConsoleAppender(log.NewJSONFormatter().Pretty("", "\t")))
	l.Info("hello world", log.Fields{"name": "Medivh"})
	l.Infof("hello %s", "Medivh")
}
