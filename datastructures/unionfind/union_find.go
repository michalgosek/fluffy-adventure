package unionfind

type UnionFind struct {
	size          int
	sz            []int
	id            []int
	numComponents int
}

func NewUnionFind(size int) *UnionFind {
	if size <= 0 {
		panic("size <= 0 is not allowed")
	}

	u := UnionFind{
		size:          size,
		sz:            make([]int, size),
		id:            make([]int, size),
		numComponents: size,
	}

	for i := 0; i < size; i++ {
		u.id[i] = i
		u.sz[i] = 1
	}
	return &u
}

func (u *UnionFind) Find(p int) int {
	root := p
	for root != u.id[root] {
		root = u.id[root]
	}

	for p != root {
		next := u.id[p]
		u.id[p] = root
		p = next
	}
	return root
}

func (u *UnionFind) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *UnionFind) ComponentSize(p int) int {
	return u.sz[u.Find(p)]
}

func (u *UnionFind) Components() int {
	return u.numComponents
}

func (u *UnionFind) Unify(p, q int) {
	if u.Connected(p, q) {
		return
	}

	root1 := u.Find(p)
	root2 := u.Find(q)

	if u.sz[root1] < u.sz[root2] {
		u.sz[root2] += u.sz[root1]
		u.id[root1] = root2
		u.sz[root1] = 0
	} else {
		u.sz[root1] += u.sz[root2]
		u.id[root2] = root1
		u.sz[root2] = 0
	}
	u.numComponents--
}
