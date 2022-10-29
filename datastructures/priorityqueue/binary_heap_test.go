package priorityqueue

import "fmt"

func ExampleNewBinaryHeapInts() {
	heap := NewBinaryHeapInts(-1, 5, 1, 8, 6, 2)

	fmt.Println(heap.IsMinHeap(0))
	fmt.Println(heap.Size())
	fmt.Println(heap.Peek())

	heap.Poll()
	fmt.Println(heap.IsMinHeap(0))
	fmt.Println(heap.Size())
	fmt.Println(heap.Peek())

	heap.Poll()
	fmt.Println(heap.Peek())

	heap.Remove(8)
	fmt.Println(heap.IsMinHeap(0))
	fmt.Println(heap.Size())

	fmt.Println(heap.Contains(6))
	fmt.Println(heap.Contains(333))

	// Output:
	// true
	// 6
	// -1
	// true
	// 5
	// 1
	// 2
	// true
	// 3
	// true
	// false
}
