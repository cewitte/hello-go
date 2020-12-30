package main

import (
	"fmt"

	money "github.com/Rhymond/go-money"
)

func main() {
	real := money.New(258, "BRL")
	thirt := money.New(13, "BRL")

	sum, _ := real.Add(thirt)

	fmt.Println(sum.Display())

}
