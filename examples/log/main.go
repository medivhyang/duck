package main

import (
	"github.com/medivhyang/duck/log"
)

func main() {
	demoDefault()
	demoText()
	demoJSON()
}

func demoDefault() {
	log.Debug("hello world", log.Fields{"name": "medivh"})
	log.Debugf("hello %s", "Medivh")
}

func demoText() {
	l := log.New("demo.text", log.LevelDebug, log.NewConsoleAppender(log.NewTextFormatter()))
	l.Debug("hello world", log.Fields{"name": "Medivh"})
	l.Debugf("hello %s", "Medivh")
}

func demoJSON() {
	l := log.New("demo.json", log.LevelDebug, log.NewConsoleAppender(log.NewJSONFormatter()))
	l.Debug("hello world", log.Fields{"name": "medivh", "age": 18})

	fileAppender, err := log.NewFileAppender("log.txt", log.NewJSONFormatter())
	if err != nil {
		panic(err)
	}
	fileAppender = fileAppender.SetLevel(log.LevelWarn)
	l.SetModule("demo.json.file").SetAppenders(
		fileAppender,
	)
	l.Debug("hello world")
	l.Warn("hello world")
}
