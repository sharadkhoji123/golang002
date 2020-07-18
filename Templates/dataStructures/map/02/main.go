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

	sages := map[string]string{
		"Delhi":       "Sharad",
		"Haryana":     "Pramod",
		"Van Coveour": "Rahul",
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
