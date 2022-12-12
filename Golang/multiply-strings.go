package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	converted_num1, _ := strconv.ParseInt(num1, 10, 64)
	converted_num2, _ := strconv.ParseInt(num2, 10, 64)
	res := converted_num1 * converted_num2

	converted_string := strconv.FormatInt(res, 10)

	return converted_string
}

func main() {
	fmt.Println(multiply("3", "4"))
}
