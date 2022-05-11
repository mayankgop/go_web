package main

import (
	// "fmt"
	"os"
	"text/template"
)
type boss struct{
	Initial string
	Name string

}
var tpl *template.Template
func init(){
	tpl=template.Must(template.ParseFiles("e.g"))
}
func main(){
	budha:=boss{
		Initial:"G",
		Name: "gautam",
	}
	


	_=tpl.ExecuteTemplate(os.Stdout,"e.g",budha)
}
