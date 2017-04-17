package main

import "fmt"

func main() {
	var mySlice []string
	var mySlices [][]string

	fmt.Println(mySlice)
	fmt.Println(mySlices)

	fmt.Println(append(mySlice, "Hoa"))


	fmt.Println(mySlice == nil)
}
