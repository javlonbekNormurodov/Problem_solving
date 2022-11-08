package main

import (
	"fmt"
	"strings"
)

func licenseKeyFormatting(s string, k int) string {
	var a string
	splited := strings.Split(s, "-")
	for i := range splited {
		if i > 1 {
			a = string(splited[1])
			a += splited[2]
			a += splited[3]
		}
	}
	spl := splited[0]
	spl += "-"
	a = strings.ToUpper(a)
	spl += a
	fmt.Println(spl)
	return spl
}

func main() {
	licence := "5F3Z-2e-9-w"
	k := 4
	licenseKeyFormatting(licence, k)
}
