package main

import (
	"os"
	"text/template"
	"time"
)

var tpl *template.Template
func init(){
	tpl=template.Must(template.New("").Funcs(fm).ParseFiles("f.g"))
}
func mdy(t time.Time)string{
	return t.Format("01-02-2020")
}
var fm = template.FuncMap{
	"fd":mdy,
}
func main(){
	_=tpl.ExecuteTemplate(os.Stdout,"f.g",time.Now())
}