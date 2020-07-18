package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", dog)
	//	http.HandleFunc("/dog.jpg", dogP)
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template doesn't exist ", err)
	}
}

// func dogP(w http.ResponseWriter, req *http.Request) {
// 	http.ServeFile(w, req, "public/pics/dog.jpg")
// }
