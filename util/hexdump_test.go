package util_test

import (
	"testing"

	"github.com/rudifa/goutil/util"
)

func TestHexDump(t *testing.T) {

	// Test: 0 bytes
	{
		data := []byte{}
		expected := ""
		actual := util.HexDump(data)
		if actual != expected {
			t.Errorf("Test failed.\nExpected:\n%s\nActual:\n%s", expected, actual)
		}
	}

	// Test 20 bytes
	{
		data := []byte{0x01, 0x02, 0x03, 0x04, 0x05,
			0x06, 0x07, 0x08, 0x09, 0x0a,
			0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
			0x10, 0x11, 0x12, 0x13, 0x14}

		expected := "00000000: 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f 10  ................\n" +
			"00000010: 11 12 13 14                                      ....\n"

		actual := util.HexDump(data)
		if actual != expected {
			t.Errorf("Test failed.\nExpected:\n%s\nActual:\n%s", expected, actual)
		}
	}

	// Test 20 bytes starting from offset 50, hex 0x32
	{
		data := []byte{0x32, 0x33, 0x34, 0x35, 0x36,
			0x37, 0x38, 0x39, 0x3a, 0x3b,
			0x3c, 0x3d, 0x3e, 0x3f, 0x40,
			0x41, 0x42, 0x43, 0x44, 0x45}

		expected := "00000000: 32 33 34 35 36 37 38 39 3a 3b 3c 3d 3e 3f 40 41  23456789:;<=>?@A\n" +
			"00000010: 42 43 44 45                                      BCDE\n"

		actual := util.HexDump(data)
		if actual != expected {
			t.Errorf("Test failed.\nExpected:\n%s\nActual:\n%s", expected, actual)
		}
	}
}
