package main

import "fmt"

const (
	printing  = iota
	paper     = iota
	finishing = iota
)

func main() {
	fmt.Println(printing)
	fmt.Println(paper)
	fmt.Println(finishing)
}
