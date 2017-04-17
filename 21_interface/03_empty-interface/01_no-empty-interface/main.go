package main

import "fmt"

type verhicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	verhicle
	Wheels int
	Doors  int
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

	cars := []car{prius, tacoma}

	for k, v := range cars {
		fmt.Println(k, " -- ", v)
	}

	boeing747 := plane{
		verhicle: verhicle{Seats: 4, MaxSpeed: 100, Color: "red"},
		Jet:      true,
	}
	boeing757 := plane{}

	planes := []plane{boeing747, boeing757}

	for k, v := range planes {
		fmt.Println(k, " -- ", v)
	}

	sanger := boat{}
	boutique := boat{
		Length:   100,
		verhicle: verhicle{4, 120, "blue"},
	}

	boats := []boat{sanger, boutique}

	for k, v := range boats {
		fmt.Println(k, " -- ", v)
	}
}
