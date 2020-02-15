package main

import "log"

import "net/http"

func main() {

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
