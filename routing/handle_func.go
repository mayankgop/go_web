package main

import (
	"io"
	"net/http"
)


func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}


func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {

	// func HandleFunc(pattern string, h func(ResponseWriter, *Request))


	http.HandleFunc("/dogs", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}