package main

import "fmt"

func main() {
	myGreetings := map[int]string{
		1: "Hello",
		2: "Xin chao",
		3: "Bonjour",
	}
	keyToDelete := 3

	fmt.Println(myGreetings)
	if _, ok := myGreetings[keyToDelete]; ok {
		delete(myGreetings, keyToDelete)
		fmt.Println(ok)
	} else {
		fmt.Println("There's no entry at ", keyToDelete)
	}

	fmt.Println(myGreetings)
}
