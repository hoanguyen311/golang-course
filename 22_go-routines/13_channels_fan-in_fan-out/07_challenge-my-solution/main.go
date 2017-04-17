package main

import (
	"fmt"
	"sync"
)

func main() {

	c := gen()
	f1 := factorial(c)
	f2 := factorial(c)
	f3 := factorial(c)
	f4 := factorial(c)
	f5 := factorial(c)
	f6 := factorial(c)
	f7 := factorial(c)
	f8 := factorial(c)



	for i := range merge(f1, f2, f3, f4, f5, f6, f7, f8) {
		fmt.Println(i)
	}
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100000; i++ {
			for j := 30; j < 40; j++ {
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

func merge(cs ...<-chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(cs))

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	go func() {
		for _, c := range cs {
			output(c)
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}