package binarysearchtree

import "fmt"

func ExampleBinarySearchTree() {
	bst := NewBinarySearchTree[int]()

	bst.Add(7)
	bst.Add(5)
	bst.Add(20)
	bst.Add(21)
	bst.Add(23)
	bst.Add(4)
	bst.Add(6)

	fmt.Println(bst.IsEmpty())
	fmt.Println(bst.Size())
	bst.Remove(5)
	fmt.Println(bst.Size())
	fmt.Println(bst.Contains(5))
	fmt.Println(bst.Height())

	// Output:
	// false
	// 7
	// 6
	// false
	// 4
}
