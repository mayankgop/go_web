package main

import (
	// "fmt"
	"os"
	"text/template"
)
var tpl *template.Template
func init(){
	tpl=template.Must(template.ParseFiles("c.g"))
}
func main(){
	_=tpl.ExecuteTemplate(os.Stdout,"c.g",42)
}
