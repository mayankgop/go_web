package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	name := "Mayank Madan"
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
	f,_:=os.Create("t.html")
	defer f.Close()
	io.Copy(f,strings.NewReader(abc))
}
