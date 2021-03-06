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
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()
	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", 18)
	if err != nil {
		log.Fatalln(err)
	}

}
