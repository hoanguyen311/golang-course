package main

import "fmt"

func main() {
	num := 3.2

	fmt.Printf("%T\n", num)
	fmt.Printf("%T\n", int(num))

	var num2 interface{} = 20

	fmt.Printf("%T\n", num2)
	//fmt.Printf("%T\n", num2.(int))
	fmt.Printf("%T\n", int(num2))
}
