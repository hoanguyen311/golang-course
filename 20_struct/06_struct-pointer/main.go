package main

import "fmt"

type Person struct {
	last string
	first string
	age int
}

func main() {
	p1 := &Person{
		last: "Nguyen",
		first: "Hoa",
		age: 26,
	}

	fmt.Println(p1)
	fmt.Printf("%T \n", p1)
	fmt.Println(p1.first, p1.last, p1.age)
}
