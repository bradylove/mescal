package store

import (
	"errors"
	"hash/fnv"
	"sync/atomic"
	"unsafe"

	"github.com/bradylove/mescal/pkg/mescal"
)

var (
	bucketCount     = uint64(104729)
	sliceCount      = uint64(41)
	allowableOffset = uint64(5)

	ErrKeyNotFound = errors.New("Key not found")
)

type Store struct {
	cache [][]unsafe.Pointer
}

func New() *Store {
	s := Store{
		cache: make([][]unsafe.Pointer, sliceCount),
	}

	for idx := range s.cache {
		s.cache[uint64(idx)] = make([]unsafe.Pointer, bucketCount)
	}

	return &s
}

func (s Store) Set(kv *mescal.KeyValue) error {
	slIdx, bIdx := buildIndex(kv.Key)

	for i := uint64(0); i < allowableOffset; i++ {
		newIdx := (bIdx + i) % bucketCount

		old := (*mescal.KeyValue)(atomic.LoadPointer(&s.cache[slIdx][newIdx]))

		if old == nil || old.Key == kv.Key {
			atomic.SwapPointer(
				&s.cache[slIdx][newIdx],
				unsafe.Pointer(kv),
			)

			return nil
		}
	}

	return nil
}

func (s Store) Get(key string) (*mescal.KeyValue, error) {
	slIdx, bIdx := buildIndex(key)

	for i := uint64(0); i < allowableOffset; i++ {
		newIdx := (bIdx + i) % bucketCount

		kv := (*mescal.KeyValue)(atomic.LoadPointer(&s.cache[slIdx][newIdx]))

		if kv != nil && kv.Key == key {
			return kv, nil
		}
	}

	return nil, ErrKeyNotFound
}

func buildIndex(key string) (sliceIdx uint64, bucketIdx uint64) {
	f := fnv.New64a()
	f.Write([]byte(key))

	sum := f.Sum64()

	return sum % sliceCount, sum % bucketCount
}
