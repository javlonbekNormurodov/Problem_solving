package main

import (
	"fmt"
	"sort"
)

func removeDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func longestConsecutive(nums []int) int {
	nums = removeDuplicateValues(nums)
	var sum, max int
	var nms []int
	sort.Ints(nums)
	fmt.Println("nums => ", nums)
	if len(nums) == 1 {
		return 1
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1]-1 {
			sum++
		} else {
			nms = append(nms, sum)
			sum = 0
		}
		nms = append(nms, sum)
	}
	sort.Ints(nms)
	if len(nms) > 0 {
		max = nms[len(nms)-1]
		return max + 1
	}
	return 0
}

func main() {
	num := []int{1, 2, 0, 1}
	fmt.Println(longestConsecutive(num))
}
