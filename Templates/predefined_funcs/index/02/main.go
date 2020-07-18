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
	xs := []string{"zero", "one", "two", "three", "four"}

	data := struct {
		Words []string
		Lname string
	}{
		xs,
		"Nautiyal",
	}
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()
	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
