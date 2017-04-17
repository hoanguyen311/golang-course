package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txts")

	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	str := string(bs)

	fmt.Println(str)
}
