package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3}

	Reverse(arr, 0, len(arr))
	assert.Equal(t, []int{3, 2, 1}, arr)

	Reverse(arr, 1, len(arr))
	assert.Equal(t, []int{3, 1, 2}, arr)

	{ // empty
		Reverse([]int{}, 1, len(arr))

		arr = nil
		Reverse(arr, 1, len(arr))
	}
}
