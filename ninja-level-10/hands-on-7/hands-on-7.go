package main

import "fmt"

func main() {
	c := make(chan int)

	for i := 1; i <= 10; i++ {
		go func() {
			for j := 1; j <= 10; j++ {
				c <- j
			}
		}()
	}

	for i := 1; i <= 100; i++ {
		fmt.Println(<-c)
	}
}
