package main

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type SLList[T comparable] struct {
	head   *Node[T]
	last   *Node[T]
	length int
}

func NewSLList[T comparable]() *SLList[T] {
	return &SLList[T]{nil, nil, 0}
}

func (list *SLList[T]) Contains(item T) bool {
	return list.IndexOf(item) != -1
}

func (list *SLList[T]) IndexOf(item T) int {
	if list.IsEmpty() {
		return -1
	}

	curr := list.head
	for i := 0; i < list.length; i++ {
		if item == curr.value {
			return i
		}
		curr = curr.next
	}

	return -1
}

func (list *SLList[T]) IsEmpty() bool {
	return list.length == 0
}

func (list *SLList[T]) Head() (*Node[T], error) {
	if list.IsEmpty() {
		return nil, errors.New("IndexError: the list is empty")
	}

	return list.head, nil
}

func (list *SLList[T]) Tail() *Node[T] {
	return list.head.next
}

func (list *SLList[T]) Peek() (*T, error) {
	if list.IsEmpty() {
		return nil, errors.New("IndexError: the list is empty")
	}

	return &list.head.value, nil
}

func (list *SLList[T]) PeekLast() (*T, error) {
	if list.IsEmpty() {
		return nil, errors.New("IndexError: the list is empty")
	}

	return &list.last.value, nil
}

func (list *SLList[T]) Pop() (*T, error) {
	if list.IsEmpty() {
		return nil, errors.New("IndexError: the list is empty")
	}

	n, err := list.Scan(list.length - 2)

	if err != nil {
		return nil, err
	}

	out := n.next.value
	n.next = nil
	list.length--

	return &out, nil
}

func (list *SLList[T]) PopFront() (*T, error) {
	if list.IsEmpty() {
		return nil, errors.New("IndexError: the list is empty")
	}

	out := list.head.value
	list.head = list.head.next
	list.length--
	return &out, nil
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

func (list *SLList[T]) PushFront(item T) {
	if list.IsEmpty() {
		list.head = &Node[T]{item, nil}
		list.last = list.head
	} else {
		temp := list.head
		list.head = &Node[T]{item, temp}
	}
	list.length++
}

func (list *SLList[T]) Remove(item T) error {
	idx := list.IndexOf(item)
	return list.RemoveAtIndex(idx)
}

func (list *SLList[T]) RemoveAtIndex(idx int) error {
	if idx == -1 {
		return nil
	}

	if idx == 0 {
		list.length--
		list.PopFront()
		return nil
	}

	if idx == list.length-1 {
		list.length--
		list.Pop()
		return nil
	}

	prev, err := list.Scan(idx - 1)
	if err != nil {
		return err
	}
	curr, err := list.Scan(idx)
	if err != nil {
		return err
	}

	list.length--
	prev.next = curr.next
	return nil
}

func (list *SLList[T]) Debug() {
	curr := list.head
	for i := 0; i < list.length; i++ {
		fmt.Printf("[%v]->", curr.value)
		curr = curr.next
	}
	fmt.Println("nil")
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
