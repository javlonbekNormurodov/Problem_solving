package problemsolving

func matchingStrings(strings []string, queries []string) []int32 {
	// Write your code here
	var (
		sum   []int32
		count int32
	)
	for _, a := range queries {
		count = 0
		for b := range strings {
			if a == strings[b] {
				count += 1
			}
		}
		sum = append(sum, count)
	}
	return sum
}