package main

import (
	"fmt"
)

func main() {
	foo()
	bar()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("foo")
	}
}
func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("bar")
	}
}
