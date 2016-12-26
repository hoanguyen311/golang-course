package main

import "fmt"

func main() {
	greetings := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	fmt.Println("[1:2]")
	fmt.Println(greetings[1:2])

	fmt.Println("[:2]")
	fmt.Println(greetings[:2])

	fmt.Println("[3:]")
	fmt.Println(greetings[3:])

	fmt.Println("[:]")
	fmt.Println(greetings[:])
}
