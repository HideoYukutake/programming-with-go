package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var answer = 17
	var count = 0
	rand.Seed(time.Now().UnixNano())
	for {
		count++
		var choice = rand.Intn(100)
		if answer == choice {
			fmt.Printf("You got it at %vth challenge!\n", count)
			break
		} else {
			fmt.Printf("%v is Wrong...\n", choice)
		}

	}
}
