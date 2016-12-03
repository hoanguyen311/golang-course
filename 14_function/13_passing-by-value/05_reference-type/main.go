package main

import "fmt"

func main() {
	names := make([]string, 2, 15)

	changeMe(names)

	fmt.Println(names)
}

func changeMe(arr []string) {
	arr[0] = "Thai Hoa"
	arr[1] = "Cuc"
}
