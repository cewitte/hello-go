package main

import "fmt"

func getF() func() {
	return func() {
		fmt.Println("I am the returned func")
	}
}

func main() {
	f := getF()

	f()
}
