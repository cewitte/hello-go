package main

import "fmt"

func mathOp(x, y float64, s string, a func(a, b float64) float64) {
	fmt.Printf("Function: %v. Result: %v\n", s, a(x, y))
}

func main() {
	add := func(a, b float64) float64 {
		return a + b
	}

	mult := func(a, b float64) float64 {
		return a * b
	}

	x := 12.5
	y := 36.2

	mathOp(x, y, "add", add)
	mathOp(x, y, "mult", mult)

}
