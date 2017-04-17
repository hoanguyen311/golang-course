package main

import "fmt"

type verhicles interface {}

type verhicle struct {
	Seats int
	MaxSpeed int
	Color string
}

type car struct {
	verhicle
	Wheels int
	Doors int
}

type boat struct {
	verhicle
	Length int
}

type plane struct {
	verhicle
	Jet bool
}

func main() {
	prius := car{}
	tacoma := car{}
	boeing747 := plane{
		verhicle: verhicle{ Seats: 4, MaxSpeed: 100, Color: "red" },
		Jet: true,
	}
	boeing757 := plane{}

	sanger := boat{}
	boutique := boat{
		Length: 100,
		verhicle: verhicle{ 4, 120, "blue" },
	}

	verhicles := []verhicles {prius, tacoma, boeing747, boeing757, sanger, boutique}

	for k, v := range verhicles {
		fmt.Println(k, " -- ", v)
	}
}