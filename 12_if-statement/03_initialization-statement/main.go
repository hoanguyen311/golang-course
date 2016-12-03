package main

import "fmt"

func switchOnType(x interface{}) string {
	switch x.(type) {
	case int:
		return "int"
	case string:
		return "string"
	default:
		return "unknown"
	}
}

func main() {
	if theType := switchOnType("Hoa"); theType == "string" {
		fmt.Println("That is string")
	}
}
