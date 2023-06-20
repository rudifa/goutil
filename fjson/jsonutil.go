package fjson

import (
	"bytes"
	"encoding/json"
)

// Prettyfmt returns a pretty formatted json string
func Prettyfmt(input string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(input), "", "\t"); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

// Compact returns a compacted json string
func Compact(jsonString string) (string, error) {
	dst := &bytes.Buffer{}
	err := json.Compact(dst, []byte(jsonString))
	if err != nil {
		return "", err
	}
	return dst.String(), nil
}
