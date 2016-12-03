package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {
	res, err := http.Get("http://lazada.vn/")

	if err != nil {
		log.Fatal(err)
	}

	page, err := ioutil.ReadAll(res.Body);

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s \n", page);
}
