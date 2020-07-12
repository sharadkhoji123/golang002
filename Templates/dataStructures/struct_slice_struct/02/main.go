package main

import (
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	a := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	b := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	c := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	d := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	e := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{a, b, c}
	cars := []car{d, e}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(nf, data)
	if err != nil {
		log.Fatalln(err)
	}
}
