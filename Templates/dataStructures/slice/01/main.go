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

func main() {

	sages := []string{"Sharad", "Pramod", "Rahul"}
	nf, err := os.Create("index1.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(nf, sages)
	if err != nil {
		log.Fatalln(err)
	}

}
