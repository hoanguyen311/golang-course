package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
)

func main() {
	res, err := http.Get("https://golang.org/pkg/net/http/")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
