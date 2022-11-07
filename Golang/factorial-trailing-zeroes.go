package main

import "fmt"

func trailingZeroes(n int) int {
	sum := 0
	i := 5
	for i <= n {
		sum += n / i
		i *= 5
	}
	return sum
}

func main() {
	fmt.Println(trailingZeroes(1674))
}
