package main

import "fmt"

func filter(nums []int, callback func(int) bool) []int {
	var result []int

	for _, v := range nums {
		if (callback(v)) {
			result = append(result, v)
		}
	}

	return result
}

func main() {
	filtered := filter([]int{1, 4, 10, 3, 5}, func(i int) bool {
		return i % 5 == 0
	})

	fmt.Println(filtered)
}
