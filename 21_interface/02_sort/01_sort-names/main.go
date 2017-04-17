package main

import (
	"fmt"
	"sort"
)

type People []string

func (p People) Len() int           { return len(p) }
func (p People) Less(i, j int) bool { return p[i] < p[j] }
func (p People) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p People) Sort() {
	sort.Sort(p)
}

type Positions []int

func (p Positions) Len() int           { return len(p) }
func (p Positions) Less(i, j int) bool { return p[i] < p[j] }
func (p Positions) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	p := []string{"Hoa", "B", "C", "A", "Cuc"}

	//sort.Sort(sort.Reverse(People(p)))

	People(p).Sort()

	fmt.Println(p)

	nums := []int{2, 3, 120, 10, 100}

	sort.Sort(sort.Reverse(Positions(nums)))

	fmt.Println(nums)
}
