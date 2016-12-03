package main

import "fmt"

func main() {
	a := 26
	fmt.Println("&a = ", &a, ", a = ", a)

	changeMe(&a)

	fmt.Println(a)
}

func changeMe(a *int) {
	fmt.Println("a - inside func", a)
	*a = 30
}
