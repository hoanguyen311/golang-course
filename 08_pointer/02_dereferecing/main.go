package main

import "fmt"

func main() {
	x := 40

	fmt.Println("x = ", x)
	fmt.Println("&x = ", &x)

	var px *int = &x

	fmt.Println("px = ", px)
	fmt.Println("*px = ", *px)

	*px = *px * 10

	fmt.Println("x = ", x)
}
