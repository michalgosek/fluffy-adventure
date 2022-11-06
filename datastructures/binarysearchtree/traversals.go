package binarysearchtree

import (
	"golang.org/x/exp/constraints"
)

type PreorderTraversal[T constraints.Ordered] struct {
	root  *Node[T]
	stack []*Node[T]
}

func NewPreorderTraversalInt(root *Node[int]) *PreorderTraversal[int] {
	p := PreorderTraversal[int]{
		root:  root,
		stack: []*Node[int]{root},
	}
	return &p
}

func (p *PreorderTraversal[T]) HasNext() bool {
	return p.root != nil && len(p.stack) != 0
}

func (p *PreorderTraversal[T]) Next() T {
	l := len(p.stack) - 1
	node := p.stack[l]
	p.stack = p.stack[:len(p.stack)-1]

	if node.right != nil {
		p.stack = append(p.stack, node.right)
	}
	if node.left != nil {
		p.stack = append(p.stack, node.left)
	}

	return node.v
}

type InorderTraversal[T constraints.Ordered] struct {
	stack []*Node[T]
	trav  *Node[T]
}

func NewInorderTraversalInt(root *Node[int]) *InorderTraversal[int] {
	p := InorderTraversal[int]{
		stack: []*Node[int]{root},
		trav:  root,
	}
	return &p
}

func (i *InorderTraversal[T]) HasNext() bool {
	return len(i.stack) != 0
}

func (i *InorderTraversal[T]) Next() T {
	for i.trav != nil && i.trav.left != nil {
		i.stack = append(i.stack, i.trav.left)
		i.trav = i.trav.left
	}

	lastIdx := len(i.stack) - 1
	node := i.stack[lastIdx]
	i.stack = i.stack[:lastIdx]

	if node.right != nil {
		i.stack = append(i.stack, node.right)
		i.trav = node.right
	}
	return node.v
}

type PostorderTraversal[T constraints.Ordered] struct {
	stack2 []*Node[T]
}

func (p *PostorderTraversal[T]) HasNext() bool {
	return len(p.stack2) != 0
}

func (p *PostorderTraversal[T]) Next() T {
	lastIdx := len(p.stack2) - 1
	node := p.stack2[lastIdx]
	p.stack2 = p.stack2[:lastIdx]
	return node.v
}

func NewPostorderTraversalInt(root *Node[int]) *PostorderTraversal[int] {
	stack1 := []*Node[int]{root}
	var stack2 []*Node[int]
	for len(stack1) != 0 {
		lastIdx := len(stack1) - 1
		node := stack1[lastIdx]

		if node != nil {
			stack2 = append(stack2, node)
			stack1 = stack1[:lastIdx]
			if node.left != nil {
				stack1 = append(stack1, node.left)
			}
			if node.right != nil {
				stack1 = append(stack1, node.right)
			}
		}
	}
	p := PostorderTraversal[int]{
		stack2: stack2,
	}
	return &p
}

type LevelOrderTraversal[T constraints.Ordered] struct {
	queue []*Node[T]
}

func NewLevelOrderTraversalInt(root *Node[int]) *LevelOrderTraversal[int] {
	p := LevelOrderTraversal[int]{
		queue: []*Node[int]{root},
	}
	return &p
}

func (l *LevelOrderTraversal[T]) HasNext() bool {
	return len(l.queue) != 0
}

func (l *LevelOrderTraversal[T]) Next() T {
	// poll
	node := l.queue[0]
	l.queue = l.queue[1:]

	if node.left != nil {
		l.queue = append(l.queue, node.left) // enqueue
	}
	if node.right != nil {
		l.queue = append(l.queue, node.right) // enqueue
	}
	return node.v
}
