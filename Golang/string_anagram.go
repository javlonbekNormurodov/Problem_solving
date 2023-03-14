package main

import (
	"fmt"
)

func alphasort(str []rune, depth int) {
	for x := range str {
		y := x + 1
		for y = range str {
			if str[x] < str[y] {
				str[x], str[y] = str[y], str[x]
			}
		}
	}
}

func stringAnagram(dictionary []string, query []string) []int32 {
	var (
		arr, arr1 []byte
		sum, sum1 byte
		resArray  []int32
	)
	m := make(map[byte]int32)
	for _, v := range dictionary {
		for _, j := range v {
			sum += byte(j)
		}
		arr = append(arr, sum)
	}

	for _, v := range query {
		for _, j := range v {
			sum1 += byte(j)
		}
		arr1 = append(arr1, sum1)
	}
	fmt.Println("arr => ", arr)
	fmt.Println("arr1 => ", arr1)
	for _, i := range arr {
		fmt.Println("i +> ", i)
		for _, j := range arr1 {
			fmt.Println("j => ", j)
			if i == j {
				m[i]++
			}
			resArray = append(resArray, m[j])
		}
	}
	return resArray
}

func main() {
	dictionary := []string{"hack", "a", "rank", "khac", "achk", "kran", "rankhacker", "a", "ab", "ba", "stairs", "raits"}
	query := []string{"a", "nark", "bs", "hack", "stair"}
	fmt.Println(stringAnagram(dictionary, query))
}
