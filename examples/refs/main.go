package main

import (
	"fmt"
	"github.com/medivhyang/duck/refs"
)

func main() {
	a := refs.String("hello")
	fmt.Println(refs.UnwrapString(a, "default"))

	b := (*string)(nil)
	fmt.Println(refs.UnwrapString(b, "default"))
}
