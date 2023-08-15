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
