package main

import "fmt"

func factorial(i int) int {
	if i == 0 {
		return 1
	}

	return i * factorial(i - 1)
}

func main() {
	fmt.Println(factorial(4))
}
