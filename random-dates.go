package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func isLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	year := rand.Intn(2100) + 1
	month := rand.Intn(12) + 1
	daysInMonth := 31
	switch month {
	case 2:
		daysInMonth = 28
		if isLeapYear(year) {
			daysInMonth = 29
		}
	case 4, 6, 9, 11:
		daysInMonth = 30
	}
	for count := 0; count < 10; count++ {
		day := rand.Intn(daysInMonth) + 1
		fmt.Println(era, year, month, day)
	}
}
