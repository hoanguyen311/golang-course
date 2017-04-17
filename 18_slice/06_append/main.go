package main

import "fmt"

func main() {
	greetings := make([]string, 3, 5)

	greetings[0] = "Hello"
	greetings[1] = "Hi"
	greetings[2] = "Xin chao"
	//greetings[3] = "bonjour"

	greetings = append(greetings, "bonjour")
	greetings = append(greetings, "gidday")
	greetings = append(greetings, "Suprabadham")

	fmt.Println(greetings)
	fmt.Println(cap(greetings))
}
