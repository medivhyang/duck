package main

import (
	"github.com/medivhyang/duck/log"
	"os"
)

func main() {
	demoText()
	demoJSON()
}

func demoText() {
	l := log.New(os.Stdout, log.LevelDebug).Module("test.text").Format(log.FormatText)
	l.Debug("hello world", map[string]interface{}{"name": "Medivh"})
}

func demoJSON() {
	l := log.New(os.Stdout, log.LevelDebug).Module("test.json").Format(log.FormatJSON)
	l.Debug("hello world", map[string]interface{}{"name": "Medivh"})
}
