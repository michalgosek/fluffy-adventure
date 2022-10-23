package linked_list

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
	l.RemoveAt(2)
	fmt.Println(l.GetAt(2))
	fmt.Println(l.Size())
	fmt.Println(l.Clear())

	// Output:
	// true
	// 0
	// 4
	// 2
	// 4
	// 3
	// true
}
