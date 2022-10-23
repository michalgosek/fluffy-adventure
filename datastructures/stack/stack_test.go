package stack

import "fmt"

func Example_intStack() {
	s := NewIntArrayIntStack()

	fmt.Println(s.IsEmpty())

	s.Push(1)
	s.Push(2)
	s.Push(3)

	peek := s.Peek()
	fmt.Println(peek)

	s.Push(4)
	fmt.Println(s.Pop())

	//Output:
	// true
	// 1
	// 4
}

func Example_stringStack() {
	s := NewStringArrayStringStack()

	fmt.Println(s.IsEmpty())

	s.Push("A")
	s.Push("B")
	s.Push("C")

	peek := s.Peek()
	fmt.Println(peek)

	s.Push("D")
	fmt.Println(s.Pop())

	//Output:
	// true
	// A
	// D
}
