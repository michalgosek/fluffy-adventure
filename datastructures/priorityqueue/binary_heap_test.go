package priorityqueue

import "fmt"

func ExampleNewBinaryHeapInts() {
	heap := NewBinaryHeapInts(1, 5, 1, 8, 6, 2)

	fmt.Println(heap.IsMinHeap(0))
	heap.Poll()

	fmt.Println(heap.Size())
	fmt.Println(heap.IsMinHeap(0))
	fmt.Println(heap.Peek())
	fmt.Println(heap.Contains(8))
	fmt.Println(heap.Contains(11))

	heap.Add(-1)
	fmt.Println(heap.Contains(-1))
	fmt.Println(heap.Peek())
	fmt.Println(heap.IsMinHeap(0))
	fmt.Println(heap.Size())

	heap.Clear()
	fmt.Println(heap.Size())

	// Output:
	// true
	// 5
	// true
	// 1
	// true
	// false
	// true
	// -1
	// true
	// 6
	// 0
}
