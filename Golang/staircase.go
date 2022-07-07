package main

import (
	"fmt"
	"strings"
)

func staircase(n int32) {
	var (
		j int = 1
	)

	for i := n; i > 0; i-- {
		fmt.Printf("%s%s\n", strings.Repeat(" ", int(i-1)), strings.Repeat("#", j))
		j += 1
	}
}
