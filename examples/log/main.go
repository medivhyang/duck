package main

import (
	"github.com/medivhyang/duck/log"
	"time"
)

func main() {
	demoSimple()
	demoModule()
	demoData()
	demoTimeLocation()
	demoJSON()
	demoFile()
	demoFileAppender()
}

func demoSimple() {
	log.Info("hello world")
	log.Infof("helo %s", "Medivh")
}

func demoModule() {
	log.Default.New("test").Info("hello world")
}

func demoData() {
	log.Error("terrible", log.Fields{"name": "ying", "age": 25})
}

func demoTimeLocation() {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	l := log.Default.New("").TimeLocation(loc)
	l.Info("hello world")
}

func demoJSON() {
	l := log.Default.New("").Append(log.NewConsoleAppender(log.NewJSONFormatter()))
	l.Error("terrible", log.Fields{"name": "ying", "age": 25})
}

func demoFile() {
	l := log.Default.New("").EnableFile(log.LevelError)
	l.Info("hello world")
	l.Error("terrible")
}

func demoFileAppender() {
	fileAppender, err := log.NewFileAppender("log.txt", log.NewJSONFormatter())
	if err != nil {
		panic(err)
	}
	l := log.New("", log.LevelDebug, fileAppender)
	l.Info("hello world")
	l.Error("terrible", log.Fields{"name": "ying", "age": 25})
}
