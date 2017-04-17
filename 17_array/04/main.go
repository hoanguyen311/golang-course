package main

import "fmt"

func main() {
	var arr [256]string

	for i := 0; i < 256; i++ {
		arr[i] = string(i)
	}

	fmt.Println(arr)

	for _, v := range arr {
		fmt.Printf("%v %T %b\n", v, v, []byte(v))
	}
}
