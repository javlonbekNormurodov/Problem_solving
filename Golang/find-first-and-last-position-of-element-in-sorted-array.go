package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	var res, result []int
	for i := range nums {
		if nums[i] == target {
			res = append(res, i)
		}
	}
	if res == nil {
		res = append(res, -1, -1)
		return res
	}
	if len(res) == 1 {
		res = append(res, res[0])
	}
	if len(res) > 2 {
		result = append(result, res[0], res[len(res)-1])
		return result
	}
	return res
}

func main() {
	nums := []int{3, 3, 3}
	fmt.Println(searchRange(nums, 3))
}
