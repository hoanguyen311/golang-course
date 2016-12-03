package main

import "fmt"

func main() {
	sum := findSum(4000000)

	fmt.Println(sum)
}

func fibonacci(term int) int {
	if (term < 2) {
		return 1
	}
	return  fibonacci(term - 2) + fibonacci(term - 1)
}

func findSum(maxValue int) int {
	sum := 0
	counter := 0

	for valueOfTerm := fibonacci(counter);
			valueOfTerm < maxValue;
		counter, valueOfTerm = counter + 1, fibonacci(counter) {
		if valueOfTerm % 2 == 0 {
			sum += valueOfTerm
		}
	}

	return sum
}