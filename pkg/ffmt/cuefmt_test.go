package ffmt_test

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

	cueval = ctx.CompileString(`
str: "hello world"
num: 42
flt: 3.14
"k8s.io/annotation": "secure-me"
`)

	assert.Equal(t, `{·str: "hello world"·num: 42·flt: 3.14·"k8s.io/annotation": "secure-me"·}`, fmt.CompactCueVal(cueval))
}
