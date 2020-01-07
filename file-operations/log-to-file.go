package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func getFilename() string {
	return "logs/" + time.Now().String() + "_log.txt"
}

func main() {
	f, err := os.Create(getFilename())
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")
}

// Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
