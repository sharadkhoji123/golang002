package main

import (
	"log"
	"os"
	"text/template"
)

type hotels struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h1 := []hotels{
		{
			Name:    "Royal Blue",
			Address: "25, Southern street Downtown",
			City:    "California",
			Zip:     "27001",
			Region:  "Southern",
		},
		{
			Name:    "La Ville",
			Address: "2, CA Street 5th Cross",
			City:    "California",
			Zip:     "27005",
			Region:  "Central",
		},
		{
			Name:    "Westend",
			Address: "4, Mountain Road",
			City:    "California",
			Zip:     "27009",
			Region:  "Northern",
		},
	}

	np, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer np.Close()
	err = tpl.Execute(np, h1)
	if err != nil {
		log.Fatalln(err)
	}
}
