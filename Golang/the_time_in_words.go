package main

func timeInWords(h int32, m int32) string {
	l := map[int32]string{
		0:  "o' clock",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twenlve",
		13: "thirteen",
		14: "fourteen",
		15: "quarter",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen"}

	decimal := map[int32]string{
		20: "twenty",
		40: "forty",
		30: "half",
		50: "fifty",
		60: "sixty",
	}
	if m > 30 {
		to := 60 - m
		if to < 20 {
			if to == 15 {
				return l[to] + " to " + l[h+1]
			}
			return l[to] + " " + "minutes to " + l[h+1]
		}
		decimalNum := to / 10
		decimalNum *= 10
		num := to % 10
		return decimal[decimalNum] + " " + l[num] + " minutes to " + l[h+1]
	} else if m == 0 {
		return l[h] + " " + l[m]
	} else {
		if m < 20 {
			if m == 15 {
				return l[m] + " past " + l[h]
			}
			if m == 1 {
				return decimal[m] + l[m] + " minute past " + l[h]
			}
			return l[m] + " " + " minutes past " + l[h]
		}
		if m == 30 {
			return decimal[m] + " past " + l[h]
		}
		decimalNum := m / 10
		decimalNum *= 10
		num := m % 10

		return decimal[decimalNum] + " " + l[num] + " minutes past " + l[h]
	}

}
