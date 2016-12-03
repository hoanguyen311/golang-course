package main

import "fmt"

func main() {
	greet("Nguyen", "Thai Hoa")
	greet("Phan", "Thi Cuc")
}

func greet(lastName , firstName string) {
	fmt.Println("Bonjour", firstName, lastName)
}
