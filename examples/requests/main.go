package main

import (
	"fmt"
	"github.com/medivhyang/duck/requests"
)

func main() {
	s, err := requests.GetText("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
