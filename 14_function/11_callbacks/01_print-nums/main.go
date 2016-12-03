package main

import "fmt"

func visit(nums []int, callback func(int)) {
	for _, v := range nums {
		callback(v)
	}
}

func main() {
	visit([]int{1, 2, 3, 4, 5}, func(i int) {
		fmt.Println("--", i)
	})
}
