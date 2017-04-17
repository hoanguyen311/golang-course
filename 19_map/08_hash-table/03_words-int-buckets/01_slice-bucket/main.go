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
	buckets := make([][]string, 12)

	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}

	for scanner.Scan() {
		w := scanner.Text()
		n := HashBucket(w, 12)
		buckets[n] = append(buckets[n], w)
	}

	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}

	fmt.Println(buckets[0])
}

func HashBucket(word string, bucket int) int {
	var sum int

	for _, v := range word {
		sum += int(v)
	}

	return sum % bucket
}
