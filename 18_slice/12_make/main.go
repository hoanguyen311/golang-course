package main

import "fmt"

func main() {

	mySlice := make([]string, 30)

	mySlice[0] = "Thai Hoa"
	mySlice[29] = "AAA"

	fmt.Println(mySlice)
}
