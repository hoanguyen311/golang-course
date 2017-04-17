package main

import "fmt"

type animal struct {
	Sound string
}

type dog struct {
	animal
	Friendly bool
}

type cat struct {
	animal
	Friendly bool
	Color string
}

func spec(a interface{}) {
	fmt.Println(a)
}

func main() {
	tom := cat{
		animal{ "Miao" },
		true,
		"black",
	}

	lucky := dog {
		animal: animal{ Sound: "Gau" },
		Friendly: true,
	}

	spec(tom)
	spec(lucky)
}
