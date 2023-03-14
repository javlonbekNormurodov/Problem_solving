package main

import "fmt"

func bubbleSort(arr []int) []int {
	var isDone = false

	for !isDone {
		isDone = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isDone = false
			}
		}
	}
	return arr
}

func main() {
	arr := []int{1, -3, 3, 4, 2, 1, 34, 4, 2, 3, 1}
	fmt.Println(bubbleSort(arr))
}
