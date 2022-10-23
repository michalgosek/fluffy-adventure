package linked_list

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type SingleLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func (s *SingleLinkedList[T]) Size() int {
	return s.size
}

func (s *SingleLinkedList[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *SingleLinkedList[T]) AddFirst(v T) {
	node := Node[T]{Value: v, Next: s.head}
	if s.IsEmpty() {
		s.head = &node
		s.tail = &node
	} else {
		s.tail = s.head
		s.head = &node
	}
	s.size++
}

func (s *SingleLinkedList[T]) AddLast(v T) {
	node := Node[T]{Value: v}
	if s.IsEmpty() {
		s.head = &node
	} else {
		s.tail.Next = &node
		s.tail = &node
	}
	s.size++
}

func (s *SingleLinkedList[T]) Clear() bool {
	trav := s.head

	for trav != nil {
		var empty T
		next := trav.Next
		trav.Next, trav.Value = nil, empty
		trav = next
	}

	s.tail = nil
	s.head = nil
	s.size = 0
	return true
}

func (s *SingleLinkedList[T]) GetAt(idx int) T {
	if idx < 0 || idx > s.size {
		panic("illegal index")
	}

	trav := s.head
	for i := 0; i != idx; i++ {
		trav = trav.Next
	}
	return trav.Value
}

func (s *SingleLinkedList[T]) AddAt(idx int, v T) {
	if idx < 0 || idx > s.size {
		panic("illegal index")
	}

	if idx == 0 {
		s.AddFirst(v)
		return
	}
	if idx == s.size {
		s.AddLast(v)
		return
	}

	trav := s.head
	for i := 0; i < idx-1; i++ {
		trav = trav.Next
	}
	node := Node[T]{Value: v, Next: trav.Next}
	trav.Next = &node
	s.size++
}

func (s *SingleLinkedList[T]) RemoveAt(idx int) {
	if idx < 0 || idx >= s.size {
		panic("illegal index")
	}
	if idx == 0 {
		s.RemoveFirst()
		return
	}

	trav := s.head
	for i := 0; i < idx-1; i++ {
		trav = trav.Next
	}

	if idx == s.size-1 {
		s.tail = trav
		trav.Next = nil
		s.size--
		return
	}

	trav.Next = trav.Next.Next
	s.size--
}

func (s *SingleLinkedList[T]) RemoveFirst() {
	newHead := s.head.Next
	s.head = newHead
	s.size--
}

func (s *SingleLinkedList[T]) RemoveLast() {
	s.RemoveAt(s.size - 1)
}

func (s *SingleLinkedList[T]) PeekFirst() T {
	return s.head.Value
}

func (s *SingleLinkedList[T]) PeekLast() T {
	return s.tail.Value
}
