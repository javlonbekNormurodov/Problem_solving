package main

import "math"

func reverse(i int32) int32 {
  var (
    res int32
  )
  for i > 0 {
    reminder := i % 10
    res = (res * 10) + reminder
    i /= 10
  }

  return res
}

func beautifulDays(i, j, k int32) int32 {
  var (
    countBeautifulDays int32
  )

  for initial := i; initial <= j; initial++ {
    var (
      diff int32
    )
    diff = int32(math.Abs(float64(initial) - float64(reverse(initial))))
    if diff%k == 0 {

      countBeautifulDays += 1
    }
  }

  return countBeautifulDays
}