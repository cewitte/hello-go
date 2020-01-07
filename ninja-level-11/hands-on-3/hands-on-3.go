package main

import "fmt"

type customErr struct {
}

func (err *customErr) Error() string {
	return "Oh my god! How could you? Some mistakes are never forgiven."
}

func foo(err error) {
	fmt.Printf("Error: %v\n", err)
}

func main() {
	err := customErr{}
	foo(&err)
}
