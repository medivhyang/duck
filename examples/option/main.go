package main

import (
	"fmt"
	"github.com/medivhyang/duck"
)

func main() {
	a := duck.OptionString{
		Valid: true,
		Value: "hello",
	}
	fmt.Println(a.ValueOrDefault())

	b := duck.OptionString{}
	fmt.Println(b.ValueOrDefault("default"))
}
