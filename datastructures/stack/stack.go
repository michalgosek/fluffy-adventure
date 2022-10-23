package stack

type ArrayStack[T any] struct {
	arr  []T
	size int
	cap  int
}

func NewIntArrayIntStack() *ArrayStack[int] {
	return &ArrayStack[int]{
		cap: 2,
		arr: make([]int, 2),
	}
}

func NewStringArrayStringStack() *ArrayStack[string] {
	return &ArrayStack[string]{
		cap: 2,
		arr: make([]string, 2),
	}
}

func (s *ArrayStack[T]) increaseCapacity() {
	s.cap *= 2
	arr := make([]T, s.cap)
	copy(arr, s.arr)
	s.arr = arr
}

func (s *ArrayStack[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *ArrayStack[T]) Push(elem T) {
	if s.size == len(s.arr) {
		s.increaseCapacity()
	}
	s.arr[s.size] = elem
	s.size++
}

func (s *ArrayStack[T]) Pop() T {
	if s.IsEmpty() {
		panic("empty stack")
	}
	pop := s.arr[s.size-1]
	s.size--
	s.arr = s.arr[:s.size]
	return pop
}

func (s *ArrayStack[T]) Peek() T {
	if s.IsEmpty() {
		panic("empty stack")
	}
	peek := s.arr[0]
	s.arr = s.arr[1:]
	s.size--
	return peek
}
