package main

import (
	"fmt"
	"time"
)

const theName string = "Nguyen Thai Hoa"
const yearOfBirth int = 1990

func main() {
	now := time.Now()


	fmt.Printf("Name: %s \n", theName)
	fmt.Printf("Age: %v \n", now.Year() - yearOfBirth)
}
