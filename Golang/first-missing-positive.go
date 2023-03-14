package main

import (
	"fmt"
)

// func firstMissingPositive(nums []int) int {
// 	var res bool
// 	sort.Ints(nums)
// 	temp := 1

// 	for i := 0; i < len(nums); i++ {
// 		fmt.Println("temp before => ", temp)
// 		res = isHave(nums, temp)
// 		fmt.Println("res => ", res)
// 		if !res {
// 			return temp
// 		}
// 		temp++
// 	}
// 	return nums[len(nums)-1] + 1
// }

// func isHave(nums []int, s int) bool {
// 	for _, v := range nums {
// 		if s == v {
// 			return true
// 		}
// 	}
// 	return false
// }
func firstMissingPositive(nums []int) int {
	m := make(map[int]bool)
	var temp int = 1
	for _, v := range nums {
		m[v] = true
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[temp]; !ok {
			return temp
		}
		m[temp] = true
		temp += 1
	}
	return temp
}

func main() {
	nums := []int{1, 1000}
	fmt.Println(firstMissingPositive(nums))
}
