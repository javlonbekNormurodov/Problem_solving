package main

// this is my code
// func isPalindrome(s string) bool {
// 	s = strings.ToLower(s)
// 	var r string
// 	for _, v := range s {
// 		if int(v) < 97 || int(v) > 122 {
// 			if int(v) < 48 || int(v) >= 58 {
// 				s = strings.Replace(s, string(v), "", len(s))
// 			}
// 		}
// 	}
// 	for i := len(s) - 1; i >= 0; i-- {
// 		r += string(s[i])
// 	}
// 	if string(r) == s {
// 		return true
// 	}
// 	return false
// }

//THIS IS THE BEST CASE
// func isPalindrome(s string) bool {
// 	l := 0
// 	r := len(s) - 1
// 	for l < r {
// 		for l < r && !unicode.IsLetter(rune(s[l])) && !unicode.IsNumber(rune(s[l])) {
// 			l++
// 		}
// 		for l < r && !unicode.IsLetter(rune(s[r])) && !unicode.IsNumber(rune(s[r])) {
// 			r--
// 		}
// 		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
// 			return false
// 		}
// 		l++
// 		r--
// 	}
// 	return true
// }

func main() {
	// str := "0P"
	// fmt.Println(isPalindrome(str))
}
