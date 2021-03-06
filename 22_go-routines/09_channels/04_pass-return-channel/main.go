package main

import "fmt"

func main() {
	inc := incrementor()
	cSum := puller(inc)

	for i := range cSum {
		fmt.Println(i)
	}
}

func incrementor() chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func puller(inc chan int) chan int {
	c := make(chan int)

	go func() {
		var sum int
		for i := range inc {
			sum += i
		}
		c <- sum

		close(c)
	}()

	return c
}
