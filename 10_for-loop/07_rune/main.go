package main

import "fmt"

func main() {
	for i := 33; i < 200; i++ {
		fmt.Println('i', " - ", string(i), []byte(string(i)))
	}

	name := "Hoa";

	fmt.Println(`
		<h1 class="title">
		` + name +  `
		</h1>
	`);
}
