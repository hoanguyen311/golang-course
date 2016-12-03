package main

import "fmt"

func main() {
	greet("Thai Hoa")
	greet("Thi Cuc")
}


func greet(name string) {
	fmt.Println("Hello", name)
}