package queue

import "fmt"

func ExampleLinkedListQueue() {
	ll := NewLinkedListQueue[int]()

	fmt.Println(ll.Size())

	ll.Enqueue(1)
	ll.Enqueue(2)
	ll.Enqueue(3)

	fmt.Println(ll.IsEmpty())

	fmt.Println(ll.Dequeue())
	fmt.Println(ll.Dequeue())
	fmt.Println(ll.Dequeue())

	// Output:
	// 0
	// false
	// 1
	// 2
	// 3
}

func ExampleDynamicArrayQueue() {
	ll := NewDynamicArrayQueue[int]()

	fmt.Println(ll.Size())

	ll.Enqueue(1)
	ll.Enqueue(2)
	ll.Enqueue(3)

	fmt.Println(ll.IsEmpty())

	fmt.Println(ll.Dequeue())
	fmt.Println(ll.Dequeue())
	fmt.Println(ll.Dequeue())

	// Output:
	// 0
	// false
	// 1
	// 2
	// 3
}
