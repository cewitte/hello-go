package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOARCH =\t", runtime.GOARCH)
	fmt.Println("Op. System =\t", runtime.GOOS)
	fmt.Println("Go version =\t", runtime.Version())
	fmt.Println("Num of CPUs =\t", runtime.NumCPU())
	fmt.Println("Go Root =\t", runtime.GOROOT())
}
