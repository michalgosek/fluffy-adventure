package linkedlist

import "fmt"

func Example_singleLinkedList() {
	var l SingleLinkedList[int]

	fmt.Println(l.IsEmpty())

	l.AddAt(0, 0)
	l.AddAt(1, 1)
	l.AddAt(2, 2)
	l.AddAt(3, 4)

	fmt.Println(l.PeekFirst())
	fmt.Println(l.PeekLast())
	fmt.Println(l.GetAt(2))

	fmt.Println(l.IndexOf(2))
	fmt.Println(l.Contains(4))
	l.RemoveAt(2)
	fmt.Println(l.Size())
	fmt.Println(l.GetAt(2))
	fmt.Println(l.RemoveFirst())
	fmt.Println(l.RemoveLast())

	l.AddFirst(9)
	l.AddLast(10)
	fmt.Println(l.PeekFirst())
	fmt.Println(l.PeekLast())
	fmt.Println(l.Clear())

	// Output:
	// true
	// 0
	// 4
	// 2
	// 2
	// true
	// 3
	// 4
	// 0
	// 4
	// 9
	// 10
	// true
}
