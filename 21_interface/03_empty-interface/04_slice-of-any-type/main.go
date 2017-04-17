package main

import "fmt"

type animal struct {
	Sound string
	Name  string
}

type dog struct {
	animal
	Friendly bool
}

type cat struct {
	animal
	Friendly bool
}

func main() {
	tom := cat{animal{Sound: "Meaw", Name: "Tom"}, true}
	tim := dog{}
	ted := dog{}
	jerry := cat{}

	myPets := []interface{}{tom, tim, ted, jerry}

	for k, v := range myPets {
		fmt.Println(k, "---", v)
	}

	myDog, ok := myPets[0].(cat)

	if ok {
		fmt.Println(myDog.Sound)
	}
}
