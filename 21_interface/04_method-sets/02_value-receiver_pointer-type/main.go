package main

import (
	"math"
	"fmt"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c := &circle{ 100 }

	info(c)
}