package fjson_test

import (
	"encoding/hex"
	"log"
	"testing"

	"github.com/rudifa/goutil/fjson"
)

func TestCompactAndPrettyfmt(t *testing.T) {

	// sample compact json string
	compactJson := `{"key":"value","nested":{"nestedKey":"nested Value"}}`

	// pretty format the compactJson string
	pretty, err := fjson.Prettyfmt(compactJson)
	if err != nil {
		t.Errorf("error: %s", err)
	}

	// test the pretty formatted string
	expected := `{
	"key": "value",
	"nested": {
		"nestedKey": "nested Value"
	}
}`

	// run go test -v to see the hex dump of the strings
	log.Println("expected:\n", hex.Dump([]byte(expected)))
	log.Println("pretty:\n", hex.Dump([]byte(pretty)))

	if pretty != expected {
		t.Errorf("expected\n: '%s' got\n: '%s'", expected, pretty)
	}

	// re-compact the pretty string
	reCompacted, err := fjson.Compact(pretty)
	if err != nil {
		t.Errorf("error: %s", err)
	}

	// test the re-compacted string against the original compacted string
	if reCompacted != compactJson {
		t.Errorf("expected: '%s' got: '%s'", compactJson, reCompacted)
	}
}
