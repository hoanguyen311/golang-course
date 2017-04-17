package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Last string
	First string
	Age int
}

func main() {
	var p Person

	fmt.Println(p.Last, p.First, p.Age)

	bs := []byte(`{ "Last" : "Nguyen", "First" : "Thai Hoa", "Age": 20 }`)

	json.Unmarshal(bs, &p)

	fmt.Println("-----------")

	fmt.Printf("%v %v %v\n", p.Last, p.First, p.Age)
	fmt.Printf("%T\n", p)
}