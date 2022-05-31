package main

import (
	"testing"

	"github.com/M-Quadra/kameria/algorithm"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3}

	algorithm.Reverse(arr, 0, len(arr))
	assert.Equal(t, []int{3, 2, 1}, arr)

	algorithm.Reverse(arr, 1, len(arr))
	assert.Equal(t, []int{3, 1, 2}, arr)

	{ // empty
		algorithm.Reverse([]int{}, 1, len(arr))

		arr = nil
		algorithm.Reverse(arr, 1, len(arr))
	}
}
