package main

import "fmt"

func main() {
	theHalf, isEvent := half(9)

	fmt.Println(theHalf, isEvent)
}

func half(n float64) (float64, bool) {
	return n / 2, int(n) % 2 == 0
}