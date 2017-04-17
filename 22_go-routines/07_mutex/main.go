package main

import (
	"sync"
	"fmt"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutext sync.Mutex

func main() {
	wg.Add(2)

	go increament("Foo: ")
	go increament("Bar: ")

	wg.Wait()

	fmt.Println(counter)
}

func increament(s string) {
	for i := 0; i < 20; i++ {

		time.Sleep(3 * time.Millisecond)
		mutext.Lock()
		x := counter
		x++
		counter = x
		mutext.Unlock()
		fmt.Println(s, counter)

	}
	wg.Done()
}
