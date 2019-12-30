package main

import "math"

import "fmt"

type shape interface {
	area()
	info()
}

type square struct {
	base       float64
	height     float64
	squareArea float64
}

type circle struct {
	radius     float64
	circleArea float64
}

func (s *square) area() {
	s.squareArea = s.base * s.height
}

func (c *circle) area() {
	c.circleArea = math.Pi * math.Pow(c.radius, 2.00)
}

func (s *square) info() {
	fmt.Println("I am a square and my area is", s.squareArea)
}

func (c *circle) info() {
	fmt.Println("I am a circle and my area is", c.circleArea)
}

func main() {
	s := square{
		base:   2.00,
		height: 2.00,
	}

	c := circle{
		radius: 2.00,
	}

	s.area()
	c.area()

	s.info()
	c.info()

}
