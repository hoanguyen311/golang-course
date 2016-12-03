package main

import "fmt"

func is_odd(i int) bool {
	return i % 2 == 0
}

func main() {
	i := 0

	for {
		i++

		if i > 50 {
			break
		}

		if is_odd(i) {
			continue
		}

		fmt.Println(i)
	}
}
