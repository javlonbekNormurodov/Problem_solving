package main

import "fmt"

func findSubarrays(nums []int) bool {
	sumMap := make(map[int]bool)
	var sum int
	for i := 0; i < len(nums)-1; i++ {
		sum = nums[i] + nums[i+1]
		if _, ok := sumMap[sum]; ok {
			return ok
		}
		sumMap[sum] = true
	}
	return false
}

func main() {
	num := []int{4, 2, 3}
	fmt.Println(findSubarrays(num))
}
