package main

import "fmt"

func duplicate(array []int, num int) bool {
	for _, v := range array {
		if v == num {
			return false
		}
	}
	return true
}

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	array := []int{}
	for i := range nums1 {
		if m[nums1[i]] == 1 || m[nums1[i]] == 0 {
			m[nums1[i]] = 1
		}
	}
	for _, v := range nums2 {
		if m[v] == 1 {
			if duplicate(array, v) {
				array = append(array, v)
			}
		}
	}
	return array
}

func main() {
	nums1 := []int{1, 3, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}
