package main

import "fmt"

func main() {
	x := 10

	fmt.Printf("X: %d\n", x)
	fmt.Println("X's memory address: ", &x)
	fmt.Printf("X's memory address in decimal: %d\n", &x)
}
