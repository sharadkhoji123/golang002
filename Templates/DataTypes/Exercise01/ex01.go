package main

import (
	"fmt"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() {
	area := s.side * s.side
	fmt.Println("Square area is ", area)
}

func (c circle) area() {
	area := 3.14 * c.radius * c.radius
	fmt.Println("Circle area is ", area)
}

type shape interface {
	area()
}

func info(s shape) {
	s.area()
}

func main() {
	c1 := circle{5.0}
	s1 := square{4}
	info(c1)
	info(s1)
}
