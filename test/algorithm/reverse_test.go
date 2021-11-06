package main

import (
	"testing"

	"github.com/M-Quadra/kameria/algorithm"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	algorithm.Reverse(nil, 0, 2)
	tmp := map[int]struct{}{}
	algorithm.Reverse(tmp, 0, 2)

	ary := []int{1, 2, 3}

	algorithm.Reverse(&ary, 0, len(ary))
	assert.Equal(t, []int{3, 2, 1}, ary)

	algorithm.Reverse(ary, 1, len(ary))
	assert.Equal(t, []int{3, 1, 2}, ary)
}
