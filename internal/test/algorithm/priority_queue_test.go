package main

import (
	"testing"

	"github.com/M-Quadra/kameria/algorithm"
	priorityqueue "github.com/M-Quadra/kameria/algorithm/priority-queue"
	"github.com/stretchr/testify/require"
)

func TestPriorityQueue(t *testing.T) {
	pq := priorityqueue.New[int](nil, func(a, b int) bool {
		return a < b
	})
	_, ok := pq.Pop()
	require.False(t, ok)
	_, ok = pq.Top()
	require.False(t, ok)

	arr := []int{}
	pq = priorityqueue.New(&arr, func(a, b int) bool {
		return a < b
	})
	pq.Push(1)

	v, ok := pq.Pop()
	require.True(t, ok)
	require.Equal(t, 1, v)

	_, ok = pq.Pop()
	require.False(t, ok)

	arr = []int{1, 2, 3, 4, 5, 6}
	arrB := make([]int, len(arr))
	copy(arrB, arr)
	algorithm.Reverse(arrB, 0, len(arr))
	pq = priorityqueue.New(&arr, func(a, b int) bool {
		return a > b
	})

	v, ok = pq.Top()
	require.True(t, ok)
	require.Equal(t, arrB[0], v)
	for v, ok := pq.Pop(); ok; v, ok = pq.Pop() {
		require.Equal(t, arrB[0], v)
		arrB = arrB[1:]
	}
}
