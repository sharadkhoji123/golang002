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
	sharad := gmap{
		Gcord: "Delhi",
		Name:  "Sharad",
		Motto: "Have Friends",
	}
	rahul := gmap{
		Gcord: "Vancouver",
		Name:  "Rahul",
		Motto: "Make Money",
	}
	pramod := gmap{
		Gcord: "Haryana",
		Name:  "Rahul",
		Motto: "Search Peace",
	}
	sages := []gmap{sharad, rahul, pramod}
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(nf, sages)
	if err != nil {
		log.Fatalln(err)
	}

}
