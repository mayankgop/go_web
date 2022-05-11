package main

import (
	"fmt"
	"os"
	"io"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println("first argument is name of file i.e. ",os.Args[0])
	fmt.Println("          ")
	fmt.Println(os.Args[1])

	abc := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en"
		<head>
		<title>Hello sample</title>
		</head>
		<body>
		<h1>` + name + `</h1>
		</body>
		</html>
	     `)
	f,_:=os.Create("u.html")
	defer f.Close()
	io.Copy(f,strings.NewReader(abc))

}
