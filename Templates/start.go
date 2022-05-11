package main

import "fmt"

func main() {
	name := "Mayank Madan"
	abc := `
		<!DOCTYPE html>
		<html lang="en"
		<head>
		<title>Hello sample</title>
		</head>
		<body>
		<h1>` + name + `</h1>
		</body>
		</html>
	     `
	fmt.Println(abc)

}
