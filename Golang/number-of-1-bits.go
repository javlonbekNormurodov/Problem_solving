package main

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	fmt.Println(num)
	m := make(map[int]int)
	number := fmt.Sprintf("%s", num)
	fmt.Println(number)
	for _, v := range number {
		if v == '1' {
			fmt.Println("v", v)
			m[1]++
		}
	}
	return m[1]
}

func main() {
	var num uint32
	num = 000000000000000000010011
	fmt.Println(hammingWeight(num))
}
