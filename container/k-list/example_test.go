package klist_test

import (
	"fmt"

	klist "github.com/M-Quadra/kameria/container/k-list"
)

func Example() {
	// Create a new list and put some numbers in it.
	l := klist.New[int]()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
}
