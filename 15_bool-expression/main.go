package main

import "fmt"

func main() {
	hasIcecream := true
	hasCake := false

	if hasCake {
		fmt.Println("A is happy")
	} else {
		fmt.Println("A is not happy")
	}

	if hasIcecream {
		fmt.Println("B is happy")
	} else {
		fmt.Println("B is not happy")
	}

	if hasCake || hasIcecream {
		fmt.Println("C is happy")
	} else {
		fmt.Println("C is not happy")
	}

	if hasCake && hasIcecream {
		fmt.Println("D is happy")
	} else {
		fmt.Println("D is not happy")
	}
}
