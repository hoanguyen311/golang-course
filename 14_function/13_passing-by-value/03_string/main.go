package main

import "fmt"

func main() {
	name := "Hoa"
	fmt.Println(name)
	changeMe(name)
	fmt.Println(name)
}


func changeMe(z string) {
	z = "Nguyen Thai Hoa"
}
