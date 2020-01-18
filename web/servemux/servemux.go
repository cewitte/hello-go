/*

	ListenAndServe on port ":8080" using the default ServeMux.

	Use HandleFunc to add the following routes to the default ServeMux:

	"/"
	"/dog/"
	"/me/

	Add a func for each of the routes.

	Have the "/me/" route print out your name.

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
	tpl.Execute(w, "me")
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("servemux.gohtml"))
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
