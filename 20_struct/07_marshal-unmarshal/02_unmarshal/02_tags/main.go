package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last string
	Age int `json:"wisdom score"`
}

func main() {
	var p Person

	bs := []byte(`{ "First": "Thai Hoa", "Last": "Nguyen", "wisdom score": 26 }`)

	json.Unmarshal(bs, &p)


	fmt.Println(p.Age)
}
