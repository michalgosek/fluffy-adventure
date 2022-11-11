package hashtable

import (
	"fmt"
	"hash/maphash"

	"golang.org/x/exp/constraints"
)

type HashObject[T constraints.Ordered] struct {
	hash maphash.Hash
}

func (h *HashObject[T]) HashCode(v T) uint64 {
	s := fmt.Sprintf("%v", v)
	_, err := h.hash.WriteString(s)
	if err != nil {
		panic(err)
	}
	hash := h.hash.Sum64()
	h.hash.Reset()
	return hash
}
