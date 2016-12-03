package main

import "fmt"

func makeGreater() func() string {
	return func() string {
		return "Hello"
	}
}

func main() {
	greet := makeGreater()

	fmt.Println(greet())
	fmt.Printf("%T\n", greet)
}
