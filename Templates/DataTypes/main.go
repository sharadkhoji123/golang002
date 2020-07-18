package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type secretagent struct {
	person
	licensetokill bool
}

func (p person) says() {
	fmt.Println(p.name, "says Good Morning!!! I am ", p.age, " years old")
}

func (sa secretagent) says() {
	fmt.Println(sa.name, "says Good Morning!!! I am ", sa.age, " years old and my license to kill is ", sa.licensetokill)
}

type human interface {
	says()
}

func saySomething(h human) {
	h.says()
}

func main() {
	p1 := person{
		name: "Rohan",
		age:  27,
	}

	saySomething(p1)

	sa1 := secretagent{
		person{
			"James Bond",
			32,
		},
		true,
	}
	saySomething(sa1)
}
