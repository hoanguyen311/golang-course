package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")


	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(res.Body)

	sc.Split(bufio.ScanWords)
	buckets := make([]int, 200)
	for sc.Scan() {
		w := sc.Text()

		buckets[hashBucket(w)]++
	}

	fmt.Println(buckets[65:123])
}

func hashBucket(word string) int {
	return int(word[0])
}