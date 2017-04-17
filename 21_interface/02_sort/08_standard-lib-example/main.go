package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type ByAge []person

func (people ByAge) Less(i, j int) bool {
	return people[i].age < people[j].age
}

func (people ByAge) Swap(i, j int) {
	people[i], people[j] = people[j], people[i]
}

func (people ByAge) Len() int {
	return len(people)
}
func (p person) String() string {
	return fmt.Sprintf("HAHA ! name: %v, age: %v", p.name, p.age)
}
func main() {
	people := ByAge{
		person{"Hoa", 26},
		person{"Cuc", 22},
		person{"Nam", 30},
		person{"Thao", 32},
		person{"Ha", 34},
	}

	fmt.Println(people[0])
}
