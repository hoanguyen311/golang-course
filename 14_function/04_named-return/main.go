package main

import "fmt"

func main() {
	fmt.Println(getFullname("Nguyen", "Thai Hoa"))
}

func getFullname(lname, fname string) (s string) {
	s = fmt.Sprint(lname, " ", fname)
	return
}
