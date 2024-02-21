package main

import (
	"errors"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type SLList[T any] struct {
	head   *Node[T]
	last   *Node[T]
	length int
}

func NewSLList[T any]() *SLList[T] {
	return &SLList[T]{nil, nil, 0}
}

func (list *SLList[T]) IsEmpty() bool {
	return list.length == 0
}

func (list *SLList[T]) Push(item T) {
	if list.IsEmpty() {
		list.head = &Node[T]{item, nil}
		list.last = list.head
	} else {
		temp := list.last
		list.last = &Node[T]{item, nil}
		temp.next = list.last
	}
	list.length++
}

func (list *SLList[T]) Pop() (*T, error) {
	if list.IsEmpty() {
		return nil, errors.New("IndexError: the list is empty.")
	} else {
		n, err := list.Scan(list.length - 2)
		if err != nil {
			return nil, err
		}
		out := n.next.value
		n.next = nil
		list.length--
		return &out, nil
	}
}

func (list *SLList[T]) Scan(idx int) (*Node[T], error) {
	if idx >= list.length {
		return nil, errors.New("IndexError: index out of bounds")
	}
	curr := list.head
	for i := 0; i < idx; i++ {
		curr = curr.next
	}
	return curr, nil
}

// AddAtIndex
// PopFront
// Pop
// Contains
// Peek
// PeekLast
// Push
// PushFront
// Remove
// Head
// Tail
// IndexOf
// IsEmpty
