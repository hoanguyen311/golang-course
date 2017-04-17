package main

import "fmt"

func main() {
	str := "Hello"
	letter := rune(str[1])
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)
}
