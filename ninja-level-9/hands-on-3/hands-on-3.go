package main

import (
	"fmt"
	"runtime"
	"sync"
)

var increment int = 1
var wg sync.WaitGroup

func incrementerA() {
	a := increment
	runtime.Gosched()
	a++
	increment = a
	fmt.Println(a)
	wg.Done()
}

func incrementerB() {
	b := increment
	runtime.Gosched()
	b++
	increment = b
	fmt.Println(b)
	wg.Done()
}

func incrementerC() {
	c := increment
	runtime.Gosched()
	c++
	increment = c
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
