package main

import "fmt"

func sayIt(str string) {
	fmt.Println(str)
}

func main() {
	defer sayIt("1) Hello")
	sayIt("2) World")
}
