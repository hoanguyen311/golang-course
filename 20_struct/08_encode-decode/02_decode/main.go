package main

import (
	"strings"
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

	rd := strings.NewReader(`{ "Last": "Nguyen", "First": "Hoa", "Age": 26 }`)

	json.NewDecoder(rd).Decode(&p)

	fmt.Println(p)
}
