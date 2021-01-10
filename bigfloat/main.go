package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := new(big.Float).SetFloat64(1.555)
	b := new(big.Float).SetFloat64(1.555)
	c := new(big.Float).Add(a, b)

	fmt.Printf("%.2f\n", c)
}
