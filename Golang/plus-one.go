package main

func plusOne(digits []int) []int {
	lastNumber := len(digits) - 1
	lastElement := digits[lastNumber]
	lastElement += 1
	if lastElement == 10 {
		if len(digits) > 1 && digits[lastNumber] == 9 {

		}
		digits[lastNumber] = 1
		digits = append(digits, 0)
		return digits
	}
	digits[len(digits)-1] = lastElement
	return digits
}
