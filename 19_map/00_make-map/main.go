package main

import "fmt"

func main() {
	var myMap = make(map[string]string)

	myMap["Thai Hoa"] = "Hello"
	myMap["Thi Cuc"] = "Hi"

	fmt.Println(myMap)
}
