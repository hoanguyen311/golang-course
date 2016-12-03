package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	res, _ := http.Get("http://lazada.vn/")

	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Printf("%", page);
}
