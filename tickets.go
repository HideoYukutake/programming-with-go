package main

import (
	"fmt"
	"math/rand"
)

func main() {
	SpaceAdventure := "Space Adventure"
	SpaceX := "SpaceX"
	VirginGalactic := "Virgin Galactic"

	distance := 62100000 // km

	fmt.Printf("%-16v %-5v %-11v %-4v\n", "Spaceline", "Days", "Trip type", "Price")
	for i := 0; i < 40; i++ {
		fmt.Print("=")
	}
	fmt.Print("\n")
	for count := 0; count < 10; count++ {
		var company = SpaceAdventure
		switch idx := rand.Intn(3); idx {
		case 0:
			company = SpaceAdventure
		case 1:
			company = SpaceX
		case 2:
			company = VirginGalactic
		}
		speed := rand.Intn(15) + 16 // km/s
		days := distance / (speed * 60 * 60 * 24)
		cost := speed + 30 // million
		trip_type := "One-way"
		is_one_way := rand.Intn(2)
		if is_one_way == 0 {
			trip_type = "Round-trip"
			cost = cost * 2
		}
		fmt.Printf("%-16v %-5v %-11v $%4v\n", company, days, trip_type, cost)
	}
}
