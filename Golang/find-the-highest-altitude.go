package main

import (
	"fmt"
)

func largestAltitude(gain []int) int {
	var (
		highest []int
		sum     int = 0
	)
	for _, v := range gain {
		sum += v
		highest = append(highest, sum)
	}
	sum = 0
	for _, v := range highest {
		if sum < v {
			sum = v
		}
	}
	return sum
}

func main() {
	gain := []int{-5, 1, 5, 0, -7}
	fmt.Println(largestAltitude(gain))
}
