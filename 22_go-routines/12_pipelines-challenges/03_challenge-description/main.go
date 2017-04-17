package main

import "fmt"

func main() {
	c := factorial(3)

	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	c := make(chan int)

	go func() {
		re := 1
		for i := n; i > 0; i-- {
			re *= i
		}

		c <- re

		close(c)
	}()

	return c
}
