package main

import "fmt"

func main() {
	myMap := map[int]string{
		1: "Thai Hoa",
		2: "Thi Cuc",
		3: "Rest of the world",
	}

	fmt.Println(myMap[3])
}
