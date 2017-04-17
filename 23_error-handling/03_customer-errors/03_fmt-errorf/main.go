package main

import (
	"log"
	"fmt"
	"math"
)

func main() {
	_, err := sqrt(-10)

	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("norgate math: can not be negative number: %v", n)
	}
	return math.Sqrt(n), nil
}