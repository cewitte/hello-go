package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func(count int) {
		for i := 1; i <= count; i++ {
			c <- i
		}
		close(c)
	}(100)

	for v := range c {
		fmt.Print(v, "\t")
	}

	fmt.Println("\n\nAbout to exit...")
}
