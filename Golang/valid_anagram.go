package main

import "fmt"

func isAnagram(s string, t string) bool {
	mapt, maps := make(map[rune]int), make(map[rune]int)
	for _, v := range s {
		maps[v]++
	}
	for _, v := range t {
		mapt[v]++
	}
	if len(mapt) > len(maps) {
		for i := range mapt {
			if mapt[i] != maps[i] {
				return false
			}
		}
	} else {
		for i := range maps {
			if maps[i] != mapt[i] {
				return false
			}
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}
