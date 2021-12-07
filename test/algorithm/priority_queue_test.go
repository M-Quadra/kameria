package main

import (
	"testing"

	"github.com/M-Quadra/kameria/algorithm"
	priorityqueue "github.com/M-Quadra/kameria/algorithm/priority-queue"
	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	var arr []int
	pq, err := priorityqueue.New(&arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	assert.Nil(t, err)
	pq.Push(1)
	assert.Equal(t, 1, pq.Pop())
	assert.Equal(t, nil, pq.Pop())

	arr = []int{1, 2, 3, 4, 5, 6}
	arrB := make([]int, len(arr))
	copy(arrB, arr)
	algorithm.Reverse(arrB, 0, len(arr))
	pq, _ = priorityqueue.New(&arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	assert.Equal(t, arrB[0], pq.Top())
	for _, v := range arrB {
		assert.Equal(t, v, pq.Pop())
	}
}
