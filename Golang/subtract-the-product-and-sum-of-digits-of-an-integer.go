package main

import "fmt"

func digits(n int) (int, int) {
	var sum, sumOfNums int
	multiplication := 1
	for n != 0 {
		sum = n % 10
		n /= 10
		sumOfNums += sum
		multiplication *= sum

	}
	return sumOfNums, multiplication
}

func subtractProductAndSum(n int) int {
	a, b := digits(n)
	fmt.Printf("a = %d and b = %d\n", a, b)
	result := b - a
	return result
}
