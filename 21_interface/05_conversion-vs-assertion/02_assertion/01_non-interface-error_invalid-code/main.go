package main

import "fmt"

func main() {
	name := "Sydney"

	if str, ok := name.(string); ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Println("Value is not a string\n")
	}


}
