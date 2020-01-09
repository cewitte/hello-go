package main

import (
	"fmt"

	"github.com/cewitte/hello-go/ninja-level-13/hands-on-2/quote"
	"github.com/cewitte/hello-go/ninja-level-13/hands-on-2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
