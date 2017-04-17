package main

import "fmt"

func main() {
	fmt.Println(factorial(3))
}


func factorial(n int) int {
	re := 1

	for i := 2; i <= n; i++ {
		re *= i
	}

	return re
}
