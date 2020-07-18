package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretagent struct {
	person
	shooter bool
}

func (p person) sayWhat() {
	fmt.Println("My name is ", p.fname, p.lname)
}

func (sa secretagent) sayWhat() {
	fmt.Println("My name is ", sa.fname, sa.lname, ", my shooter status is ", sa.shooter)
}

type human interface {
	sayWhat()
}

func speak(h human) {
	h.sayWhat()
}

func main() {
	p1 := person{
		fname: "Ron",
		lname: "Gaurd",
	}
	sa1 := secretagent{
		person{
			"Bon",
			"Jovi",
		},
		true,
	}

	speak(p1)
	speak(sa1)
}
