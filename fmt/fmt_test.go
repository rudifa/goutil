package fmt_test

import (
	"testing"

	"github.com/rudifa/goutil/fmt"
	"github.com/stretchr/testify/assert"
)

func TestCustomFormat(t *testing.T) {

	str := "\n\tA\nB\n"

	result := fmt.CompactFmt(str)
	t.Logf("str: %s", str)
	t.Logf("result: %s", result)

	assert.Equal(t, "·A·B·", result)
}
