package priorityqueue

import (
	"container/heap"
)

type subHeap[T any] struct {
	slicePtr *[]T
	less     func(i, j T) bool
}

var _ heap.Interface = subHeap[int]{}

func (sh subHeap[T]) Len() int {
	return len(*sh.slicePtr)
}

func (sh subHeap[T]) Swap(i, j int) {
	arr := *sh.slicePtr
	arr[i], arr[j] = arr[j], arr[i]
}

func (sh subHeap[T]) Less(i, j int) bool {
	return sh.less((*sh.slicePtr)[i], (*sh.slicePtr)[j])
}

func (sh subHeap[T]) Push(x any) {
	*sh.slicePtr = append(*sh.slicePtr, x.(T))
}

func (sh subHeap[T]) Pop() any {
	arr := *sh.slicePtr
	opt := arr[len(arr)-1]
	*sh.slicePtr = arr[:len(arr)-1]
	return opt
}
