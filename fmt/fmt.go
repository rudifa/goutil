// custom formatting
package fmt

import (
	"regexp"
	"strings"
)

// CompactFmt formats a string, reducing the whitespace to a minimum
// (1) replaces every newline character by a middle dot;
// (2) replaces any adjacent spaces or tabs following the replaced newline by nothing,
// (3) replaces any other adjacent adjacent spaces or tabs by a single space
func CompactFmt(formatString string) string {
	replacer := strings.NewReplacer("\n", "·")

	result := replacer.Replace(formatString)

	re := regexp.MustCompile(`"[^"]*"`)
	matches := re.FindAllStringIndex(result, -1)

	for i := len(matches) - 1; i >= 0; i-- {
		start, end := matches[i][0], matches[i][1]
		substr := result[start:end]

		replaceNewline := strings.ReplaceAll(substr, "\n", "·")
		result = result[:start] + replaceNewline + result[end:]
	}

	re = regexp.MustCompile(`[ \t]+`)
	result = re.ReplaceAllStringFunc(result, func(match string) string {
		if strings.Contains(match, "\n") {
			return strings.ReplaceAll(match, "\n", "")
		}
		return " "
	})

	result = strings.ReplaceAll(result, "· ", "·")

	return result
}
