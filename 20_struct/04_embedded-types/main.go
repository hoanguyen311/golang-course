package main

import "fmt"

type Person struct {
	First string
	Last string
	Age int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last: "Bone",
			Age: 30,
		},
		LicenseToKill: true,
	}
	p2 := Person{
		First: "Thai Hoa",
		Last: "Nguyen",
	}
	//p1 := DoubleZero{
	//	Person{ "James", "Bone", 30 },
	//	true,
	//}

	fmt.Println( p1.First, p1.Last, p1.LicenseToKill)
	fmt.Println( p2.First, p2.Last )
}
