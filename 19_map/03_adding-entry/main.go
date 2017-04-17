package main

import "fmt"

func main() {
	//myMap := map[string]string{
	//	"Hoa": "FE Dev",
	//	"Cuc": "German teacher",
	//}
	//
	var myMap = make(map[string]int)

	fmt.Println(myMap)

	myMap["Hieu"] = 10

	fmt.Println(myMap)

	fmt.Println(myMap["Hic"])

}
