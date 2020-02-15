package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, go! This is my Twitter account: @cewitte")
	fmt.Println(time.Now())
	t := time.Now()
	fmt.Printf("time:%v\n", t)
}
