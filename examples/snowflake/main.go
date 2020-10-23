package main

import (
	"fmt"
	"github.com/medivhyang/duck/snowflake"
)

func main() {
	fmt.Println(snowflake.Generate().String())
	fmt.Println(snowflake.Generate().String())
	fmt.Println(snowflake.Generate().Base64())
	fmt.Println(snowflake.Generate().Base64())
	fmt.Println(snowflake.Generate().MD5())
	fmt.Println(snowflake.Generate().MD5())
	fmt.Println(snowflake.Generate().Base36())
	fmt.Println(snowflake.Generate().Base36())
	fmt.Println(snowflake.Generate().Base2())
	fmt.Println(snowflake.Generate().Base2())
}
