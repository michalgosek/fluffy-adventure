package priorityqueue

import (
	"fluffy-adventure/datastructures/arrays"
	"golang.org/x/exp/constraints"
	"math"
)

type BinaryHeap[T constraints.Ordered] struct {
	heap *arrays.DynamicArray[T]
}

func (b *BinaryHeap[T]) swap(i, j int) {
	first := b.heap.Get(i)
	second := b.heap.Get(j)
	b.heap.Set(i, second)
	b.heap.Set(j, first)
}

// less checks whether the heap min invariant is kept between parent and child nodes stored at i, j indices.
func (b *BinaryHeap[T]) less(i, j int) bool {
	first := b.heap.Get(i)
	second := b.heap.Get(j)
	return b.heap.CompareTo(first, second) <= 0
}

func (b *BinaryHeap[T]) swim(k int) {
	parent := (k - 1) / 2
	for k > 0 && b.less(k, parent) {
		b.swap(parent, k)
		k = parent
		parent = (k - 1) / 2
	}
}

func (b *BinaryHeap[T]) sink(k int) {
	heapSize := b.Size()
	for {
		left := 2*k + 1
		right := 2*k + 2
		smallest := left

		if right < heapSize && b.less(right, left) {
			smallest = right
		}
		if left >= heapSize || b.less(k, smallest) {
			break
		}

		b.swap(smallest, k)
		k = smallest
	}
}

func (b *BinaryHeap[T]) removeAt(i int) T {
	if b.IsEmpty() {
		var empty T
		return empty
	}

	indexOfLastElem := b.Size() - 1
	removedData := b.heap.Get(i)
	b.swap(i, indexOfLastElem)

	b.heap.RemoveAt(indexOfLastElem)

	if i == indexOfLastElem {
		return removedData
	}

	elem := b.heap.Get(i)
	b.sink(i)

	if b.heap.Get(i) == elem {
		b.swim(i)
	}
	return removedData
}

func (b *BinaryHeap[T]) IsEmpty() bool {
	return b.heap.Size() == 0
}

func (b *BinaryHeap[T]) Size() int {
	return b.heap.Size()
}

func (b *BinaryHeap[T]) Contains(v T) bool {
	for i := 0; i < b.Size(); i++ {
		if b.heap.Get(i) == v {
			return true
		}
	}
	return false
}

func (b *BinaryHeap[T]) Add(v T) {
	var empty T
	if v == empty {
		panic("Illegal argument")
	}

	b.heap.Append(v)
	idxOfLastElem := b.Size() - 1
	b.swim(idxOfLastElem)
}

func (b *BinaryHeap[T]) Peek() T {
	if b.IsEmpty() {
		var noop T
		return noop
	}
	return b.heap.Get(0)
}

func (b *BinaryHeap[T]) Poll() T {
	return b.removeAt(0)
}

func (b *BinaryHeap[T]) Clear() {
	b.heap.Clear()
}

func (b *BinaryHeap[T]) Remove(v T) bool {
	var empty T
	if v == empty {
		return false
	}
	for i := 0; i < b.Size(); i++ {
		if v == b.heap.Get(i) {
			b.removeAt(i)
			return true
		}
	}
	return false
}

func (b *BinaryHeap[T]) IsMinHeap(k int) bool {
	size := b.Size()
	if k >= size {
		return true
	}

	left := 2*k + 1
	right := 2*k + 2

	if left < size && !b.less(k, left) {
		return false
	}
	if right < size && !b.less(k, right) {
		return false
	}
	return b.IsMinHeap(left) && b.IsMinHeap(right)
}

func NewBinaryHeapInts(nums ...int) *BinaryHeap[int] {
	heapSize := len(nums)

	b := BinaryHeap[int]{
		heap: arrays.NewDynamicIntArray(),
	}
	for i := 0; i < len(nums); i++ {
		b.heap.Append(nums[i])
	}

	// Heapify process, O(n)
	for i := math.Max(0, float64((heapSize/2)-1)); i >= 0; i-- {
		b.sink(int(i))
	}
	return &b
}
