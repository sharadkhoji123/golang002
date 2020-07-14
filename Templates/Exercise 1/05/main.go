package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("table.csv"))
}

type data struct {
	Date time.Time
	High float64
}

func main() {
	np, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(np, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
