package klist

// Element in List
type Element[T any] struct {
	next, prev *Element[T]
	list       *List[T]
	Value      T
}

// Next returns the next list element or nil.
func (e *Element[T]) Next() *Element[T] {
	if p := e.next; p != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev returns the previous list element or nil.
func (e *Element[T]) Prev() *Element[T] {
	if p := e.prev; p != nil && p != &e.list.root {
		return p
	}
	return nil
}
