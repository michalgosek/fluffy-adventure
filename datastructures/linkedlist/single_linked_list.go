package linkedlist

type Node[T any] struct {
	Value any
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
		s.AddFirst(v)
		return
	}

	s.tail.Next = &node
	s.tail = &node
	s.size++
}

func (s *SingleLinkedList[T]) Contains(v T) bool {
	return s.IndexOf(v) != -1
}

func (s *SingleLinkedList[T]) IndexOf(v any) int {
	if s.size == 0 {
		panic("empty list")
	}
	trav := s.head
	for i := 0; trav != nil; i++ {
		if trav.Value == v {
			return i
		}
		trav = trav.Next
	}
	return -1
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

func (s *SingleLinkedList[T]) Remove(v any) any {
	if v == nil {
		return nil
	}

	if s.head.Value == v {
		return s.RemoveFirst()
	}
	if s.tail.Value == v {
		return s.RemoveLast()
	}

	trav := s.head
	var i int
	for trav != nil {
		if trav.Value == v {
			break
		}
		trav = trav.Next
		i++
	}
	return s.RemoveAt(i)
}

func (s *SingleLinkedList[T]) GetAt(idx int) any {
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

func (s *SingleLinkedList[T]) RemoveAt(idx int) any {
	if idx < 0 || idx >= s.size {
		panic("illegal index")
	}
	if idx == 0 {
		return s.RemoveFirst()
	}

	trav := s.head
	for i := 0; i < idx-1; i++ {
		trav = trav.Next
	}

	if idx == s.size-1 {
		v := trav.Next.Value
		s.setNewTail(trav)
		return v
	}

	trav.Next = trav.Next.Next
	s.size--
	return trav.Value
}

func (s *SingleLinkedList[T]) setNewTail(prev *Node[T]) {
	prev.Next = nil
	s.tail = prev
	s.size--
}

func (s *SingleLinkedList[T]) RemoveFirst() any {
	curr := s.head
	s.head = curr.Next
	s.size--
	curr.Next = nil
	return curr.Value
}

func (s *SingleLinkedList[T]) RemoveLast() any {
	return s.RemoveAt(s.size - 1)
}

func (s *SingleLinkedList[T]) PeekFirst() any {
	return s.head.Value
}

func (s *SingleLinkedList[T]) PeekLast() any {
	return s.tail.Value
}
