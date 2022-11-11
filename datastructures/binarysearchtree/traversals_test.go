package binarysearchtree

import "fmt"

func ExamplePreorderTraversal() {
	bst := NewBinarySearchTree[int]()

	bst.Add(11)
	bst.Add(6)
	bst.Add(3)
	bst.Add(8)
	bst.Add(15)
	bst.Add(13)
	bst.Add(17)
	bst.Add(3)
	bst.Add(1)
	bst.Add(5)
	bst.Add(12)
	bst.Add(14)
	bst.Add(19)

	t := NewPreorderTraversalInt(bst.root)
	for t.HasNext() {
		fmt.Println(t.Next())
	}

	// Output:
	// 11
	// 6
	// 3
	// 1
	// 5
	// 8
	// 15
	// 13
	// 12
	// 14
	// 17
	// 19
}

func ExampleInorderTraversal() {
	bst := NewBinarySearchTree[int]()

	bst.Add(11)
	bst.Add(6)
	bst.Add(8)
	bst.Add(15)
	bst.Add(13)
	bst.Add(17)
	bst.Add(3)
	bst.Add(1)
	bst.Add(5)
	bst.Add(12)
	bst.Add(14)
	bst.Add(19)

	t := NewInorderTraversalInt(bst.root)
	for t.HasNext() {
		fmt.Println(t.Next())
	}

	// Output:
	// 1
	// 3
	// 5
	// 6
	// 8
	// 11
	// 12
	// 13
	// 14
	// 15
	// 17
	// 19
}

func ExamplePostorderTraversal() {
	bst := NewBinarySearchTree[int]()

	bst.Add(11)
	bst.Add(6)
	bst.Add(3)
	bst.Add(8)
	bst.Add(15)
	bst.Add(13)
	bst.Add(17)
	bst.Add(3)
	bst.Add(1)
	bst.Add(5)
	bst.Add(12)
	bst.Add(14)
	bst.Add(19)

	t := NewPostorderTraversalInt(bst.root)
	for t.HasNext() {
		fmt.Println(t.Next())
	}

	// Output:
	// 1
	// 5
	// 3
	// 8
	// 6
	// 12
	// 14
	// 13
	// 19
	// 17
	// 15
	// 11
}

func ExampleLevelOrderTraversal() {
	bst := NewBinarySearchTree[int]()

	bst.Add(11)
	bst.Add(6)
	bst.Add(3)
	bst.Add(8)
	bst.Add(15)
	bst.Add(13)
	bst.Add(17)
	bst.Add(3)
	bst.Add(1)
	bst.Add(5)
	bst.Add(12)
	bst.Add(14)
	bst.Add(19)

	t := NewLevelOrderTraversalInt(bst.root)
	for t.HasNext() {
		fmt.Println(t.Next())
	}

	// Output:
	// 11
	// 6
	// 15
	// 3
	// 8
	// 13
	// 17
	// 1
	// 5
	// 12
	// 14
	// 19
}
