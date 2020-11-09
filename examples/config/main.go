package main

import (
	"fmt"
	"github.com/medivhyang/duck/config"
)

func main() {
	instance := struct {
		Addr  string `json:"addr"`
		Debug bool   `json:"debug"`
	}{
		Addr:  ":8080",
		Debug: true,
	}
	if err := config.LoadOrStoreFile(config.JSON, "config.json", &instance); err != nil {
		panic(err)
	}
	fmt.Println(instance)
}
