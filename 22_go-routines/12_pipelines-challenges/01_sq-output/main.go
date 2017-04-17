package main

import "fmt"

func main() {
	cGen := gen(2, 3, 5, 10)
	cSquare := sq(cGen)

	for i := range cSquare {
		fmt.Println(i)
	}
}

func gen(num ...int) chan int {
	c := make(chan int)

	go func() {
		for _, i := range num {
			c <- i
		}
		close(c)
	}()

	return c;
}

func sq(in chan int) chan int {
	out := make(chan int)

	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()
	return out
}