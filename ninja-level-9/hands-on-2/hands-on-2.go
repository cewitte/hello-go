package main

import "fmt"

type person struct {
	name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println(p.name)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		name: "Amanda Miranda Witte",
	}

	p.speak()

	saySomething(&p)

}
