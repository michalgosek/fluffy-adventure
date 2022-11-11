package binarysearchtree

import (
	"math"

	"fluffy-adventure/cmp"
	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	v           T
	left, right *Node[T]
}

type BinarySearchTree[T constraints.Ordered] struct {
	n    int
	root *Node[T]
}

func NewBinarySearchTree[T constraints.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

func (b *BinarySearchTree[T]) add(node *Node[T], v T) *Node[T] {
	if node == nil {
		node = &Node[T]{v: v}
	} else {
		if cmp.Compare(v, node.v) < 0 {
			node.left = b.add(node.left, v)
		} else {
			node.right = b.add(node.right, v)
		}
	}
	return node
}

func (b *BinarySearchTree[T]) findMin(node *Node[T]) *Node[T] {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (b *BinarySearchTree[T]) findMax(node *Node[T]) *Node[T] {
	for node.right != nil {
		node = node.right
	}
	return node
}

func (b *BinarySearchTree[T]) remove(node *Node[T], v T) *Node[T] {
	if node == nil {
		return nil
	}
	compare := cmp.Compare(v, node.v)
	if compare < 0 {
		node.left = b.remove(node.left, v)
	} else if compare > 0 {
		node.right = b.remove(node.right, v)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			tmp := b.findMin(node.right)

			node.v = tmp.v
			node.right = b.remove(node.right, tmp.v)
		}
	}
	return node
}

func (b *BinarySearchTree[T]) Remove(v T) bool {
	if !b.contains(b.root, v) {
		return false
	}
	b.root = b.remove(b.root, v)
	b.n--
	return true
}

func (b *BinarySearchTree[T]) Add(v T) bool {
	if b.contains(b.root, v) {
		return false
	}
	b.root = b.add(b.root, v)
	b.n++
	return true
}

func (b *BinarySearchTree[T]) Contains(v T) bool {
	return b.contains(b.root, v)
}

func (b *BinarySearchTree[T]) contains(node *Node[T], v T) bool {
	if node == nil {
		return false
	}

	compare := cmp.Compare(v, node.v)
	if compare < 0 {
		return b.contains(node.left, v)
	} else if compare > 0 {
		return b.contains(node.right, v)
	}
	return true
}

func (b *BinarySearchTree[T]) IsEmpty() bool {
	return b.Size() == 0
}

func (b *BinarySearchTree[T]) Size() int {
	return b.n
}

func (b *BinarySearchTree[T]) Height() int {
	return b.height(b.root)
}

func (b *BinarySearchTree[T]) height(node *Node[T]) int {
	if node == nil {
		return 0
	}
	left := float64(b.height(node.left))
	right := float64(b.height(node.right))
	max := int(math.Max(left, right)) + 1
	return max
}
