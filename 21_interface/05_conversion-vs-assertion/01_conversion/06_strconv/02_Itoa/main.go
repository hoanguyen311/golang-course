package main

import (
	"strconv"
	"fmt"
)

func main() {
	str := strconv.Itoa(100)

	fmt.Println(str)
	fmt.Printf("%T\n", str)
}
