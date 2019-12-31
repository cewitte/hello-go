package main

import "fmt"

type person struct {
	first string
}

func (p *person) changeMe() {
	p.first = "James"
}

func main() {
	carlos := person{
		first: "Carlos",
	}

	fmt.Println(carlos.first)

	carlos.changeMe()

	fmt.Println(carlos.first)
}
