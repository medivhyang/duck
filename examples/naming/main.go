package main

import (
	"fmt"
	"github.com/medivhyang/duck/naming"
)

func main() {
	s := "helloWorld"
	fmt.Println(naming.ToPascal(s))
	fmt.Println(naming.ToCamel(s))
	fmt.Println(naming.ToSnake(s))
	fmt.Println(naming.ToKebab(s))
}
