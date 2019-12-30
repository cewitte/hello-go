package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v %v and I am %v yearls old.", p.first, p.last, p.age)
	fmt.Println()
}

func main() {
	p1 := person{
		first: "Carlos",
		last:  "Witte",
		age:   47,
	}

	p1.speak()
}
