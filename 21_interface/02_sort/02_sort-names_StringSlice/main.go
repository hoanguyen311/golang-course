package main

import (
	"fmt"
	"sort"
)

func main() {
	p := []string{"A", "C", "B", "Z", "E", "H"}

	sort.Sort(sort.StringSlice(p))
	sort.StringSlice(p).Sort()

	fmt.Println(p)
}
