package main

import (
	"sync"
	"sync/atomic"
	"time"
	"math/rand"
	"fmt"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go increment("Foo")
	go increment("Bar")
	wg.Wait()

	fmt.Println(counter)
}

func increment(s string) {
	for i:= 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, counter)
	}
	wg.Done()
}
