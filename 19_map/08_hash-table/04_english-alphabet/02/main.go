package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
	"os"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")

	if err != nil {
		log.Fatal(err)
	}

	words := make(map[string]string)

	sc := bufio.NewScanner(res.Body)

	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
		fmt.Println(sc.Text())
	}

	if err := sc.Err(); err != nil {
		fmt.Println(os.Stderr, "reading input:", err)
	}
	i := 0
	for k := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++
	}
}
