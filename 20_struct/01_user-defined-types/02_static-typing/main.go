package main

import "fmt"

type person struct {
	firstName string
	lastName string
	age int
}

func main() {
	p1 := person{"Hoa", "Nguyen", 26}
	p2 := person{"Cuc", "Phan", 22}
	p3 := person{}

	fmt.Println(p1.firstName, p1.lastName, p1.age)
	fmt.Println(p2.firstName, p2.lastName, p2.age)
	fmt.Println(p3.age)
}
