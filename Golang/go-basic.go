package main

import (
	"fmt"
	"sort"
)

func RemainderSorting(strArr []string) []string {
	var (
		sum                              int
		sumArr1, sumArr2, sumArr0, summa []string
	)
	sort.Strings(strArr)
	for _, v := range strArr {
		sum = len(v) % 3
		if sum == 0 {
			sumArr0 = append(sumArr0, v)
		} else if sum == 1 {
			sumArr1 = append(sumArr1, v)
		} else {
			sumArr2 = append(sumArr2, v)
		}
	}
	summa = append(summa, sumArr0...)
	summa = append(summa, sumArr1...)
	summa = append(summa, sumArr2...)
	return summa
}

func main() {
	strArr := []string{"Salam", "Hello", "friend"}
	fmt.Println(RemainderSorting(strArr))
}
