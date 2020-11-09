package main

import (
	"fmt"
	"github.com/medivhyang/duck/slices"
)

func main() {
	source := []string{"1", "2", "3", "3", "4", "5"}

	fmt.Println(slices.UniqueStrings(source))
	fmt.Println(slices.Strings(source).Unique())

	fmt.Println(slices.RemoveStrings(source, []string{"1", "2"}))
	fmt.Println(slices.Strings(source).Remove([]string{"1", "2"}))

	fmt.Println(slices.ToInts(source))
	fmt.Println(slices.Strings(source).ToInts())

	fmt.Println(source)
}
