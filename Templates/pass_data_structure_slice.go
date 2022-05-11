package main

import (
	// "fmt"
	"os"
	"text/template"
)
var tpl *template.Template
func init(){
	tpl=template.Must(template.ParseFiles("a.g"))
}
func main(){
	n:=[]string{"bose","maharana","budha","ambedkar"}
	_=tpl.ExecuteTemplate(os.Stdout,"a.g",n)
}
