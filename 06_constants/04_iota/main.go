package main

import "fmt"

const (
	_ = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

func main() {
	hello := "Hello Hoa";
	//fmt.Printf("KB (binary): %b\n", KB)
	//fmt.Printf("MB (binary): %b\n", MB)
	//fmt.Printf("GB (binary): %b\n", GB)
	//fmt.Printf("TB (binary): %b\n", TB)
	fmt.Printf("KB (decimal): %d\n", KB)
	fmt.Printf("MB (decimal): %d\n", MB)
	fmt.Printf("GB (decimal): %d\n", GB)
	fmt.Printf("TB (decimal): %d\n", TB)

	fmt.Printf("%T: %v\n", "Hello Hoa", "Hello Hoa");
	fmt.Printf("%T: %v\n", hello, hello);
}
