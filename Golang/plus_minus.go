package problemsolving

import "fmt"

func plusMinus(arr []int32) {
	// Write your code here
	var (
		sumP, sumN, sumZ, elements float64
	)
	for a := range arr {
		if arr[a] < 0 {
			sumN += 1
		} else if arr[a] == 0 {
			sumZ += 1
		} else {
			sumP += 1
		}
		elements += 1
	}
	fmt.Printf("%.6f\n", sumP/elements)
	fmt.Printf("%.6f\n", sumN/elements)
	fmt.Printf("%.6f\n", sumZ/elements)
}
