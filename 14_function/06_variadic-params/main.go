package main

import "fmt"

func main() {
	a := average(10, 20)

	fmt.Println(a)
}

func average(nums ...float64) float64 {
	var total float64
	for _, v := range nums {
		total += v
	}
	return total / float64(len(nums))
}
