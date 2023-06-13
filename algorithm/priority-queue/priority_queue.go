package priorityqueue

import (
	"container/heap"
)

// Queue std::priority_queue
type Queue[T any] struct {
	sub subHeap[T]
}

// Size len(slice)
func (q Queue[T]) Size() int {
	return q.sub.Len()
}

// Empty len(slice) == 0
func (q Queue[T]) IsEmpty() bool {
	return q.Size() <= 0
}

// Push priority_queue.push(x)
func (q Queue[T]) Push(x T) {
	heap.Push(q.sub, x)
}

// Pop priority_queue.pop() with result
func (q Queue[T]) Pop() T {
	return heap.Pop(q.sub).(T)
}

// Top slice[0]
func (q Queue[T]) Top() T {
	return (*q.sub.slicePtr)[0]
}

/*
New PriorityQueue

slicePtr can be nil, then it will be initialized as []T{}

less can't be nil, it's the compare function
less(i, j int) { return a[i]<b[i] } is min-heap
*/
func New[T any](slicePtr *[]T, less func(a, b T) bool) *Queue[T] {
	if slicePtr == nil {
		slicePtr = &[]T{}
	}
	if less == nil {
		return nil
	}

	q := Queue[T]{
		sub: subHeap[T]{
			slicePtr: slicePtr,
			less:     less,
		},
	}
	heap.Init(q.sub)
	return &q
}
