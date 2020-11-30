package main

import (
	"fmt"
	"github.com/medivhyang/duck/rsa"
)

func main() {
	if err := rsa.GenerateKeyPairFiles(2048, "public.pem", "private.pem"); err != nil {
		panic(err)
	}

	plainText := "hello world"
	cipherText, err := rsa.EncryptWithFile([]byte(plainText), "public.pem")
	if err != nil {
		panic(err)
	}

	result, err := rsa.DecryptWithFile(cipherText, "private.pem")
	if err != nil {
		panic(err)
	}

	if string(result) != plainText {
		fmt.Printf("want: %s, got: %s\n", plainText, result)
		return
	}

	fmt.Println("pass!")
}
