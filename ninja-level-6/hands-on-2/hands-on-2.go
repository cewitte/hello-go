package main

import "fmt"

func foo(an ...int) int {
	var num int

	for v := range an {
		num += v
	}
	return num
}

func bar(an []int) int {
	var num int

	for v := range an {
		num += v
	}

	return num
}

func main() {

	fmt.Println("Foor returns...\t", foo(2, 3, 5, 7, 11, 13))

	n2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Bar returns...\t", bar(n2))

}
