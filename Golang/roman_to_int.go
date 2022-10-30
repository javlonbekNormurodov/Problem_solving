package main

func romanToInt(s string) int {
	roman := make(map[string]int)
	roman["I"] = 1
	roman["V"] = 5
	roman["X"] = 10
	roman["L"] = 50
	roman["C"] = 100
	roman["D"] = 500
	roman["M"] = 1000

	arr := []string{}
	result := []int{}

	for i := 0; i < len(s); i++ {
		arr = append(arr, s)
	}
	for _, val := range arr {
		if roman[val] < roman[val] {
			result = append(result, roman[val]-roman[val])
		}
	}
	var sum int

	for _, v := range result {
		sum += v
	}
	return sum
}
