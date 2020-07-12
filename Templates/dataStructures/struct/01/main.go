package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type gmap struct {
	Gcord string
	Name  string
	Motto string
}

func main() {

	sages := gmap{
		Gcord: "Delhi",
		Name:  "Sharad",
		Motto: "Have Friends",
	}
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(nf, sages)
	if err != nil {
		log.Fatalln(err)
	}

}
