package fenwicktree

import "fmt"

func ExampleFenwickTree() {
	nums := []float32{0, 1, 2, 2, 4}

	f := NewRangeQueryPointUpdate(nums...)
	fmt.Println(f.Sum(1, 4))

	f.Add(3, 1)
	fmt.Println(f.Get(3))

	// Output:
	// 9
	// 3
}

func ExampleFenwickTreeRangeUpdate() {
	f := NewRangeUpdatePointQuery(0, 1, 2)
	f.UpdateRange(1, 2, 1)

	fmt.Println(f.Get(1))
	fmt.Println(f.Get(2))

	// Output:
	// 2
	// 3
}
