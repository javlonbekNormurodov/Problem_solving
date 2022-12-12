package main

import (
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n > 0 {
		for i := 0; i < n; i++ {
			nums1[i+m] = nums2[i]
		}
		sort.Ints(nums1)
	}
}

func main() {
	nums1 := []int{0}
	m := 1
	nums2 := []int{1}
	n := 0
	merge(nums1, m, nums2, n)
}
