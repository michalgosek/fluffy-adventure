package hashtable

import (
	"fluffy-adventure/cmp"
	"fluffy-adventure/datastructures/linkedlist"
	"golang.org/x/exp/constraints"
)

type Entry[K constraints.Ordered, V any] struct {
	hash uint64
	k    K
	v    V
}

type SeparateChaining[K constraints.Ordered, V any] struct {
	size  int
	table []*linkedlist.SingleLinkedList[*Entry[K, V]]
	h     HashObject[K]
}

func NewSeparateChaining[K constraints.Ordered, V any]() *SeparateChaining[K, V] {
	s := SeparateChaining[K, V]{
		table: make([]*linkedlist.SingleLinkedList[*Entry[K, V]], 3),
	}
	return &s
}

func (s *SeparateChaining[K, V]) isNilKey(key K) bool {
	var nilKey K
	return nilKey == key
}

func (s *SeparateChaining[K, V]) normalizeIndex(keyHash uint64) int {
	return int(keyHash&0x7FFFFFFF) % cap(s.table)
}

func (s *SeparateChaining[K, V]) bucketSeekEntry(idx int, key K) *Entry[K, V] {
	if s.isNilKey(key) {
		return nil
	}

	bucket := s.table[idx]
	if bucket == nil {
		return nil
	}

	for i := 0; i < bucket.Size(); i++ {
		e := s.getBucketEntry(bucket, i)
		if cmp.Compare(e.k, key) == 0 {
			return e
		}
	}
	return nil
}

func (s *SeparateChaining[K, V]) getBucketEntry(bucket *linkedlist.SingleLinkedList[*Entry[K, V]], i int) *Entry[K, V] {
	e, ok := bucket.GetAt(i).(*Entry[K, V])
	if !ok {
		panic("bucket entry assertion failure")
	}
	return e
}

func (s *SeparateChaining[K, V]) bucketInsertEntry(idx int, e Entry[K, V]) {
	bucket := s.table[idx]
	if bucket == nil {
		bucket = &linkedlist.SingleLinkedList[*Entry[K, V]]{}
		s.table[idx] = bucket
	}

	existing := s.bucketSeekEntry(idx, e.k)
	if existing != nil {
		existing.v = e.v
		return
	}
	bucket.AddLast(&e)
	s.size++
}

func (s *SeparateChaining[K, V]) bucketRemoveEntry(idx int, key K) {
	entry := s.bucketSeekEntry(idx, key)
	if entry == nil {
		return
	}

	bucket := s.table[idx]
	bucket.Remove(entry)
	s.size--
	if bucket.IsEmpty() {
		s.table[idx] = nil
	}
}

func (s *SeparateChaining[K, V]) Size() int {
	return s.size
}

func (s *SeparateChaining[K, V]) IsEmpty() bool {
	return s.size == 0
}

func (s *SeparateChaining[K, V]) Insert(key K, value V) {
	if s.isNilKey(key) {
		panic("specified nil key value")
	}

	entry := Entry[K, V]{
		k:    key,
		v:    value,
		hash: s.h.HashCode(key),
	}
	idx := s.normalizeIndex(entry.hash)
	s.bucketInsertEntry(idx, entry)
}

func (s *SeparateChaining[K, V]) Remove(key K) {
	if s.isNilKey(key) {
		panic("specified nil key value")
	}

	idx := s.normalizeIndex(s.h.HashCode(key))
	s.bucketRemoveEntry(idx, key)
}

func (s *SeparateChaining[K, V]) Get(key K) V {
	var empty V
	if s.isNilKey(key) {
		return empty
	}

	idx := s.normalizeIndex(s.h.HashCode(key))
	entry := s.bucketSeekEntry(idx, key)
	if entry != nil {
		return entry.v
	}
	return empty
}

func (s *SeparateChaining[K, V]) Keys() []K {
	var keys []K
	for _, t := range s.table {
		if t == nil {
			continue
		}
		for i := 0; i < t.Size(); i++ {
			entry := s.getBucketEntry(t, i)
			keys = append(keys, entry.k)
		}
	}
	return keys
}

func (s *SeparateChaining[K, V]) Values() []V {
	var values []V
	for _, t := range s.table {
		if t == nil {
			continue
		}
		for i := 0; i < t.Size(); i++ {
			entry := s.getBucketEntry(t, i)
			values = append(values, entry.v)
		}
	}
	return values
}
