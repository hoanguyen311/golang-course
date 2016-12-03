package main

import "fmt"

var z = 200

func main() {
	x := 40

	fmt.Println("x = ", x)
	{
		fmt.Println("in side inner scope x = ", x)
		y := 100
		fmt.Println("y = ", y)
	}

	fmt.Println("z = ", z)
}
