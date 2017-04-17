package main

import "fmt"

func main() {
	str := "Ab"
	letter := rune(str[0])
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)
}
