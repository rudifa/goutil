/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/

package stacktrace

import (
	"regexp"
	"runtime"
	"strings"

	"github.com/rudifa/goutil/util"
)

// TODO describe usage with examples
// TODO separate tests into wb_test and bb_test

// stacktrace api ----------------------------------------------------------------------------------

// Stacktrace represents a stacktrace and has transform methods
type Stacktrace struct {
	// each pair of raw lines from runtime.Stack() describing a call is joined into a single line
	PairedRawLines []string
}

// CapturedStacktrace returns a new instance of Stacktrace
// initialized from runtime.Stack() at the point where it was called
func CapturedStacktrace() *Stacktrace {
	buf := make([]byte, 4096)
	runtime.Stack(buf, false)
	return NewStacktraceFrom(string(buf))
}

// NewStacktraceFrom returns a new instance of Stacktrace
// initialized from a string obtained from runtime.Stack()
func NewStacktraceFrom(rawString string) *Stacktrace {
	pairedLines := strings.Split(JoinLinePairs(rawString), "\n")
	return &Stacktrace{pairedLines}
}

// StacklineCallersigs returns a one-line string representation of the stacktrace
// with the function signatures (path (receiver) funcname) for each call
// captured at the point of creation of the Stacktrace instance, e.g.
// "main.main => main.runParseAndFormat_2274 => main.runParseAndFormat => format.Node => format.(*config).fprint =>..."
func (st Stacktrace) StacklineCallersigs() string {
	pairedRawLines := st.PairedRawLines
	return StacklineCallersigs(pairedRawLines)
}

// StacklineCallersigs returns an one-line string representation of the stacktrace, e.g.
// "main.main => main.runParseAndFormat_2567 => main.runParseAndFormat => format.(*printer).writeString => format.(*printer).append => format.LogStackTrace"
func StacklineCallersigs(rawLines []string) string {
	funcSigs := []string{}
	for _, line := range rawLines {
		funcSig := ExtractFuncSignature(line)
		if funcSig != "" {
			funcSigs = append(funcSigs, funcSig)
		}
	}
	util.ReverseSlice(funcSigs)
	return strings.Join(funcSigs, " => ")
}

// StacklineCallpoints returns an one-line string representation of the stacktrace
// with the calling function name and file:line for each call
// captured at the point of creation of the Stacktrace instance, e.g.
// "main main.go:26 => runParseAndFormat_2567 main.go:37 => runParseAndFormat main.go:104 => writeString printer.go:365 => append printer.go:454 => LogStackTrace printer.go:506"
func (st Stacktrace) StacklineCallpoints() string {
	return StacklineCallpoints(st.PairedRawLines)
}

// StacklineCallpoints returns an one-line string representation of the stacktrace
// with the calling function name and file:line for each call
func StacklineCallpoints(pairedLines []string) string {
	funcCallSigs := []string{}
	for _, line := range pairedLines {
		funcSig := ExtractCallpoint(line)
		if funcSig != "" {
			funcCallSigs = append(funcCallSigs, funcSig)
		}
	}
	util.ReverseSlice(funcCallSigs)
	return strings.Join(funcCallSigs, " => ")
}

// stacktrace utilities ----------------------------------------------------------------------------

// ExtractFuncSignature extracts the function signature from a stack line
// that was generated by runtime.Stack(), e.g.
// "format.(*printer).writeString"
func ExtractFuncSignature(line string) string {
	r := regexp.MustCompile(`((\w+)\.((\(\*\w+\))\.)?(\w+))\((.*?)\)`)
	matches := r.FindStringSubmatch(line)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// ExtractCallsite returns a `Callsite` from a stack line, e.g. "printer.go:506"
func ExtractCallsite(line string) string {
	r := regexp.MustCompile(`(\w+\.go:\d+)`)
	matches := r.FindStringSubmatch(line)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// ExtractCallpoint returns a `Callpoint` e.g. "LogStackTrace printer.go:506"
func ExtractCallpoint(line string) (callpoint string) {
	// TODO combine the regex into one
	re1 := regexp.MustCompile(`((\w+)\.((\(\*\w+\))\.)?(\w+))\((.*?)\)`)
	re2 := regexp.MustCompile(`.+(\b\w+\.go:\d+)`)

	matches1 := re1.FindStringSubmatch(line)
	// log.Printf("len(matches)=%d\n", len(matches1))
	// for i, match := range matches1 {
	// 	log.Printf("matches[%d]=%s\n", i, match)
	// }
	matches2 := re2.FindStringSubmatch(line)
	// log.Printf("len(matches)=%d\n", len(matches2))
	// for i, match := range matches2 {
	// 	log.Printf("matches[%d]=%s\n", i, match)
	// }

	if len(matches1) > 5 && len(matches2) > 1 {
		return matches1[5] + ` ` + matches2[1]
	}
	return ""
}

// JoinLinePairs returns a string with each pair of lines describing a call joined into a single line, e.g.
// "cuelang.org/go/cue/format.LogStackTrace() /Users/rudifarkas/GitHub/golang/src/cue/cue/format/printer.go:506 +0x3a"
func JoinLinePairs(rawString string) string {
	const (
		pattern1 = `(\w+.\w+.+)`
		pattern2 = `\n\s+`
		pattern3 = `(.+)`
		pattern  = pattern1 + pattern2 + pattern3
	)
	// re should match pairs of lines in rawString from runtime.Stack(buf, false)
	re := regexp.MustCompile(pattern)

	// log.Printf("rawString=\n%s\n", rawString)

	// in every occurrence of re in rawString, replace pattern2 by ` `
	result := re.ReplaceAllStringFunc(rawString, func(s string) string {
		matches := re.FindStringSubmatch(s)
		// log.Printf("len(matches)=%d\n", len(matches))

		// for i, match := range matches {
		// 	log.Printf("matches[%d]=%s\n", i, match)
		// }
		if len(matches) > 0 {
			// log.Printf("matches[0]=%s\n", matches[0])
			return matches[1] + " " + matches[len(matches)-1]
		}
		return s
	})

	return result
}
