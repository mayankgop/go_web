package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}
// ServeHTTP writes reply headers and data to ResponseWriter and then return
// since ServeHTTP method is attached with method 'm'(of type hotdog) so 'm' is also a handler

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
	// here d is our handler
	// A Handler responds to an http request
	// ListenAndServe asks us for the address and the handler 
}




// type Handler interface{
//     SeverHTTP(ResponseWriter, *Request)
// }
