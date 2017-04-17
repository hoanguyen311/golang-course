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

	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)
	buckets := make(map[int]map[string]int)

	for i := 0; i < 12; i++ {
		buckets[i] = make(map[string]int)
	}

	for scanner.Scan() {
		w := scanner.Text()
		n := HashBucket(w, 12)
		buckets[n][w]++
	}

	fmt.Println(buckets[6])
}

func HashBucket(word string, bucket int) int {
	//var sum int
	//
	//for _, v := range word {
	//	sum += int(v)
	//}
	//
	//return sum % bucket
	return len(word) % bucket
}
