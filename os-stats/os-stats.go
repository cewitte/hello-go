// Test exercise.
package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS Arch:\t\t", runtime.GOOS)
	fmt.Println("Number of CPUs:\t\t", runtime.NumCPU())

	wg.Add(8)
	for i := 1; i <= 8; i++ {
		go foo(0, i)
	}
	wg.Wait()
}

func foo(s, e int) {
	for i := s; i <= e; i++ {
		fmt.Println(math.Pow(2, float64(i)))
	}

	wg.Done()
}
