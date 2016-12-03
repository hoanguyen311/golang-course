package main

import "fmt"

func findSum(num int) int {
	var sum int

	for i:= 0; i < num; i++ {
		if (i % 3 == 0 || i % 5 == 0) {
			sum += i;
		}
	}

	return sum
}

func main() {
	fmt.Println(findSum(1000))
}
