package main

import "fmt"

func main() {
	var num interface{} = 100

	fmt.Println(num.(int) + 10)
}
