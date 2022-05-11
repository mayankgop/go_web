package main

import (
	// "log"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
// it will automatically call the file with name index.html just bcoz of the name
}