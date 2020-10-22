package main

import (
	"fmt"
	"github.com/medivhyang/duck/options"
)

func main() {
	a := options.String{
		Valid: true,
		Value: "hello",
	}
	fmt.Println(a.Unwrap())

	b := options.String{}
	fmt.Println(b.Unwrap("default"))
}
