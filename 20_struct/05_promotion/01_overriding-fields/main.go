package main

import "fmt"

type Person struct {
	first string
	last string
	age int
}

type DoubleZero struct {
	Person
	first string
}


func main() {
	p1 := Person{"Hoa", "Nguyen Thai", 26}
	p2 := DoubleZero{
		Person: Person{
			"Hoa",
			"James",
			30,
		},
		first: "Bone",
	}

	fmt.Println(p1.first)
	fmt.Println(p2.first)
}
