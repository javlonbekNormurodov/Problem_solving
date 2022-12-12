package main

import "fmt"

func lengthOfLastWord(s string) int {
	lenS := len(s)
	fmt.Println(lenS)

	lenLastWord := 0
	for i := lenS - 1; i >= 0; {
		for i >= 0 && string(s[i]) == " " {
			i--
		}
		if i < 0 {
			return 0
		}

		for i >= 0 && string(s[i]) != " " {
			//fmt.Println(i)
			//fmt.Println(string(s[i]))
			i--
			lenLastWord++
		}

		return lenLastWord
	}

	return 0
}

func main() {
	length := lengthOfLastWord("computerscience ")
	fmt.Println(length)
}
