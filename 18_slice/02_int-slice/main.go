package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5)

	fmt.Println(mySlice)
	fmt.Println(cap(mySlice))
	fmt.Println(len(mySlice))


	for i := 0; i < 100; i++ {
		mySlice = append(mySlice, i);

		fmt.Printf("Slice at %v: %v, Cap: %v, Length: %v\n", i, mySlice[i], cap(mySlice), len(mySlice))
	}

	secondSlice := []int{2, 3, 4}
}
