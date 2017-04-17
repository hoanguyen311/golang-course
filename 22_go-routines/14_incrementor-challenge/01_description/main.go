package main

import (
	"sync"
	"fmt"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int32

func main() {
	wg.Add(2)

	go incrementor("Hoa")
	go incrementor("Cuc")

	wg.Wait()
	fmt.Println("Final: ", counter)
}


func incrementor(name string) {
	for i := 0; i < 20; i++ {
		atomic.AddInt32(&counter, 1)
		fmt.Println(name, "\t", i)

	}
	wg.Done()
}