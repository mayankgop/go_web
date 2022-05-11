package main

import (
	"example/urlshort"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	mux := http.NewServeMux() //created a new multiplexer
	mux.HandleFunc("/", hello)
	paths := map[string]string{
		"/youtube": "https://youtu.be/wi-MGFhrad0",
		"/gmail":   "https://youtu.be/g6zYF7fbJk8",
	}
	maphandler := urlshort.MapHandler(paths, mux)

	
	yaml, err := ioutil.ReadFile("./a.yaml")
	if err != nil {
		fmt.Println(err)
	}
	
	yamlHandler, err := urlshort.YAMLHandler(yaml, maphandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on port 8080 -----")
	http.ListenAndServe(":8080", yamlHandler)

}

func hello(w http.ResponseWriter, res *http.Request) {
	fmt.Fprint(w, "this is hello")
}
