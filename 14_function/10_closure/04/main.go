package main

import "fmt"

func wrapper() func() int {
	var y int
	return func() int {
		y++
		return y
	}
}

func main() {
	increment := wrapper()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
