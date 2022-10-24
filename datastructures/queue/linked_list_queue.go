package queue

import (
	"fluffy-adventure/datastructures/linkedlist"
	"golang.org/x/exp/constraints"
)

type LinkedListQueue[T constraints.Ordered] struct {
	list linkedlist.SingleLinkedList[T]
}

func NewIntLinkedListQueue() *LinkedListQueue[int] {
	var l linkedlist.SingleLinkedList[int]
	return &LinkedListQueue[int]{
		list: l,
	}
}

func (l *LinkedListQueue[T]) Enqueue(v T) {
	l.list.AddLast(v)
}

func (l *LinkedListQueue[T]) Dequeue() T {
	return l.list.RemoveFirst()
}

func (l *LinkedListQueue[T]) IsEmpty() bool {
	return l.list.IsEmpty()
}

func (l *LinkedListQueue[T]) Peek() T {
	return l.list.PeekFirst()
}

func (l *LinkedListQueue[T]) Size() int {
	return l.list.Size()
}

func (l *LinkedListQueue[T]) Contains(v T) bool {
	return l.list.Contains(v)
}
