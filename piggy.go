package main

import (
	"fmt"
	"math/rand"
)

func main() {
	piggyBank := 0.0
	distance := 20.0 // $

	for piggyBank < distance {
		switch idx := rand.Intn(3); idx {
		case 0:
			piggyBank += 0.05
		case 1:
			piggyBank += 0.1
		case 2:
			piggyBank += 0.25
		}
	}
	fmt.Printf("%05.2f\n", piggyBank)
}
