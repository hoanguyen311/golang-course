package main

import (
	"fmt"
	"time"
)

var publisherID int
var workerID int

func main() {
	input := make(chan string)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)

	go publisher(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)

	time.Sleep(1 * time.Millisecond)
}

func workerProcess(in chan string) {
	workerID++
	thisID := workerID

	for {
		fmt.Printf("%d waiting for input\n", thisID)
		input := <-in
		fmt.Printf("%d: input is \"%s\"\n", thisID, input)
	}
}

func publisher(out chan string) {
	publisherID++
	thisID := publisherID
	dataID := 0

	for {
		dataID++
		fmt.Printf("publisher ID %d is pushing data\n", thisID)
		out <- fmt.Sprintf("Data from publisher %d. Data %d", thisID, dataID)
	}
}
