package main

import "fmt"

func main() {
	nums := []float64{20, 10, 15, 100}

	a := average(nums...)

	fmt.Println(a)
}

func average(nums ...float64) float64 {
	var result float64

	for _, v := range nums {
		result += v
	}

	return result / float64(len(nums))
}
