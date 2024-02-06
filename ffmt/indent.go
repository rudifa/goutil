package ffmt

/*
Copyright © 2024 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/

import "fmt"

// IndentNestedBrackets indents nested brackets in a string
// where brackets is a string of length 2, e.g. "{}", "<>", "[]"
// and indent is the string to be used for indentation, e.g. "  ", "\t"
func IndentNestedBrackets(s string, brackets string, indent string) (string, error) {
	if len(brackets) != 2 {
		return "", fmt.Errorf("brackets must be a string of length 2")
	}

	indentstr := ""
	result := ""
	openingBracket := rune(brackets[0])
	closingBracket := rune(brackets[1])

	for _, c := range s {
		switch c {
		case openingBracket:
			indentstr += indent
			result += string(c) + "\n" + indentstr
		case closingBracket:
			indentstr = indentstr[:len(indentstr)-len(indent)]
			result += "\n" + indentstr + string(c)
		default:
			result += string(c)
		}
	}
	return result, nil
}
