package main

import "fmt"

import "github.com/thaihoa311/hello/04_scope/myMath";




func main() {
	add4 := myMath.CreateAdd(10);

	fmt.Println(add4(5));
	fmt.Println(myMath.CreateAdd(2)(5));

	subtract := func (a int, b int) int {
		return a - b;
	}

	fmt.Println(subtract(10, 1))
	foo();
}

func foo() {
	fmt.Printf("FE Developer: %v\n", developer);
}

var developer string = "Hoa Nguyen";