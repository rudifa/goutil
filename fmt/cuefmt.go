package fmt

import (
	"fmt"

	"cuelang.org/go/cue"
)

// CompactCueVal formats a cue.Value as a string,
// replacing newline characters with middle dot (·)
// and squashing the whitespace to a minimum
func CompactCueVal(v cue.Value) string {

	str := fmt.Sprintf("%v", v)

	return CompactFmt(str)
}
