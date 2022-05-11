package main

import (
	// "log"
	"os"
	"text/template"
)
func main(){
	tpl,_:=template.ParseFiles("tpl.gohtml")
	// if err!=nil{
	// 	log.Fatalln(err)
	// }
	_=tpl.Execute(os.Stdout,nil)
}