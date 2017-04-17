package main

import (
	"fmt"
	"sort"
)

func main() {
	is := []int{2, 1, 10, 9, 100, 20, 50}
	fmt.Println(is)
	sort.Sort(sort.IntSlice(is))
	fmt.Println(is)
}
