package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hash(word string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(word), bcrypt.MinCost)

	return string(bytes)
}

func main() {
	word := "Leve 1, Pague 2"
	hashed := hash(word)
	fmt.Println(hashed)
	fmt.Println("Hash is", len(hashed), "long")
}
