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

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CS20", "Introduction to C Programming", "4"},
				course{"CS121", "Introduction to Java Programming", "20"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"CS31", "Advanced Go", "3"},
				course{"CS91", "Advance Web Programming", "12"},
			},
		},
	}

	np, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer np.Close()
	err = tpl.Execute(np, y)
	if err != nil {
		log.Fatalln(err)
	}
}
