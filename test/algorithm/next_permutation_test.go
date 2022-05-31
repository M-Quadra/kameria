package main

import (
	"testing"

	"github.com/M-Quadra/kameria/algorithm"
	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {
	arr := []int{1, 2, 3}
	for _, v := range [][]int{
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
		{},
	} {
		ok := algorithm.NextPermutation(arr, 0, len(arr), func(a, b int) bool {
			return a < b
		})
		isFinale := len(v) <= 0
		if !isFinale {
			assert.Equal(t, arr, v)
		}
		assert.Equal(t, isFinale, !ok)
	}

	{ // empty
		ok := algorithm.NextPermutation([]int{}, 0, len(arr), func(a, b int) bool {
			return a < b
		})
		assert.False(t, ok)

		ok = algorithm.NextPermutation(arr, 0, len(arr), nil)
		assert.False(t, ok)

		arr = nil
		ok = algorithm.NextPermutation(arr, 0, len(arr), func(a, b int) bool {
			return a < b
		})
		assert.False(t, ok)
	}
}
