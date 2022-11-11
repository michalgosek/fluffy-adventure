package queue

import (
	"fluffy-adventure/datastructures/linkedlist"
	"golang.org/x/exp/constraints"
)

type LinkedListQueue[T constraints.Ordered] struct {
	list linkedlist.SingleLinkedList[T]
}

func NewLinkedListQueue[T constraints.Ordered]() *LinkedListQueue[T] {
	var l linkedlist.SingleLinkedList[T]
	return &LinkedListQueue[T]{
		list: l,
	}
}

func (l *LinkedListQueue[T]) Enqueue(v T) {
	l.list.AddLast(v)
}

func (l *LinkedListQueue[T]) Dequeue() any {
	return l.list.RemoveFirst()
}

func (l *LinkedListQueue[T]) IsEmpty() bool {
	return l.list.IsEmpty()
}

func (l *LinkedListQueue[T]) Peek() any {
	return l.list.PeekFirst()
}

func (l *LinkedListQueue[T]) Size() int {
	return l.list.Size()
}

func (l *LinkedListQueue[T]) Contains(v T) bool {
	return l.list.Contains(v)
}
