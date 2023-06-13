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

	arr := []int{}
	pq = priorityqueue.New(&arr, func(a, b int) bool {
		return a < b
	})
	pq.Push(1)

	require.Equal(t, 1, pq.Pop())

	arr = []int{1, 2, 3, 4, 5, 6}
	arrB := make([]int, len(arr))
	copy(arrB, arr)
	algorithm.Reverse(arrB, 0, len(arr))
	pq = priorityqueue.New(&arr, func(a, b int) bool {
		return a > b
	})

	require.Equal(t, arrB[0], pq.Top())
	for !pq.IsEmpty() {
		require.Equal(t, arrB[0], pq.Pop())
		arrB = arrB[1:]
	}
}
