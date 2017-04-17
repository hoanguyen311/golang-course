package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	area() float64
}

func info(s Shape) {
	fmt.Println(s.area())
}

func main() {
	s1 := Square{10}
	c1 := Circle{9.2}

	info(s1)
	info(c1)
}
