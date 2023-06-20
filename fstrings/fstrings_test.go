package fstrings_test

import (
	"testing"

	"github.com/rudifa/goutil/fstrings"
)

//"github.com/rudifa/goutil/pkg/fstrings"

func TestTruncate(t *testing.T) {

	// sample string
	str := "This is a long string that will be truncated"

	// truncate the string to 10 runes
	truncated := fstrings.Truncate(str, 10)

	// test the truncated string
	expected := "This is a ..."
	if truncated != expected {
		t.Errorf("expected: '%s' got: '%s'", expected, truncated)
	}
}

func TestHeadAndTail(t *testing.T) {

	// sample string
	str := "This is a long string that will be shortened"

	// truncate the string to 10 + 10 runes
	shortened := fstrings.HeadAndTail(str, 10)

	// test the truncated string
	expected := "This is a ··· shortened"
	if shortened != expected {
		t.Errorf("expected: '%s' got: '%s'", expected, shortened)
	}
}
