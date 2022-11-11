package arrays

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type DynamicArray[T constraints.Ordered] struct {
	cap int
	arr []T
	l   int
}

func (d *DynamicArray[T]) Append(v T) {
	if d.l+1 >= d.cap {
		d.cap = d.cap * 2
		arr2 := make([]T, d.cap)
		for i := 0; i < d.l; i++ {
			arr2[i] = d.arr[i]
		}
		d.arr = arr2
	}
	d.arr[d.l] = v
	d.l++
}

func (d *DynamicArray[T]) Clear() {
	var empty T
	for i := 0; i < d.l; i++ {
		d.arr[i] = empty
	}
	d.l = 0
}

func (d *DynamicArray[T]) Reverse() {
	for i := 0; i < d.l/2; i++ {
		tmp := d.arr[i]
		d.arr[i] = d.arr[d.l-i-1]
		d.arr[d.l-i-1] = tmp
	}
}

func (d *DynamicArray[T]) RemoveAt(idx int) T {
	if idx >= d.l || idx < 0 {
		panic("index out of range")
	}
	data := d.arr[idx]
	newArr := make([]T, d.l-1)
	for i, j := 0, 0; i < d.l; i, j = i+1, j+1 {
		if i == idx {
			j--
			continue
		}
		newArr[j] = d.arr[i]
	}
	d.arr = newArr
	d.l--
	d.cap = d.l
	return data
}

func (d *DynamicArray[T]) Remove(v T) bool {
	for i := 0; i < d.l; i++ {
		if d.arr[i] == v {
			d.RemoveAt(i)
			return true
		}
	}
	return false
}

func (d *DynamicArray[T]) Get(idx int) T {
	return d.arr[idx]
}

func (d *DynamicArray[T]) Set(idx int, v T) {
	d.arr[idx] = v
}

func (d *DynamicArray[T]) Size() int {
	return d.l
}

func (d *DynamicArray[T]) IsEmpty() bool {
	return d.l == 0
}

func (d *DynamicArray[T]) Sort() {
	slices.Sort(d.arr)
}

func NewDynamicArray[T constraints.Ordered]() *DynamicArray[T] {
	d := DynamicArray[T]{cap: 1}
	return &d
}
