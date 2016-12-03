package main

import "fmt"

func main() {
	age := 26

	fmt.Println(age)

	changeMe(age)

	fmt.Println(age)
}

func changeMe(z int) {
	z = 28
}
