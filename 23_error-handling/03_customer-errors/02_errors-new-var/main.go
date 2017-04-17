package main

import (
	"log"
	"math"
	"errors"
	"fmt"
)

var ErrNorgate = errors.New("norgate math: can not be negative")

func main() {
	//_, err := sqrt(-10)
	n, err := sqrt(10)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(n)
}

func sqrt(n float64) (float64, error) {
	if n < 10 {
		return 0, ErrNorgate
	}
	return math.Sqrt(n), nil
}