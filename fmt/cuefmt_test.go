package fmt_test

import (
	"testing"

	"cuelang.org/go/cue/cuecontext"
	"github.com/rudifa/goutil/fmt"
	"github.com/stretchr/testify/assert"
)

func TestCompactCueVal(t *testing.T) {

	ctx := cuecontext.New()
	cueval := ctx.CompileString(`{
		"msg": "Hello world!"
		"bye": "And thanks for all the fish!"
	}`)

	assert.Equal(t, `{·msg: "Hello world!"·bye: "And thanks for all the fish!"·}`, fmt.CompactCueVal(cueval))
}
