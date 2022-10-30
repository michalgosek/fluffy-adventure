package unionfind

import (
	"fmt"
)

func ExampleNewUnionFind() {
	uf := NewUnionFind(5)

	fmt.Println(uf.Components())

	uf.Unify(0, 1)
	fmt.Println(uf.Components())

	uf.Unify(1, 0)
	fmt.Println(uf.Components())

	uf.Unify(1, 2)
	fmt.Println(uf.Components())

	uf.Unify(0, 2)
	fmt.Println(uf.Components())

	uf.Unify(2, 1)
	fmt.Println(uf.Components())

	uf.Unify(3, 4)
	fmt.Println(uf.Components())

	uf.Unify(4, 3)
	fmt.Println(uf.Components())

	uf.Unify(1, 3)
	fmt.Println(uf.Components())

	uf.Unify(4, 0)
	fmt.Println(uf.Components())

	// Output:
	// 5
	// 4
	// 4
	// 3
	// 3
	// 3
	// 2
	// 2
	// 1
	// 1
}
