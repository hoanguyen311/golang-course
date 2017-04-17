package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	xc := fanOut(in, 10)

	for n := range merge(xc...) {
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
	}()

	return out
}

func fanOut(in <-chan int, n int) []<-chan int {
	xc := make([]<-chan int, n)

	for i := 0; i < n; i++ {
		xc = append(xc, factorial(in))
	}

	return  xc
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)

	fact := func(n int) int {
		re := 1;
		for i := n; i > 0; i-- {
			re *= i
		}
		return re;
	}

	go func() {
		for i := range in {
			out <- fact(i)
		}
		close(out)
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	output := func(c <-chan int) {
		for i := range c {
			out <- i
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