package main

func unique(s []int) []int {
	inResult := make(map[int]bool)
	var result []int
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
