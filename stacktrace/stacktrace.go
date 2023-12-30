/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/

package stacktrace

import (
	"regexp"
	"runtime"
	"strings"
)

// Stacktrace represents a stacktrace and transform methods
type Stacktrace struct {
	RawLines []string
}

// CapturedStacktrace returns a new instance of Stacktrace
// initialized from runtime.Stack() at the point where it was called
func CapturedStacktrace() *Stacktrace {
	buf := make([]byte, 4096)
	runtime.Stack(buf, false)
	return NewStacktraceFrom(string(buf))
}

// NewStacktrace returns a new instance of Stacktrace
// initialized from a string logged using runtime.Stack()
func NewStacktraceFrom(rawString string) *Stacktrace {
	rawLines := strings.Split(rawString, "\n")
	ReverseSlice(rawLines)
	return &Stacktrace{rawLines}
}

// OnelineString returns a one-line string representation of the stacktrace
func (st Stacktrace) OnelineString() string {
	return OnelineString(st.RawLines)
}

// OnelineString returns a one-line string representation of the stacktrace
func OnelineString(rawLines []string) string {
	funcSigs := []string{}
	for _, line := range rawLines {
		funcSig := ExtractFuncSignature(line)
		if funcSig != "" {
			funcSigs = append(funcSigs, funcSig)
		}
	}
	return strings.Join(funcSigs, " => ")
}

// ExtractFuncSignature extracts the function signature from a stack line
// that was generated by runtime.Stack()
// `cuelang.org/go/cue/format.(*printer).writeString(0xc00031b2c0, {0x18ce528, 0x1}, 0x0)` =>
// `format.(*printer).writeString`
func ExtractFuncSignature(line string) string {
	r := regexp.MustCompile(`((\w+)\.((\(\*\w+\))\.)?(\w+))\((.*?)\)`)
	matches := r.FindStringSubmatch(line)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// ReverseSlice reverses a slice of strings
func ReverseSlice(s []string) {
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
}

// AreSlicesEqual returns true if the slices are equal
func AreSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
