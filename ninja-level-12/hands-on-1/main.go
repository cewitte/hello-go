package main

import (
	"fmt"
	"github.com/cewitte/hello-go/ninja-level-12/hands-on-1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	// dogYears := Years(7)
	// fmt.Printf("Human's age in dog years is %v", dogYears)

	fido := canine{
		name: "Fido",
		age:  dog.Years(7),
	}

	fmt.Println(fido)

}
