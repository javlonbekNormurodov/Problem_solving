package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// "bufio"
// "fmt"
// "io"
// "os"
// "strconv"
// "strings"

func getNumDraws(year int32) int32 {
	years := strconv.Itoa(int(year))
	url, err := http.Get("http://jsonmock.hackerrank.com/api/football_matches?year=" + years + "&page=1")
	if err != nil {
		fmt.Println("Error ==> ", err)
		return 0
	}
	res := &url.Body
	fmt.Println("res ==> ", res)
	fmt.Println("url ==> ", url.Header)
	return 0
}

func main() {
	year := int32(2011)
	getNumDraws(year)
}
