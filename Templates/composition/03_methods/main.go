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

type person struct {
	Name string
	Age  int
}

func (p person) SomeProcessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakeArgs(x int) int {
	return x * 2
}

func main() {
	p1 := person{
		Name: "Will",
		Age:  42,
	}

	np, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(np, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
