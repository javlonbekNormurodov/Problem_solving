package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	str := strings.Replace(address, ".", "[.]", 3)
	fmt.Println(str)
	return address
}

func main() {
	ip := "1.1.1.1"
	defangIPaddr(ip)
}
