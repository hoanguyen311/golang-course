package main

import "fmt"

func main() {
	fmt.Println(foo(1, 2))
	fmt.Println(foo(10, 10, 10))
	nums := []int{1, 2, 3}

	fmt.Println(foo(nums...))

	fmt.Println(foo())
}

func foo(nums ...int) int {
	var sum int

	for _, val := range nums {
		sum += val
	}

	return sum
}

