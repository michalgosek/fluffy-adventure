package queue

import (
	"fluffy-adventure/datastructures/arrays"
	"golang.org/x/exp/constraints"
)

type DynamicArrayQueue[T constraints.Ordered] struct {
	arr *arrays.DynamicArray[T]
}

func NewDynamicArrayIntQueue() *DynamicArrayQueue[int] {
	return &DynamicArrayQueue[int]{
		arr: arrays.NewDynamicIntArray(),
	}
}

func (d *DynamicArrayQueue[T]) Peek() T {
	return d.arr.Get(0)
}

func (d *DynamicArrayQueue[T]) Enqueue(v T) {
	d.arr.Append(v)
}

func (d *DynamicArrayQueue[T]) Dequeue() T {
	return d.arr.RemoveAt(0)
}

func (d *DynamicArrayQueue[T]) Size() int {
	return d.arr.Size()
}

func (d *DynamicArrayQueue[T]) IsEmpty() bool {
	return d.arr.IsEmpty()
}
