package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println("Before change: ", m)
	changeMe(m)
	fmt.Println("After change: ", m["Hoa"])
}

func changeMe(z map[string]int) {
	z["Hoa"] = 1
}
