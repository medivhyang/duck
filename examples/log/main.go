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
	log.Debug("hello world", map[string]interface{}{"name": "Medivh"})
	log.Debugf("hello %s", "Medivh")
}

func demoText() {
	l := log.New(log.WithEncoding(log.EncodingText), log.WithModule("test.text"))
	l.Debug("hello world", map[string]interface{}{"name": "Medivh"})
	l.Debugf("hello %s", "Medivh")
}

func demoJSON() {
	l := log.New(log.WithEncoding(log.EncodingJSON), log.WithModule("test.json"))
	l.Debug("hello world", map[string]interface{}{"name": "Medivh"})
	l.Debugf("hello %s", "Medivh")
}
