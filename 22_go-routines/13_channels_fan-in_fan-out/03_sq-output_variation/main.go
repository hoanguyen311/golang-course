package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3, 4, 5, 10)

	c1 := sq(in)
	c2 := sq(in)

	for c := range merge(c1, c2) {
		fmt.Println(c)
	}
}

func gen(nums ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, n := range nums {
			c <- n
		}
		close(c)
	}()

	return c
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)

	calc := func(num int) int {
		return num * num
	}

	go func() {
		for n := range in {
			out <- calc(n)
		}
		close(out)
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	wg.Add(len(cs))
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
