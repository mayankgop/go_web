package main

import (
	// "fmt"
	"os"
	"text/template"
)
var tpl *template.Template
func init(){
	tpl=template.Must(template.ParseFiles("d.g"))
}
func main(){
	n:=map[string]string{
		"B": "Budha",
		"G": "Gautam",
		"b": "Bose",
		"j": "jesus",
	}


	_=tpl.ExecuteTemplate(os.Stdout,"d.g",n)
}
