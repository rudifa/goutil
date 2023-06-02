// custom formatting
package fmt

import "strings"

// CompactFmt formats a string,
// replacing newline characters with middle dot (·)
// and skipping tab characters
func CompactFmt(formatString string) string {
	result := strings.Builder{}

	for _, c := range formatString {
		switch c {
		case '\n', '\u0085', '\u2028', '\u2029':
			result.WriteRune('·')
		case '\t':
			// Skip tab characters
			continue
		default:
			result.WriteRune(c)
		}
	}

	return result.String()
}
