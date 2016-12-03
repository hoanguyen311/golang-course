package main

import "fmt"

func main() {
	fmt.Println(getFullname("Nguyen", "Thai Hoa"))
}

func getFullname(lname, fname string) string {
	return fmt.Sprint(lname, " ", fname)
}
