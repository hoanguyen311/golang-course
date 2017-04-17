package main

import "fmt"

func main() {
	c := factorial(gen())

	for n := range c {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()

	return out
}

func factorial(nums <-chan int) <-chan int {
	c := make(chan int)

	go func() {
		for num := range nums {
			c <- fact(num)
		}
		close(c)
	}()

	return c
}

func fact(n int) int {
	re := 1
	for i := n; i > 0; i-- {
		re *= i
	}
	return re
}