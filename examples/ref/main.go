package main

import (
	"fmt"
	"github.com/medivhyang/duck"
)

func main() {
	a := duck.RefString("hello")
	fmt.Println(duck.DerefString(a, "default"))

	b := (*string)(nil)
	fmt.Println(duck.DerefString(b, "default"))
}
