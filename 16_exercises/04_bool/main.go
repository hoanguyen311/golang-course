package main

import "fmt"

func main() {
	if (true && false) || (false && true) || !(false && false) {
		fmt.Println("This ran")
	}
}
