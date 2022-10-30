package main

func digit(num []int) int {
	t := 1
	var sum int
	for _, val := range num {
		if t != 0 {
			sum += t * val
			t = t * 10
		}
	}

	return sum
}
