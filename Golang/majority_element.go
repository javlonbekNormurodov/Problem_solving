package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	m := make(map[int]int)
	l := make(map[int]int)
	// sort.Ints(nums)
	for _, v := range nums {
		m[v] += 1
	}
	fmt.Println("m => ", m)
	max := 0
	for _, i := range nums {
		fmt.Println("max => ", max)
		if m[i] > max {
			max = m[i]
			l[max] = i
		}
	}
	return l[max]
}

func main() {
	nums := []int{-1, 1, 1, 1, 2, 1}
	fmt.Println(majorityElement(nums))
}
