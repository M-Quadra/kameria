package priorityqueue

import (
	"container/heap"
	"errors"
	"reflect"
)

// PriorityQueue std::priority_queue
type PriorityQueue struct {
	sub subHeap
}

// Size len(slice)
func (pq PriorityQueue) Size() int {
	return pq.sub.Len()
}

// Empty len(slice) == 0
func (pq PriorityQueue) Empty() bool {
	return pq.Size() <= 0
}

// Push priority_queue.push(x)
func (pq PriorityQueue) Push(x interface{}) {
	heap.Push(pq.sub, x)
}

// Pop priority_queue.pop() with result
func (pq PriorityQueue) Pop() interface{} {
	if pq.sub.Len() <= 0 {
		return nil
	}

	return heap.Pop(pq.sub)
}

// Top slice[0]
func (pq PriorityQueue) Top() interface{} {
	if pq.sub.Len() <= 0 {
		return nil
	}

	return pq.sub.slicePtr.Elem().Index(0)
}

// New need point of sliceP
//  less(i, j int) { return a[i]<b[i] } is min-heap
func New(slicePtr interface{}, less func(i, j int) bool) (*PriorityQueue, error) {
	rv := reflect.ValueOf(slicePtr)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Slice {
		return nil, errors.New("slicePtr must be a slice pointer")
	}
	if less == nil {
		return nil, errors.New("less must not be nil")
	}

	opt := &PriorityQueue{
		sub: subHeap{
			slicePtr: rv,
			less:     less,
		},
	}
	heap.Init(opt.sub)
	return opt, nil
}
