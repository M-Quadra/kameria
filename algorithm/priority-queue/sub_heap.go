package priorityqueue

import (
	"container/heap"
	"reflect"
)

type subHeap struct {
	heap.Interface

	slicePtr reflect.Value
	less     func(i, j int) bool
}

func (sh subHeap) Len() int {
	return sh.slicePtr.Elem().Len()
}

func (sh subHeap) Swap(i, j int) {
	reflect.Swapper(sh.slicePtr.Elem().Interface())(i, j)
}

func (sh subHeap) Less(i, j int) bool {
	return sh.less(i, j)
}

func (sh subHeap) Push(x interface{}) {
	slice := reflect.Append(sh.slicePtr.Elem(), reflect.ValueOf(x))
	sh.slicePtr.Elem().Set(slice)
}

func (sh subHeap) Pop() interface{} {
	opt := sh.slicePtr.Elem().Index(sh.Len() - 1).Interface()
	sh.slicePtr.Elem().SetLen(sh.Len() - 1)
	return opt
}
