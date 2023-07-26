package util

import "fmt"

func HexDump(data []byte) string {
	var result string
	for i := 0; i < len(data); i += 16 {
		row := data[i:min(i+16, len(data))]
		hex := ""
		ascii := ""
		for j := 0; j < len(row); j++ {
			hex += fmt.Sprintf("%02x ", row[j])
			if row[j] >= 32 && row[j] <= 126 {
				ascii += string(row[j])
			} else {
				ascii += "."
			}
		}
		result += fmt.Sprintf("%08x: %-48s %s\n", i, hex, ascii)
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
