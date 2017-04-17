package main

import "fmt"

func main() {
	x := 10;

	s := evalInt(x)

	fmt.Println(s)
}

func evalInt(n int) string {
	if (n > 10) {
		return fmt.Sprintf("%d is greater than 10", n)
	} else {
		return fmt.Sprintf("%d is not greater than 10", n)
	}
}
