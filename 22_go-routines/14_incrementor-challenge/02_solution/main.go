package main

import (
	"fmt"
	"sync"
)

func main() {

	c := incrementor(4)
	var count int

	for n := range c {
		count++
		fmt.Println(n)
	}

	fmt.Println(count)
}
func incrementor(n int) chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(n)

	for k := 0; k < n; k++ {
		go func(procId int) {
			for i := 0; i < 20; i++ {
				out <- fmt.Sprintf("Process %d is printing %d", procId, i)
			}
			wg.Done()
		}(k + 1)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}