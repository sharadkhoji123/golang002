package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p1 := person{
		Name: "Rohan",
		Age:  32,
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(nf, p1)
	if err != nil {
		log.Fatalln(err)
	}

}
