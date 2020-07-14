package main

import (
	"log"
	"os"
	"text/template"
)

type meals struct {
	Breakfast string
	Lunch     string
	Dinner    string
}

type restaurant struct {
	Name  string
	Mtype []meals
}

type Restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r := Restaurants{
		restaurant{
			Name: "Mirchi Hotel",
			Mtype: []meals{
				meals{
					Breakfast: "Idli Dosa",
					Lunch:     "Rajma Chawal",
					Dinner:    "Shahi Paneer",
				},
				meals{
					Breakfast: "Sambhar Vada",
					Lunch:     "Kadi Pakoda",
					Dinner:    "Mix Veg",
				},
			},
		},
		restaurant{
			Name: "Jogi Dhaba",
			Mtype: []meals{
				meals{
					Breakfast: "Chole Bhature",
					Lunch:     "Kadhi Chawal",
					Dinner:    "Mix Veg",
				},
				meals{
					Breakfast: "Sandwich/Shakes",
					Lunch:     "Thali/Chinese",
					Dinner:    "Shahi Paneer",
				},
			},
		},
	}
	np, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer np.Close()
	err = tpl.Execute(np, r)
	if err != nil {
		log.Fatalln(err)
	}

}
