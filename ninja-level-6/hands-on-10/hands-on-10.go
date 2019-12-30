package main

import "fmt"

func foo() func() int {
	x := 0
	return func() int {
		return x
	}
}

func main() {
	fmt.Println(foo())
}
