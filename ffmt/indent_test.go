package ffmt_test

import (
	"testing"

	"github.com/rudifa/goutil/ffmt"
	"github.com/stretchr/testify/assert"
)

func TestIndentNestedBrackets1(t *testing.T) {

	// a simple example, using `{}` as brackets

	const str = "{Hello{World}}"
	const want = `{
  Hello{
    World
  }
}`
	got, err := ffmt.IndentNestedBrackets(str, "{}", "  ")
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestIndentNestedBrackets2(t *testing.T) {

	// a cue debugStr example, using `<>` as brackets

	const str = "<[d0// 2423/2423.comment.in.empty.struct.BAD.cue] ex0: <[d1// empty h] <[d1// empty g] <[d1// empty f] <[d1// empty e] <[d1// empty d] <[d1// empty c] <[d0// empty a] [d1// empty b] {}&{}>&{}>&{}>&{}>&{}>&{}>&{}>>"
	const want = `<
  [d0// 2423/2423.comment.in.empty.struct.BAD.cue] ex0: <
    [d1// empty h] <
      [d1// empty g] <
        [d1// empty f] <
          [d1// empty e] <
            [d1// empty d] <
              [d1// empty c] <
                [d0// empty a] [d1// empty b] {}&{}
              >&{}
            >&{}
          >&{}
        >&{}
      >&{}
    >&{}
  >
>`
	got, err := ffmt.IndentNestedBrackets(str, "<>", "  ")
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestIndentNestedBrackets3(t *testing.T) {

	// a short example, using `[]` as brackets

	const str = "[a[b[c[d]]]]"
	const want = `[
   :a[
   :   :b[
   :   :   :c[
   :   :   :   :d
   :   :   :]
   :   :]
   :]
]`
	got, err := ffmt.IndentNestedBrackets(str, "[]", "   :")
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestIndentNestedBrackets4(t *testing.T) {

	// wrong brackets

	const str = "[a[b[c[d]]]]"
	_, err := ffmt.IndentNestedBrackets(str, "[]]", "   :")
	assert.Error(t, err)
}
