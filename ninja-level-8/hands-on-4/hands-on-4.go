package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xs)

	fmt.Println("Unsorted:\t", xs)
	sort.Slice(xs, func(i, j int) bool { return xs[i] < xs[j] })
	fmt.Println("Sorted:\t\t", xs)
	fmt.Println()
	fmt.Println("Unsorted:\t", xi)
	sort.Slice(xi, func(i, j int) bool { return xi[i] < xi[j] })
	fmt.Println("Sorted:\t\t", xi)
}
