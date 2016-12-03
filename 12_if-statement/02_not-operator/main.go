package main

import "fmt"

func main() {
	if !true {
		fmt.Println("This code will not run")
	}
	if !false {
		fmt.Println("This code will run")
	}
}
