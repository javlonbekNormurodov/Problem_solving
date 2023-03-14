package main

import "fmt"

func removeElement(nums []int, val int) int {
	var sum int
	for _, v := range nums {
		if v != val {
			sum += 1
		}
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 2, 2}
	n := 2
	fmt.Println(removeElement(nums, n))
}
