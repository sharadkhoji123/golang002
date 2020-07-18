package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func double(x int) int {
	return x + x
}
func square(x int) float64 {
	return float64(x * x)
}
func tvalue(t time.Time) string {
	return t.Format("02-02-2006")
}

var fm = template.FuncMap{
	"fdbl": double,
	"fsqr": square,
	"tval": tvalue,
}

func main() {
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()
	data := struct {
		dbl int
		sqr float64
		tv  time.Time
	}{
		dbl: 5,
		sqr: 6,
		tv:  time.Now(),
	}

	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
