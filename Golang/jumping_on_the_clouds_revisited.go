package main

import "fmt"

func jumpingOnClouds(c []int32, k int32) int32 {
	var (
		e = 100
	)
	for i := 0; i < len(c); i += int(k) {
		if c[i] != 1 {
			e -= 1
		} else {
			// e = e - (int(k) + 1)
			e = e - 3
		}
		if i+int(k) > len(c) {
			fmt.Println(i)
			break
		}
	}
	return int32(e)
}
