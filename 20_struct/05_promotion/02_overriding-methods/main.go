package main

import "fmt"

type Person struct {
	first string
	last string
	age int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func (p Person) Greetings() {
	fmt.Println("Hello, I am a regular person")
}

func (dz DoubleZero) Greetings() {
	fmt.Println("Miss Moneypenny, so good to see you")
}

func main() {
	p1 := Person{"Hoa", "Nguyen Thai", 26}
	p2 := DoubleZero{
		Person: Person{
			"Bone",
			"James",
			30,
		},
		LicenseToKill: false,
	}

	p2.Greetings()
	p1.Greetings()
	p2.Person.Greetings()
}
