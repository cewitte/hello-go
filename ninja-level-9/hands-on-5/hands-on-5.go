package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var increment int64 = 1
var wg sync.WaitGroup

func incrementerA() {
	atomic.AddInt64(&increment, 1)
	runtime.Gosched()
	a := atomic.LoadInt64(&increment)
	fmt.Println(a)
	wg.Done()
}

func incrementerB() {
	atomic.AddInt64(&increment, 1)
	runtime.Gosched()
	b := atomic.LoadInt64(&increment)
	fmt.Println(b)
	wg.Done()
}

func incrementerC() {
	atomic.AddInt64(&increment, 1)
	runtime.Gosched()
	c := atomic.LoadInt64(&increment)
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
