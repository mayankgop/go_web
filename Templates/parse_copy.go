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
	nf,_:=os.Create("d.html")
	defer nf.Close()
	_=tpl.Execute(nf,nil)
}