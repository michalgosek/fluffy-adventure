package arrays

import "fmt"

func ExampleDynamicArray_int() {
	arr := NewDynamicArray[int]()

	fmt.Println(arr.IsEmpty())

	arr.Append(10)
	arr.Append(3)
	arr.Append(4)

	fmt.Println(arr.Size())

	arr.Set(0, 11)
	fmt.Println(arr.Get(0))

	arr.RemoveAt(2)
	fmt.Println(arr.Get(1))
	fmt.Println(arr.Size())

	// Output:
	// true
	// 3
	// 11
	// 3
	// 2
}
