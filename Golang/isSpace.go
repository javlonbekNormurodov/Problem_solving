package main

import (
	"regexp"
)

func isSpace(word string) bool {

	whitespace := regexp.MustCompile(`\s`).MatchString(word)
	return whitespace
}
