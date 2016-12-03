package main

import "fmt"

type Contact struct {
	name string
	greeting string
}

func switchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("Contact")
	default:
		fmt.Println("Unknown")
	}
}

func main() {
	switchOnType(10)
	switchOnType("Hoa")
	t := Contact{"Cuc", "Hello Cuc"}
	switchOnType(t)
}
