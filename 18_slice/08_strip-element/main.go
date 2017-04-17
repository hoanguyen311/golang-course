package main

import "fmt"

func main() {
	mySlice := []string{"Sunday", "Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(append(mySlice[1:], myOtherSlice[:len(myOtherSlice) - 1]...))
}
