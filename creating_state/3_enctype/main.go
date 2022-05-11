package main

import (
	"text/template"

	// "log"
	"net/http"
	// "io/ioutil"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	// var s string

	// body
	var body string
	if req.Method == http.MethodPost {

		bs := make([]byte, req.ContentLength)
		req.Body.Read(bs)
		body = string(bs)
		_ = tpl.ExecuteTemplate(w, "index.gohtml", body)

	} else {
		tpl.ExecuteTemplate(w, "index.gohtml", nil)
	}

}
