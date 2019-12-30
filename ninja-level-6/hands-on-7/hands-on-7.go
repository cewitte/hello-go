package main

import "fmt"

func main() {
	a := func() {
		for i := 0.00; i <= 1000; i++ {
			fmt.Println(i)
		}
	}

	a()

	fmt.Printf("I'm type... %T\n", a)
}
