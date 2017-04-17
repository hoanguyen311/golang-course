package main

import "fmt"

func main() {

	for i := range sq(sq(gen(2, 3))) {
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