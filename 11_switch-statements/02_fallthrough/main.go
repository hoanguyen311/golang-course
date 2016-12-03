package main

import "fmt"

func main() {
	name := "Cucdd"

	switch name {
	case "Cuc":
		fmt.Println("Wassup Cuc")
		fallthrough
	case "Hoa":
		fmt.Println("Wassup Hoa")
		fallthrough
	case "Binh":
		fmt.Println("Wassup Binh")
	default:
		fmt.Println("wassup default")
	}
}