package main

import "fmt"

func main() {
	myGreetings := map[int]string{
		1: "Hello",
		2: "Xin chao",
		3: "Bonjour",
	}

	for key, value := range myGreetings {
		fmt.Println(key, "---", value)
	}
}
