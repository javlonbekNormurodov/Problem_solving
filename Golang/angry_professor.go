package problemsolving

func angryProfessor(k int32, a []int32) string {
	// Write your code here
	var minus int32
	for _, v := range a {
		if v <= 0 {
			minus += 1
		}
	}
	if minus < k {
		return "YES"
	} else {
		return "NO"
	}
}