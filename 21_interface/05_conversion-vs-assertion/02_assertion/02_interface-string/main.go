package main

import "fmt"

func main() {
	var name interface{} = "Sydney"

	if str, ok := name.(string); ok {
		fmt.Printf("value: %v, type: %T\n", str, str)
	} else {
		fmt.Println("value is not a string")
	}
}
