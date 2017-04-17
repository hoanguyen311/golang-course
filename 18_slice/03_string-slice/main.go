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

	for _, g := range greetings {
		fmt.Println(g)
	}

	fmt.Println("--------------------")

	for i := 0; i < len(greetings); i++ {
		fmt.Println(greetings[i])
	}
}
