package main

import (
	"fmt"
	"sort"
)

func main() {
	is := []int{2, 10, 8, 4, 5, 3, 100, 99}
	sort.Ints(is)

	fmt.Println(is)
}
