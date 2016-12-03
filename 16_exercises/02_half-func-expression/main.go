package main

import "fmt"

func main() {
	half := func(n int) (int, bool) {
		return n / 2, n % 2 == 0
	}

	theHalf, isEvent := half(9)

	fmt.Println(theHalf, isEvent)
}
