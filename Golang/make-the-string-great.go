package main

import (
	"fmt"
	"strings"
)

func makeGood(s string) string {
	for i := range s {
		// if i == len(s)-1 {
		// 	break
		// }
		if s[i] > 64 && s[i] < 91 {
			fmt.Println(s[i])
			s = string(s)
			a := s[i]
			s = strings.Replace(s, string(a), "", 1)
			a2 := s[i-1]
			s = strings.Replace(s, string(a2), "", 1)
		}
	}
	return s
}
func main() {
	s := "aAbBcC"
	fmt.Println(makeGood(s))

}
