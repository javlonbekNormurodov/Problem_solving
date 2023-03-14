package main

import "fmt"

func linearSearch(arr []int, m int) bool {
	for i := 0; i < len(arr); i++ {
		if m == arr[i] {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 4, 1, 3, 2, 3, 4, 6}
	m := 6
	fmt.Println(linearSearch(arr, m))
}
