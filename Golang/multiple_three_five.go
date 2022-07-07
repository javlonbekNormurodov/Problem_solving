package main

import (
	"fmt"
)

func biggerIsGreater(w string) {
	// Write your code here
	runeStr := []rune(w)
	var (
		result []int
		count  int
		str    string
	)
	for i := 0; i < len(runeStr); i++ {
		result = append(result, int(runeStr[i]))
	}
	for j := 0; j < len(result); j++ {
		if result[0] < result[j] {
			count += 1
		}
	}
	if count == 0 {
		fmt.Println("no answer")
	}
	for i := len(result) - 1; i >= 0; i-- {
		j := 1
		if result[i] > result[i-j] {
			temp := result[i-j]
			temp2 := result[i]
			result[i] = temp2
			result[i-j] = temp
			break
		}
	}
	for i := 0; i < len(result); i++ {
		var b byte
		b = byte(result[i])
		str += string(b)
	}
	fmt.Println(str)
}
