package main

import (
	"fmt"
)

func main() {
	done := make(chan bool)
	c := make(chan int)

	go func() {
		for i := 0; i < 9; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 9; i++ {
			c <- i
		}
		done <- true
	}()

	<-done
	<-done
	close(c)

	for i := range c {
		fmt.Println(i)
	}

}
