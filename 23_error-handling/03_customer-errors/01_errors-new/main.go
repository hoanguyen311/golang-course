package main

import (
	"log"
	"errors"
	"math"
	"fmt"
)

func main() {
	n, err := sqrt(-10)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(n)
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("You can not pass a negative number")
	}
	return math.Sqrt(n), nil
}
