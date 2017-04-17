package main

import "fmt"

func main() {
	var arr [58]string

	fmt.Println(arr);

	for i:= 65; i <= 122; i++ {
		arr[i-65] = string(i)
	}

	fmt.Println(arr)
	fmt.Println(arr[25])
}
