package main

import (
	"strconv"
	"fmt"
)

func main() {
	b, _ := strconv.ParseBool("true")

	fmt.Println(b)

	f, _ := strconv.ParseFloat("12.3434", 64)

	fmt.Println(f)

	i, _ := strconv.ParseInt("-10", 0, 64)

	fmt.Println(i)

	u, _ := strconv.ParseUint("10", 0, 64)

	fmt.Println(u)

	sB := strconv.FormatBool(true)
	sF := strconv.FormatFloat(12.3434, 'E', -1, 64)
	sI := strconv.FormatInt(12, 16)
	sU := strconv.FormatUint(10, 16)

	fmt.Println(sB, sF, sI, sU)
}
