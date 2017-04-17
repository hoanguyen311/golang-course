package main

import "fmt"

func main() {
	fmt.Println(<-factorial(3))
}

func factorial(n int) chan int {
	c := make(chan int)
	go func() {
		re := 1
		for i := range incrementor(n) {
			re *= i
		}
		c <- re
		close(c)
	}()

	return c
}

func incrementor(n int) chan int {
	c := make(chan int)

	go func() {
		for i := 2; i <= n; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}