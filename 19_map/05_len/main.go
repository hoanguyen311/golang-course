package main

import "fmt"

func main() {
	myGreetings := map[string]string {
		"en": "Hello",
		"vi": "Xin chao",
		"fr": "Bounjour",
	};

	fmt.Println(len(myGreetings))
}
