package main

import "fmt"

func main() {
	buckets := make([]int, 1)

	buckets[0] = 40
	fmt.Println(buckets)

	buckets[0]++
	fmt.Println(buckets)


}
