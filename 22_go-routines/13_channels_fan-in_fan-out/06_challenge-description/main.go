package main

import "fmt"

func main() {
	c := gen()
	f := factorial(c)

	for i := range f {
		fmt.Println(i)
	}
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			for j := 3; j < 13; j++ {
				c <- j
			}
		}
		close(c)
	}()

	return c
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()

	return out
}

func fact(n int) int {
	re := 1
	for i := n; i > 0; i-- {
		re *= i
	}
	return re
}