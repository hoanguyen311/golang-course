package main

import "fmt"

func main() {
	var name interface{} = 70

	if str, ok := name.(string); ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Println("That is not a string")
	}
}
