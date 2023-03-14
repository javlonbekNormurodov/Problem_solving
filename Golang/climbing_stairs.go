package main

import (
	"fmt"
	"math"
)

func climbStairs(n int) int {
	result := math.Pow(float64(n), 0.5)
	result += float64(n) - result
	return int(result)
}

func main() {
	fmt.Println(climbStairs(5))
}
