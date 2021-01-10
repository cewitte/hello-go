package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	p := message.NewPrinter(language.BrazilianPortuguese)
	p.Printf("Há %v flores no meu jardim. \n", 1500.19)

	p = message.NewPrinter(language.AmericanEnglish)
	p.Printf("There are %v flowers in my garden. \n", 1500.19)

	message.SetString(language.BrazilianPortuguese, "Hello %s", "Olá %s, tudo bem?")
	message.SetString(language.AmericanEnglish, "Hello %s", "Hello %s, how are you?")

	p = message.NewPrinter(language.BrazilianPortuguese)
	p.Printf("Hello %s", "Carlos")
	fmt.Println()

	p = message.NewPrinter(language.AmericanEnglish)
	p.Printf("Hello %s", "Carlos")
	fmt.Println()

}
