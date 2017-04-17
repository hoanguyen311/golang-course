package main

import (
	"fmt"
	"sort"
)

func main() {
	p := []string{"Hoa", "Cuc", "An", "Binh"}

	sort.Strings(p)
	fmt.Println(p)
}
