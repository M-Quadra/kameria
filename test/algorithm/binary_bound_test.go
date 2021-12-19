package main

import (
	"testing"

	"github.com/M-Quadra/kameria/algorithm"
	"github.com/stretchr/testify/assert"
)

func TestBinaryBound(t *testing.T) {
	arr := []int{0, 1, 2, 3, 5}
	// lower_bound
	i := algorithm.BinaryBound(0, len(arr), func(middle int) bool {
		return arr[middle] < 3
	})
	assert.Equal(t, 3, arr[i])
	// upper_bound
	i = algorithm.BinaryBound(0, len(arr), func(middle int) bool {
		return arr[middle] <= 3
	})
	assert.Equal(t, 5, arr[i])
}
