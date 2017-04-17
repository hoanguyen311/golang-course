package main

import "fmt"

func main() {
	myGreetings := make(map[string]string)

	myGreetings["Peter"] = "Hello"
	myGreetings["Luyen"] = "Xin chao"

	fmt.Println(myGreetings["Luyen"])
}
