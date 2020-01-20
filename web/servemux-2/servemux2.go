/*
	1. Take the previous program and change it so that:
	* func main uses http.Handle instead of http.HandleFunc

	Contstraint: Do not change anything outside of func main
*/

package main

import "net/http"

import "html/template"

func root(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, "the root level")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, "dog")
}

func me(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, r.RequestURI)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("servemux.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(root))
	http.Handle("/me/", http.HandlerFunc(me))
	http.Handle("/dog/", http.HandlerFunc(dog))

	http.ListenAndServe(":8080", nil)
}
