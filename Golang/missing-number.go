package main

import (
	"fmt"
)

// func missingNumber(nums []int) int {
// 	var copiedNums []int
// 	sort.Ints(nums)
// 	copiedNums = append(copiedNums, 0)
// 	for i := range nums {
// 		copiedNums = append(copiedNums, copiedNums[i]+1)
// 	}

// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] != copiedNums[i] {
// 			return copiedNums[i]
// 		}
// 	}
// 	if len(nums) == 1 {
// 		return 1
// 	}
// 	if nums[len(nums)-1] != copiedNums[len(copiedNums)-1] && len(nums) != 1 {
// 		return copiedNums[len(copiedNums)-1]
// 	}

// 	return 0
// }
// This is the best case
func missingNumber(nums []int) int {
	n := len(nums)
	res := n
	for i := 0; i < len(nums); i++ {
		res ^= nums[i] ^ i
	}
	return res
}

func main() {
	nums := []int{1, 2}
	fmt.Println(missingNumber(nums))

}
