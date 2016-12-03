package main

import "fmt"



func main() {
	x := 100

	if x % 2 == 0 {
		fmt.Printf("%d is even number\n", x)
	} else {
		fmt.Printf("%d is odd number\n", x)
	}
}
