package main

import (
	"fmt"
	"math"
	"sync"
)

var wg sync.WaitGroup

func rePow(x, y, count int) {
	fmt.Printf("\nGo Routine %v result: %v \n", count, math.Pow(float64(x), float64(y)))
	wg.Done()
}

func main() {
	wg.Add(3)
	go rePow(2, 4, 1)
	go rePow(3, 5, 2)
	go rePow(4, 6, 3)
	wg.Wait()
}
