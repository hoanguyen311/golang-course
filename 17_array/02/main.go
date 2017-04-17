package main

import "fmt"

func main() {
	var arr [256]int

	for i := 0; i < 256; i++ {
		arr[i] = i
	}

	fmt.Println(arr)

	for _, v := range arr {
		fmt.Printf("%v %T %b\n", v, v, v)
		if v > 50 {
			break;
		}
	}
}
