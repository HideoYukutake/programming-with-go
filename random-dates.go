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
	year := 2100
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
	day := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, month, day, daysInMonth)
}
