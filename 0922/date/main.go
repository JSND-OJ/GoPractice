package main

import (
	"GoPractice/myPKG/calender"
	"fmt"
)

func main() {
	date := calender.Date{}
	date.Year = 2000
	date.Month = 14
	date.Day = 40
	fmt.Println(date)

	date = calender.Date{Year: 0, Month: 0, Day: -2}
	fmt.Println(date)
}
