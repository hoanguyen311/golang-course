package main

import "fmt"

func main() {
	x := 40

	fmt.Println("x = ", x);
	fmt.Println("&x = ", &x);

	var px *int = &x

	fmt.Println("px = ", px);
}
