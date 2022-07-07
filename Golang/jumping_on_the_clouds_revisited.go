package main

// import "fmt"

func jumpingOnClouds(c []int32, k int32) int32 {
	var (
		e = 100
	)
	for i := 0; i < len(c); i += int(k) {
		e -= 1
		if c[i] == 1 {
			// e = e - (int(k) + 1)
			e = e - 2
		}
	}
	return int32(e)
}
// func main() {
// 	c := []int32{1,1,1,0,0,1,1,0,0,0}
// 	k := int32(2)
// 	fmt.Println(jumpingOnClouds(c,k))
// }