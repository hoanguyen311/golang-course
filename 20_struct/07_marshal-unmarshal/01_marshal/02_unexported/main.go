package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	last string
	first string
	age int
}

func main() {
	p := Person{"Nguyen", "Thai Hoa", 26}
	fmt.Println(p)

	bs, _ := json.Marshal(p)

	fmt.Println(bs)
	fmt.Println(string(bs))
}