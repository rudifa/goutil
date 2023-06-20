package fstrings

import (
	"unicode/utf8"
)

// Truncate string to specified length
func Truncate(str string, length int) string {

	if utf8.RuneCountInString(str) < length {
		return str + "\n"
	}

	return string([]rune(str)[:length]) + "..."
}

// HeadAndTail returns the first n runes and the last n runes of the string
func HeadAndTail(str string, n int) string {
	// Get the length of the string in runes
	strLen := utf8.RuneCountInString(str)

	// If the string length is less than or equal to 2n, return the entire string
	if strLen <= 2*n {
		return str
	}

	// Get the first n runes of the string
	headRunes := []rune(str)[:n]

	// Get the last n runes of the string
	tailRunes := []rune(str)[strLen-n:]

	// Create the resulting string with head runes, ellipsis, and tail runes
	result := string(headRunes) + "···" + string(tailRunes)

	return result
}
