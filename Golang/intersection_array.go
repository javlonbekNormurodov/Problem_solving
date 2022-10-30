package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	var (
		ans   []int
		small int
	)

	sort.Ints(nums1)
	sort.Ints(nums2)
	if len(nums1) < len(nums2) {
		small = len(nums1)
	} else if len(nums1) > len(nums2) {
		small = len(nums2)
	} else {
		small = len(nums1)
	}
	for _, v := range nums1 {
		for _, val := range nums2 {
			if v == val {
				ans = append(ans, v)
			}
		}
	}
	fmt.Println(small)

	return ans
}
