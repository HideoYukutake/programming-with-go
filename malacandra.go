package main

import (
	"fmt"
)

func main() {
	const distance_to_malacandra = 56000000 // km
	var time_to_malacandra = 28 * 24        // 28日=28*24時間
	fmt.Printf("%v km/h\n", distance_to_malacandra/time_to_malacandra)
}
