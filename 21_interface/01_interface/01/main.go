package main

import "fmt"

//Square Something
type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

type Shape interface {
	area() float64
}

func info(s Shape) {
	fmt.Println(s.area())
}

func main() {
	s1 := Square{10}

	info(s1)
}
