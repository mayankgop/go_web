package main

import (
	// "fmt"
	"os"
	"strings"
	"text/template"
)

type boss struct {
	Initial string
	Name    string
}

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s

}

func init() {
	tpl = template.Must(template.New("mayank").Funcs(fm).ParseFiles("e.g"))
}

func main() {
	budha := boss{
		Initial: "Budha",
		Name:    "gautam",
	}

	_ = tpl.ExecuteTemplate(os.Stdout, "e.g", budha)
}
