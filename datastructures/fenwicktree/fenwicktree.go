package fenwicktree

type RangeQueryPointUpdate struct {
	n    int
	tree []float32
}

func NewRangeQueryPointUpdate(nums ...float32) *RangeQueryPointUpdate {
	if len(nums) == 0 {
		panic("numbers array cannot be nil")
	}

	N := len(nums)
	nums[0] = 0

	tree := make([]float32, N)
	copy(tree, nums)

	for i := 1; i < N; i++ {
		j := lsb(i) + i
		if j <= N {
			tree[j] += tree[i]
		}
	}

	f := RangeQueryPointUpdate{
		n:    N,
		tree: tree,
	}
	return &f
}

func (r *RangeQueryPointUpdate) prefixSum(i int) float32 {
	var sum float32
	for i != 0 {
		sum += r.tree[i]
		i -= lsb(i)
	}
	return sum
}

func (r *RangeQueryPointUpdate) Add(i int, v float32) {
	for i < r.n {
		r.tree[i] += v
		i += lsb(i)
	}
}

func (r *RangeQueryPointUpdate) Get(i int) float32 {
	return r.Sum(i, i)
}

func (r *RangeQueryPointUpdate) Sum(left, right int) float32 {
	if right < left {
		panic("left bound must be greater or equal left")
	}
	return r.prefixSum(right) - r.prefixSum(left-1)
}

type RangeUpdatePointQuery struct {
	n        int
	current  []float32
	original []float32
}

func NewRangeUpdatePointQuery(nums ...float32) *RangeUpdatePointQuery {
	if len(nums) == 0 {
		panic("numbers array cannot be nil")
	}

	N := len(nums)
	nums[0] = 0

	tree := make([]float32, len(nums))
	copy(tree, nums)

	for i := 1; i < N; i++ {
		j := i + lsb(i)
		if j < N {
			tree[j] += tree[i]
		}
	}

	original := tree
	cp := make([]float32, len(original))
	copy(cp, original)

	r := RangeUpdatePointQuery{
		n:        N,
		current:  cp,
		original: original,
	}
	return &r
}

func (r *RangeUpdatePointQuery) prefixSum(i int, tree []float32) float32 {
	var sum float32
	for i != 0 {
		sum += tree[i]
		i -= lsb(i)
	}
	return sum
}

func (r *RangeUpdatePointQuery) Get(i int) float32 {
	return r.prefixSum(i, r.current) - r.prefixSum(i-1, r.original)
}

func (r *RangeUpdatePointQuery) Add(i int, v float32) {
	for i < r.n {
		r.current[i] += v
		i += lsb(i)
	}
}

func (r *RangeUpdatePointQuery) UpdateRange(left, right int, val float32) {
	r.Add(left, +val)
	r.Add(right+1, -val)
}

func lsb(i int) int {
	return i & -i
}
