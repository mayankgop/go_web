package main

import (
	"fmt"
	"os"
	"text/template"
)
func main(){

	tpl,_:=template.ParseFiles("a.g")
	fmt.Println()

	_=tpl.Execute(os.Stdout,nil)
	fmt.Println()
	tpl,_=tpl.ParseFiles("b.g","c.g")
	_=tpl.ExecuteTemplate(os.Stdout,"c.g",nil)
	fmt.Println()

	_=tpl.ExecuteTemplate(os.Stdout,"b.g",nil)
	fmt.Println()

	_=tpl.ExecuteTemplate(os.Stdout,"a.g",nil)
	fmt.Println()

	_=tpl.Execute(os.Stdout,nil)// it will execute which it will find first
	

}