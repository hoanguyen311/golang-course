package main

import (
	"log"
	"fmt"
)

type NorgateError struct {
	lat, lng string
	err error
}

func (n *NorgateError) Error() string {
	return fmt.Sprintf("a norgate math occured: %v, %v, %v", n.lat, n.lng, n.err)
}

func main() {
	_, err := sqrt(-10)

	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(n float64) (float64, error) {
	if n < 10 {
		nme := fmt.Errorf("norgate math redux: square root of negative number %v", n)
		return 0, &NorgateError{"100.00", "203.00", nme}
	}
	return 42, nil
}
