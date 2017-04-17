package main

import (
	"fmt"
	"sort"
)

func main() {
	p := []string{"Hoa", "Cuc", "An", "Lam"}

	sort.Sort(sort.Reverse(sort.StringSlice(p)))

	fmt.Println(p)

	r := sort.Reverse(sort.StringSlice(p))

	fmt.Println(r)
	//fmt.Printf("%T \n", r)
}
