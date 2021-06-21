package main

import (
	"fmt"
	"github.com/medivhyang/duck/naming"
	"github.com/medivhyang/duck/reflectutil"
)

func main() {
	type user struct {
		Name string
		Age  int
	}
	m := map[string]interface{}{
		"name": "Medivh",
		"age":  3,
	}
	var u user
	reflectutil.ParseMapToStruct(m, &u, naming.ToPascal)
	fmt.Println(u)
}
