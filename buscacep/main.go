package main

import (
	"fmt"

	"github.com/igorhalfeld/lagoinha"
)

func main() {
	address, _ := lagoinha.GetAddress("06543315")
	fmt.Printf("Complete Address %v:", address)
}

// Alternativa JS
// https://viacep.com.br/exemplo/javascript/
