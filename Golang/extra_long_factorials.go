package main

import (
	"fmt"
	"math/big"
)

func extraLongFactorials(n int32) {
	// Write your code here
	b := big.NewInt(0)
	b.MulRange(1, int64(n))
	fmt.Println(b)
}
