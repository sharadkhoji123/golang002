package main

import (
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.New("Check").Parse("What's up Doug!!!"))
	tpl.ExecuteTemplate(os.Stdout, "Check", nil)
}
