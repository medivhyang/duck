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
	if err := requests.SaveFile("https://www.baidu.com", "demo.txt"); err != nil {
		panic(err)
	}
}
