package main

import "fmt"

func maxProfit(prices []int) int {
	var minimum, max int
	minimum = prices[0]
	for i := 0; i < len(prices); i++ {
		if minimum > prices[i] {
			minimum = prices[i]
		}
		if prices[i]-minimum > max {
			max = prices[i] - minimum
		}
	}

	return max
}

func main() {
	x := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(x))
}
