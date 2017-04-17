package main

import (
	"strconv"
	"fmt"
)

func main() {

	if i, err := strconv.Atoi("100"); err == nil {
		fmt.Println(i)
		fmt.Printf("%T\n", i)
	}
}