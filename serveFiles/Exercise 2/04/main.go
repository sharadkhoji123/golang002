package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", dog)
	//	http.HandleFunc("/dog.jpg", dogP)
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/resources/", http.StripPrefix("/resources/", fs))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("Template does not exist ", err)
	}
}
