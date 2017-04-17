package main

import (
	"encoding/json"
	"os"
)

type Person struct{
	First string
	Last string
	Age int
}

func main() {
	p := Person{
		First: "Thai Hoa",
		Last: "Nguyen",
		Age: 26,
	}

	json.NewEncoder(os.Stdin).Encode(p)
}
