package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Last string
	First string
	Age int
	isUnexported bool
}

func main() {
	p := Person{"Nguyen", "Thai Hoa", 26, false}

	bs, _ := json.Marshal(p)

	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
}
