package hashtable

import (
	"fluffy-adventure/cmp"
	"golang.org/x/exp/constraints"
)

type LinearProbing struct {
	cap    int
	linear int
}

func (p *LinearProbing) SetCapacity(x int) {
	p.cap = x
}

func (p *LinearProbing) Capacity() int {
	return p.cap
}

func (p *LinearProbing) AdjustCapacity() {
	for gcd(p.linear, p.cap) != 1 {
		p.cap++
	}
}

func NewLinearProbing() *LinearProbing {
	l := LinearProbing{
		linear: 17,
	}
	return &l
}

func (p *LinearProbing) Probe(x int) int {
	return p.linear * x
}

type ProbingType interface {
	AdjustCapacity()
	Probe(x int) int
	Capacity() int
	SetCapacity(x int)
}

type OpenAddressing[K constraints.Ordered, V any] struct {
	loadFactor                   float32
	threshold, modificationCount int
	keys                         []K
	values                       []V
	tombstone                    K
	keyCount, usedBuckets        int
	hasher                       HashObject[K]
	prob                         ProbingType
}

func NewOpenAddressing[K constraints.Ordered, V any](p ProbingType, tombstone K) *OpenAddressing[K, V] {
	const defaultCapacity = 7
	const defaultLoadFactor = 0.65

	p.SetCapacity(defaultCapacity)

	threshold := float32(defaultCapacity) * defaultLoadFactor
	s := OpenAddressing[K, V]{
		loadFactor: defaultLoadFactor,
		threshold:  int(threshold),
		tombstone:  tombstone,
		keys:       make([]K, p.Capacity()),
		values:     make([]V, p.Capacity()),
		prob:       p,
	}
	return &s
}

func (o *OpenAddressing[K, V]) increaseCapacity() {
	x := (o.prob.Capacity() * 2) + 1
	o.prob.SetCapacity(x)
}

func (o *OpenAddressing[K, V]) clear() {
	var emptyV V
	var emptyK K
	for i := 0; i < o.prob.Capacity(); i++ {
		o.keys[i], o.values[i] = emptyK, emptyV
	}
	o.keyCount, o.usedBuckets = 0, 0
	o.modificationCount++
}

func (o OpenAddressing[K, V]) normalizeIndex(k int) int {
	return (k & 0x7FFFFFFF) % o.prob.Capacity()
}

func (o *OpenAddressing[K, V]) resizeTable() {
	o.increaseCapacity()
	o.prob.AdjustCapacity()

	o.threshold = int(float32(o.prob.Capacity()) * o.loadFactor)

	newKeys := make([]K, o.prob.Capacity())
	newValues := make([]V, o.prob.Capacity())

	oldKeys := make([]K, len(o.keys))
	oldValues := make([]V, len(o.values))

	copy(oldKeys, o.keys)
	copy(oldValues, o.values)

	o.values = newValues
	o.keys = newKeys

	o.keyCount = 0
	o.usedBuckets = 0

	for i := 0; i < len(oldKeys); i++ {
		var emptyK K
		if oldKeys[i] != emptyK && oldKeys[i] != o.tombstone {
			o.insert(oldKeys[i], oldValues[i])
		}
	}
}

func (o *OpenAddressing[K, V]) insert(key K, value V) V {
	var emptyKey K
	if key == emptyKey {
		panic("nil key")
	}
	if o.usedBuckets >= o.threshold {
		o.resizeTable()
	}

	offset := o.normalizeIndex(int(o.hasher.HashCode(key)))

	for i, x, j := offset, 1, -1; ; i, x = o.normalizeIndex(offset+o.prob.Probe(x)), x+1 {
		var emptyKey K
		if o.keys[i] == o.tombstone {
			if j == -1 {
				j = i
			}
		} else if o.keys[i] != emptyKey {
			if cmp.Compare(o.keys[i], key) == 0 {
				oldValue := o.values[i]
				if j == -1 {
					o.values[i] = value
				} else {
					var emptyValue V
					o.keys[i] = o.tombstone
					o.values[i] = emptyValue
					o.keys[j] = key
					o.values[j] = value
				}
				o.modificationCount++
				return oldValue
			}
		} else {
			if j == -1 {
				o.usedBuckets++
				o.keyCount++
				o.keys[i] = key
				o.values[i] = value
			} else {
				o.keyCount++
				o.keys[j] = key
				o.values[j] = value
			}
			o.modificationCount++

			var empty V
			return empty
		}
	}
}

func (o *OpenAddressing[K, V]) HasKey(key K) bool {
	var emptyKey K
	var emptyValue V
	if key == emptyKey {
		panic("nil key")
	}

	offset := o.normalizeIndex(int(o.hasher.HashCode(key)))
	for i, x, j := offset, 1, -1; ; i, x = o.normalizeIndex(offset+o.prob.Probe(x)), x+1 {
		if o.keys[i] == o.tombstone {
			if j == -1 {
				j = i
			}
		} else if o.keys[i] != emptyKey {
			if cmp.Compare(o.keys[i], key) == 0 {
				if j != -1 { // lazy relocation/deletion
					o.keys[j] = o.keys[i]
					o.values[j] = o.values[i]
					o.keys[i] = o.tombstone
					o.values[i] = emptyValue
				}
				return true
			}
		} else {
			return false
		}
	}
}

func (o *OpenAddressing[K, V]) Get(key K) V {
	var emptyKey K
	var emptyValue V
	if key == emptyKey {
		panic("nil key")
	}

	offset := o.normalizeIndex(int(o.hasher.HashCode(key)))
	for i, j, x := offset, -1, 1; ; i, x = o.normalizeIndex(offset+o.prob.Probe(x)), x+1 {
		if o.keys[i] == o.tombstone {
			if j == -1 {
				j = i
			}
		} else if o.keys[i] != emptyKey {
			if j != -1 {
				o.keys[j] = o.keys[i]
				o.values[j] = o.values[i]
				o.keys[i] = o.tombstone
				o.values[i] = emptyValue
				return o.values[j]
			}
			return o.values[i]
		} else {
			var empty V
			return empty
		}
	}
}

func (o *OpenAddressing[K, V]) Remove(key K) V {
	var emptyKey K
	var emptyValue V

	if key == emptyKey {
		panic("nil key")
	}

	offset := o.normalizeIndex(int(o.hasher.HashCode(key)))
	for i, x := offset, 1; ; i, x = o.normalizeIndex(offset+o.prob.Probe(x)), x+1 {
		if o.keys[i] == o.tombstone {
			continue
		}

		if o.keys[i] == emptyKey {
			var empty V
			return empty
		}

		if cmp.Compare(o.keys[i], key) == 0 {
			o.keyCount--
			o.modificationCount--
			old := o.values[i]
			o.keys[i] = o.tombstone
			o.values[i] = emptyValue
			return old
		}
	}
}

func (o *OpenAddressing[K, V]) Size() int {
	return o.keyCount
}

func (o *OpenAddressing[K, V]) Capacity() int {
	return o.prob.Capacity()
}

func (o *OpenAddressing[K, V]) IsEmpty() bool {
	return o.keyCount == 0
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
