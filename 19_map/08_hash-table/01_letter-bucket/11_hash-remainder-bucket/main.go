package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
)

func main() {
	res, err := http.Get("http://websitetips.com/articles/copy/lorem/ipsum.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)

	scanner.Split(bufio.ScanWords)
	buckets := make([]int, 12)

	for scanner.Scan() {
		n := HashBucket(scanner.Text(), 12)
		buckets[n]++
	}

	fmt.Println(buckets)
}

func HashBucket(word string, bucket int) int {
	return int(word[0]) % bucket
}
