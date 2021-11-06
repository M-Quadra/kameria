package main

import (
	"testing"

	"github.com/M-Quadra/kameria/algorithm"
	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {
	ary := []int{1, 2, 3}
	assert.False(t, algorithm.NextPermutation(nil, 0, len(ary), func(i, j int) bool {
		return true
	}))
	assert.False(t, algorithm.NextPermutation(ary, 0, len(ary), nil))

	ok := algorithm.NextPermutation(ary, 0, len(ary), func(i, j int) bool {
		return ary[i] < ary[j]
	})
	assert.True(t, ok)
	assert.Equal(t, []int{1, 3, 2}, ary)

	ok = algorithm.NextPermutation(&ary, 0, len(ary), func(i, j int) bool {
		return ary[i] < ary[j]
	})
	assert.True(t, ok)
	assert.Equal(t, []int{2, 1, 3}, ary)
}
