package main

import (
	// "fmt"
	"os"
	"text/template"
)
var tpl *template.Template
func init(){
	tpl=template.Must(template.ParseFiles("b.g"))
}
func main(){
	_=tpl.ExecuteTemplate(os.Stdout,"b.g","yes it is it that i passed")
}
