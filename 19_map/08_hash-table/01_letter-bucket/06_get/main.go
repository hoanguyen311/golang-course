package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")

	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", bs)
}
