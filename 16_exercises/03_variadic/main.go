package main

import "fmt"

func main() {
	max := findMax(2, 24, 45, 3, 10, 100)

	fmt.Println("Max:", max)
}

func findMax(nums ...int) int {
	max := nums[0]

	for _, val := range nums {
		if val > max {
			max = val
		}
	}

	return max
}