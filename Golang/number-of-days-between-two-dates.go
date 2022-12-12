package main

import (
	"fmt"
	"strconv"
)

func daysBetweenDates(date1 string, date2 string) int {
	var (
		digits   int = 1000
		month    int
		year     int
		day      int
		month2   int
		year2    int
		day2     int
		resMonth int
		resYear  int
		resDay   int
	)
	for i := 0; i <= 3; i++ {
		year, _ = strconv.Atoi(string(date1[i]))
		year *= digits
		digits /= 10
	}
	digits = 10
	for i := 5; i <= 6; i++ {
		month, _ = strconv.Atoi(string(date1[i]))
		month *= digits
		digits /= 10
	}
	digits = 10
	for i := 7; i <= 8; i++ {
		day, _ = strconv.Atoi(string(date1[i]))
		day *= digits
		digits /= 10
	}
	for i := 0; i <= 3; i++ {
		year2, _ = strconv.Atoi(string(date2[i]))
		year2 *= digits
		digits /= 10
	}
	digits = 10
	for i := 5; i <= 6; i++ {
		month2, _ = strconv.Atoi(string(date2[i]))
		month2 *= digits
		digits /= 10
	}
	digits = 10
	for i := 7; i <= 8; i++ {
		day2, _ = strconv.Atoi(string(date2[i]))
		day2 *= digits
		digits /= 10
	}
	resYear = year - year2

	if resYear >= 0 {
		resMonth = (month + (resYear * 12)) - month2
	}
	if resMonth >= 0 {
		resDay = (day + (resMonth * 31)) - day2
	}
	return resDay
}

func main() {
	date1 := "2019-06-29"
	date2 := "2019-06-30"
	fmt.Println(daysBetweenDates(date1, date2))
}
