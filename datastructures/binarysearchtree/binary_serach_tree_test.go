package binarysearchtree

import "fmt"

func ExampleBinarySearchTree() {
	bst := NewIntBinarySearchTree()

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

func ExampleTreeTraversals_PreOrder() {
	bst := NewIntBinarySearchTree()

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

	bst.PreOrder()

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

func ExampleTreeTraversals_InOrder() {
	bst := NewIntBinarySearchTree()

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

	bst.InOrder()

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

func ExampleTreeTraversals_PostOrder() {
	bst := NewIntBinarySearchTree()

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

	bst.PostOrder()

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

func ExampleTreeTraversals_LevelOrder() {
	bst := NewIntBinarySearchTree()

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

	bst.LevelOrder()

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
