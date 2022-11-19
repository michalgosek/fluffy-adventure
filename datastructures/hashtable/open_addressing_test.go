package hashtable

import "fmt"

func ExampleOpenAddressing() {
	linear := NewLinearProbing()
	m := NewOpenAddressing[string, string](linear, "tomb")

	fmt.Println(m.IsEmpty())
	m.insert("key1", "key1")
	fmt.Println(m.Get("key1"))
	fmt.Println(m.Size())
	m.insert("key1", "updated")
	fmt.Println(m.Get("key1"))

	m.insert("key2", "key2")
	fmt.Println(m.Size())
	fmt.Println(m.Get("key2"))
	fmt.Println(m.Remove("key2"))
	fmt.Println(m.Size())

	// Output:
	// true
	// key1
	// 1
	// updated
	// 2
	// key2
	// key2
	// 1
}
