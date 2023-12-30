package util

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {

	assert.True(t, Contains([]string{"a", "b", "c"}, "a"))
	assert.True(t, Contains([]string{"a", "b", "c"}, "b"))
	assert.True(t, Contains([]string{"a", "b", "c"}, "c"))
	assert.False(t, Contains([]string{"a", "b", "c"}, "d"))
}

func TestMap(t *testing.T) {

	assert.Equal(t, []string{"a", "b", "c"}, Map([]string{"A", "B", "C"}, func(s string) string {
		return strings.ToLower(s)
	}))

	assert.Equal(t, []int{1, 4, 9}, Map([]int{1, 2, 3}, func(i int) int {
		return i * i
	}))
}

func TestAreEqual(t *testing.T) {

	assert.True(t, AreEqual([]int{1, 2, 3}, []int{1, 2, 3}))
	assert.False(t, AreEqual([]int{1, 2, 3}, []int{1, 2, 3, 4}))
	assert.True(t, AreEqual([]string{"a", "b", "c"}, []string{"a", "b", "c"}))
	assert.False(t, AreEqual([]string{"a", "b", "c"}, []string{"a", "b", "c", "d"}))
	assert.False(t, AreEqual([]string{"a", "b", "c"}, []string{"a", "b", "d"}))
}

func TestReverseSlice(t *testing.T) {

	s := []string{"a", "b", "c"}
	ReverseSlice(s)
	assert.Equal(t, []string{"c", "b", "a"}, s)

	s = []string{"a", "b", "c", "d"}
	ReverseSlice(s)
	assert.Equal(t, []string{"d", "c", "b", "a"}, s)

	// ints
	i := []int{1, 2, 3}
	ReverseSlice(i)
	assert.Equal(t, []int{3, 2, 1}, i)

}
