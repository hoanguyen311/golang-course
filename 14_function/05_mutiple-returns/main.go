package main

import "fmt"

func main() {
	fmt.Println(getFullname("Hoa", "Nguyen"))
}

func getFullname(fname, lname string) (string, string) {
	return fmt.Sprint(fname, " ", lname), fmt.Sprint(lname, " ", fname)
}