package main

import (
	"fmt"
	"github.com/medivhyang/duck/requests"
	"log"
)

func main() {
	s, err := requests.GetText("https://www.baidu.com")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(s)
}
