package hashtable

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func ExampleSeparateChaining() {
	m := NewSeparateChaining[int, string]()

	m.Insert(1, "John Doe")
	m.Insert(2, "Carl Smith")
	m.Insert(3, "Jane Williams")

	fmt.Println(m.IsEmpty())
	fmt.Println(m.Size())
	fmt.Println(m.Get(1))
	fmt.Println(m.Get(2))
	fmt.Println(m.Get(3))

	m.Remove(1)
	fmt.Println(m.Get(1))
	fmt.Println(m.Size())

	expectedKeys := []int{2, 3}
	gotKeys := m.Keys()

	expectedValues := []string{"Carl Smith", "Jane Williams"}
	gotValues := m.Values()

	slices.Sort(expectedKeys)
	slices.Sort(expectedValues)

	equalKeys := slices.Equal(expectedKeys, gotKeys)
	equalValues := slices.Equal(expectedValues, gotValues)

	fmt.Println(equalKeys)
	fmt.Println(equalValues)

	// Output:
	// false
	// 3
	// John Doe
	// Carl Smith
	// Jane Williams
	//
	// 2
	// true
	// true
}
