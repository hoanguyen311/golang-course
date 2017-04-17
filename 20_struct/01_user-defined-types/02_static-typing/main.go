package main

import "fmt"

type foo int

func main() {
	var myAge foo

	myAge = 26

	fmt.Printf("%T %v\n", myAge, myAge)

	var yourAge int

	yourAge = 22

	fmt.Printf("%T %v\n", yourAge, yourAge)

	//fmt.Println(myAge + yourAge)
	fmt.Println(int(myAge) + yourAge)
}
