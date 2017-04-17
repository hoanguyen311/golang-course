package main

import "fmt"

func main() {
	myGreetings := map[int]string{
		1: "Hello",
		2: "Xin chao",
		3: "Bonjour",
	}

	delete(myGreetings, 2)

	fmt.Println(myGreetings)
}
