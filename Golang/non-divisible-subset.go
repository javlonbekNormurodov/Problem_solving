package main

func nonDivisibleSubset(k int32, s []int32) int32 {
	// Write your code here
	var (
		// firstCount int32
		// resultCount int32
		count int32
	)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if (s[i]+s[j])%k != 0 {
				count += 1
			}
		}
	}
	return count

}
