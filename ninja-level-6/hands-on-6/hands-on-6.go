package main

import (
	"fmt"
)

func main() {

	func() {
		for i := 0.00; i <= 1000; i++ {
			fmt.Println(i)
		}
	}()

	fmt.Println("Done")

}
