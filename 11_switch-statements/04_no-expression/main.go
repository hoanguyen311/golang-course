package main

import "fmt"

func main() {
	name := "Ha"

	switch {
	case len(name) == 2:
		fmt.Println("Your friend's name has length 2")
	case name == "Hoa":
		fmt.Println("Wassup N T Hoa")
	case name == "Cuc":
		fmt.Println("Wassup P T Cuc")
	default:
		fmt.Println("Nothing matched, default case")
	}
}
