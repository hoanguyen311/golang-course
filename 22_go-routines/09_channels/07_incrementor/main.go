package main

import "fmt"

func main() {
	inc1 := incrementor("Foo: ")
	inc2 := incrementor("Bar: ")

	cSum1 := puller(inc1)
	cSum2 := puller(inc2)

	fmt.Println("sum Foo + Bar", <-cSum1 + <-cSum2)

}

func incrementor(name string) <-chan int {
	c := make(chan int)

	go func () {
		for i := 0; i < 10; i++ {
			fmt.Println(name, i)
			c <- i
		}
		close(c)
	}()

	return c
}

func puller(inc <-chan int) <- chan int {
	c := make(chan int)

	go func() {
		var sum int
		for  i := range inc {
			sum += i
		}
		c <- sum
		close(c)
	}()

	return c
}
