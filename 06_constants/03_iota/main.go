package main

import "fmt"

const (
	A = iota
	B = iota
	C
)

const (
	D = iota * 10
	E = iota * 10
)

func main() {
	fmt.Printf("%v - %v - %v\n", A, B, C)
	fmt.Println(D)
	fmt.Println(E)
}
