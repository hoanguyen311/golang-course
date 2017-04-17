package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) fullName() string {
	return fmt.Sprintf("%v %v", p.first, p.last)
}

func main() {
	p1 := person{"Hoa", "Nguyen Thai", 26}
	p2 := person{"Cuc", "Phan Thi", 26}

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
