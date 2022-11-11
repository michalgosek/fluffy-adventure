package hashtable

import "fmt"

func ExampleHashObject() {
	type Person struct {
		Name string
		Age  int
	}

	p1 := Person{
		Name: "Jack",
		Age:  22,
	}
	p2 := p1

	var h HashObject[int]

	h1 := h.HashCode(p1.Age)
	h2 := h.HashCode(p2.Age)

	fmt.Println(h1 == h2)

	// Output:
	// true
}
