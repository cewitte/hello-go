package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, "This is from dog")
}

func doggy(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpg")
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", doggy)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalln("Oh shit")
		return
	}

}
