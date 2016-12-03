package main

import "fmt"

func main() {
	name := "Cuc"

	switch name {
	case "Cuc", "Hoa":
		fmt.Println("Wassup Cuc or Hoa")
	case "Binh":
		fmt.Println("Wassup Binh")
	}
}
