package main

import (
	"fmt"
	"sync"
)

var increment int = 1
var wg sync.WaitGroup
var mu sync.Mutex

func incrementerA() {
	mu.Lock()
	a := increment
	a++
	increment = a
	mu.Unlock()
	fmt.Println(a)
	wg.Done()
}

func incrementerB() {
	mu.Lock()
	b := increment
	b++
	increment = b
	mu.Unlock()
	fmt.Println(b)
	wg.Done()
}

func incrementerC() {
	mu.Lock()
	c := increment
	c++
	increment = c
	mu.Unlock()
	fmt.Println(c)
	wg.Done()
}

func main() {
	wg.Add(3)
	go incrementerA()
	go incrementerB()
	go incrementerC()

	wg.Wait()
}
