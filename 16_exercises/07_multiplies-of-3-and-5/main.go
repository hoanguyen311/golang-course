package main

import "fmt"

func main() {

	fmt.Println(findSum(1000))
}

func findSum(max int) int {
	sum := 0

	for i := max - 1; i > 0; i-- {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	return  sum
}
