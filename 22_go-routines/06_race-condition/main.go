package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)

	go increament("Foo")
	go increament("Bar")

	wg.Wait()

	fmt.Println("Result: ", counter)
}

func increament(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Microsecond)
		counter++
		fmt.Println(s, ", i: ", i, ", counter: ", counter)
	}
	wg.Done()
}

