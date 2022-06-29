package algorithm

import "container/list"

// Queue std::queue
type Queue[T any] struct {
	deque *list.List
}

// NewQueue create a new Queue
func NewQueue[T any](x ...T) Queue[T] {
	que := Queue[T]{
		deque: list.New(),
	}
	for _, v := range x {
		que.deque.PushBack(v)
	}
	return que
}

func (que Queue[T]) Empty() bool {
	return que.deque.Len() == 0
}

func (que Queue[T]) Push(v T) {
	que.deque.PushBack(v)
}

func (que Queue[T]) Pop() T {
	v := que.deque.Front()
	que.deque.Remove(v)
	return v.Value.(T)
}

func (que Queue[T]) Size() int {
	return que.deque.Len()
}

func (que Queue[T]) Front() T {
	return que.deque.Front().Value.(T)
}

func (que Queue[T]) Back() T {
	return que.deque.Back().Value.(T)
}
